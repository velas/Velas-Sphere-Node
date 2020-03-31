package handler

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/sorenvonsarvort/velas-sphere/internal/entity"
	"github.com/sorenvonsarvort/velas-sphere/internal/merkletree"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"google.golang.org/grpc"
)

func store(data string, challenges []string) (merkletree.Node, error) {
	items := []string{}

	for _, challenge := range challenges {
		items = append(items, data+challenge)
	}

	return merkletree.CalculateRoot(items...)
}

// TODO: improve
func encrypt(key []byte, message []byte) (encmess []byte, err error) {
	plainText := message

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = cipherText
	return
}

// TODO: improve
func decrypt(key []byte, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Ciphertext block size is too short!")
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("Ciphertext block size is too short!")
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func NewPostFileHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.DB == nil {
			log.Println("no db provided")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		db := config.DB

		requestBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to read the request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		request := entity.File{}

		err = json.Unmarshal(requestBytes, &request)
		if err != nil {
			log.Println("failed to unmarshal the request body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// TODO: generate random key
		aeskey := []byte("myawesomekey0000")

		encryptedFile, err := encrypt(aeskey, request.Data)
		if err != nil {
			log.Println("failed to encrypt")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// TODO: generate random challenges
		challenges := []string{
			"x",
			"w",
			"z",
			"j",
		}

		merkleTree, err := store(string(encryptedFile), challenges)
		if err != nil {
			log.Println("failed create the merkle tree")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		merkleTreeJSONBytes, err := json.Marshal(merkleTree)
		if err != nil {
			log.Println("failed to marshal the merkle tree:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		challengeToPath := map[string]string{}

		for _, challenge := range challenges {
			hash := sha256.New()
			hash.Write([]byte(string(encryptedFile) + challenge))
			hashBytes := hash.Sum(nil)

			path, err := merkletree.FindPath(merkleTree, hashBytes)
			if err != nil {
				log.Println("failed to find the challenge path:", err, hex.EncodeToString(hashBytes))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// TODO: store challengeToPathMap into db?
			challengeToPath[challenge] = strings.Join(path, ":")
		}

		conn, err := grpc.Dial(request.Target, grpc.WithInsecure())
		if err != nil {
			log.Println("failed to dial the target")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		c := resources.NewStorageClient(conn)

		providerResponse, err := c.SaveFile(
			r.Context(),
			&resources.FileStorageRequest{
				Name:           request.Name,
				Data:           encryptedFile,
				MerkleTreeRoot: merkleTreeJSONBytes,
			},
		)
		if err != nil {
			log.Println("failed to request file storage:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id := providerResponse.GetId()
		verificationToken := providerResponse.GetVerificationToken()

		fileJSONBytes, err := json.Marshal(
			entity.File{
				ID:                 id,
				DecryptionKey:      aeskey,
				MerkleTreeRoot:     merkleTree,
				Target:             request.Target,
				GetBackToken:       providerResponse.GetGetBackToken(),
				VerificationToken:  verificationToken,
				RequesterPublicKey: "requester_public_key", // TODO: replace by the real one
			},
		)
		if err != nil {
			log.Println("failed to marshal the file: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = db.Put([]byte("requested_files/"+id), fileJSONBytes, nil)
		if err != nil {
			log.Println("failed to save the file:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// verification routine
		go func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Println("got panic:", err)
				}
			}()

			challengeToPath := challengeToPath
			id := id
			verificationToken := verificationToken

			// TODO: send challenges and verify them
			for challenge, path := range challengeToPath {
				conn, err := grpc.Dial(request.Target, grpc.WithInsecure())
				if err != nil {
					log.Println("failed to dial the target")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer conn.Close()

				c := resources.NewStorageClient(conn)

				verification, err := c.VerifyFileStorage(
					context.TODO(),
					&resources.FileStorageVerificationRequest{
						Id:                         id,
						Challenge:                  []byte(challenge),
						VerificationTokenSignature: verificationToken, // TODO: replace by a real signature
					},
				)

				if err != nil {
					log.Println("verification failed:", err)
					return
				}

				if verification.GetPath() != path {
					// TODO: request ban!
					log.Println("file storage verification failed", path, "verification path:", verification.GetPath())
					return
				}

				verificationToken = verification.GetNewVerificationToken()

				log.Println("verification ok!")

				time.Sleep(time.Minute)
			}
		}()

		response, err := json.Marshal(
			map[string]interface{}{
				"id": id,
			},
		)
		if err != nil {
			log.Println("failed to marshal the provider response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(response)
	}
}
