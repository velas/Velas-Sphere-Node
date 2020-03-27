package handler

import (
	"context"
	"crypto/aes"
	"encoding/json"
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
			log.Println("failed to unmarshal the request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// TODO: improve key management

		cipher, err := aes.NewCipher([]byte("mycoolkey"))
		if err != nil {
			log.Println("failed to create the cipher")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		encryptedFile := []byte{}

		cipher.Encrypt(request.Data, encryptedFile)

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

		_ = merkleTree

		challengeToPath := map[string]string{}

		for _, challenge := range challenges {
			path, err := merkletree.FindPath(merkleTree, challenge)
			if err != nil {
				log.Println("failed to find the challenge path")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// TODO: store challengeToPathMap?
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

		providerResponse, err := c.SaveFile(r.Context(), &resources.FileStorageRequest{
			Name: request.Name,
			Data: string(request.Data),
		})
		if err != nil {
			log.Println("failed to request task execution")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fileJSONBytes, err := json.Marshal(
			entity.File{
				ID:                providerResponse.GetId(),
				DecryptionKey:     "cipher",     // TODO: replace by the real aes cipher
				MerkleTreeRoot:    "merkleTree", // TODO: replace by the real merkle tree
				Target:            request.Target,
				GetBackToken:      providerResponse.GetGetBackToken(),
				VerificationToken: providerResponse.GetVerificationToken(),
				// RequesterPublicKey: do we really need to save it?
			},
		)
		if err != nil {
			log.Println("failed to marshal the file: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = db.Put([]byte(providerResponse.GetId()), fileJSONBytes, nil)
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
			id := providerResponse.GetId()

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
						Challenge:                  challenge,
						VerificationTokenSignature: "signature", // TODO: replace by a real signature
					},
				)

				if verification.GetPath() != path {
					// TODO: request ban!
					log.Println("file storage verification failed!")
					return
				}

				time.Sleep(time.Minute)
			}
		}()

		providerResponseJSONBytes, err := json.Marshal(providerResponse)
		if err != nil {
			log.Println("failed to marshal the provider response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(providerResponseJSONBytes)
	}
}
