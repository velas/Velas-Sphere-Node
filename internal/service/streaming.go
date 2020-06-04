package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Index struct {
	Node string `json:"node"`
	ID   string `json:"id"`
	Key  string `json:"key"`
	// Size?
}

type SuperIndex struct {
	Parts   [][]Index `json:"parts"`
	Title   string    `json:"title"`
	Len     string    `json:"len"`
	Preview Index     `json:"preview"`
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
		return nil, fmt.Errorf("failed to create new cipher: %w", err)
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

func store(item []byte, keyHex string) (Index, error) {

	aeskey := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, aeskey)
	if err != nil {
		return Index{}, fmt.Errorf("failed to read all from rand reader: %w", err)
	}

	encryptedFile, err := encrypt(aeskey, item)
	if err != nil {
		return Index{}, fmt.Errorf("failed to encrypt: %w", err)
	}

	node := "127.0.0.1:3000"

	buf := bytes.Buffer{}

	multipartWriter := multipart.NewWriter(&buf)
	multipartWriter.WriteField("key", keyHex)

	part, err := multipartWriter.CreateFormFile("file", "file")
	if err != nil {
		return Index{}, fmt.Errorf("failed to create the multipart file: %w", err)
	}

	_, err = part.Write(encryptedFile)
	if err != nil {
		return Index{}, fmt.Errorf("failed to write the multipart file: %w", err)
	}

	multipartWriter.Close()

	req, err := http.NewRequest("POST", "http://"+node+"/file", &buf)
	if err != nil {
		return Index{}, fmt.Errorf("failed to create the request: %w", err)
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return Index{}, fmt.Errorf("failed to perform the request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Index{}, fmt.Errorf("failed to read the resp body: %w", err)
	}

	respMap := map[string]interface{}{}
	err = json.Unmarshal(respBody, &respMap)
	if err != nil {
		return Index{}, fmt.Errorf("failed to unmarshal the response: %w", err)
	}

	askeyBase64 := base64.URLEncoding.EncodeToString(aeskey)

	return Index{
		Node: node,
		ID:   respMap["id"].(string),
		Key:  askeyBase64,
	}, nil
}

func upload() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// 1000MB
		err := r.ParseMultipartForm(1000 << 20)
		if err != nil {
			fmt.Println("failed to parse multipart form", err)
			return
		}

		keyHex := r.FormValue("key")

		len := r.FormValue("len")

		videoTitle := r.FormValue("title")

		file, _, err := r.FormFile("video")
		if err != nil {
			fmt.Println("failed to from file", err)
			return
		}

		defer file.Close()

		parts := [][]Index{}
		i := 0
		for {
			replica := []Index{}

			buf := make([]byte, 10<<20)

			n, _ := io.ReadFull(file, buf)
			// if err != nil {
			// 	if n != 10<<20 {
			// 		fmt.Println("failed to read the file:", n, err)
			// 		return
			// 	}
			// }
			if n == 0 {
				break
			}

			for r := 0; r < 3; r++ {
				storedPart, err := store(buf[:n], keyHex)
				if err != nil {
					fmt.Println("failed to store the part:", err)
					return
				}

				fmt.Println("part", i, "replica", r, storedPart.ID)

				replica = append(replica, storedPart)

				if err == io.ErrUnexpectedEOF {
					break
				}
				if err != nil {
					fmt.Println("error happened:", err)
					return
				}

			}
			parts = append(parts, replica)
			i++
		}

		previewFile, _, err := r.FormFile("preview")
		if err != nil {
			fmt.Println("failed to get form file", err)
			return
		}

		defer previewFile.Close()

		previewFileBytes, err := ioutil.ReadAll(previewFile)
		if err != nil {
			fmt.Println("failed to read the preview", err)
			return
		}

		previewIndex, err := store(previewFileBytes, keyHex)
		if err != nil {
			fmt.Println("failed to store the preview", err)
			return
		}

		superIndex := SuperIndex{
			Parts:   parts,
			Title:   videoTitle,
			Len:     len,
			Preview: previewIndex,
		}

		superIndexBytes, err := json.Marshal(superIndex)
		if err != nil {
			fmt.Println("failed to marshal the index:", err)
			return
		}

		storedSuperIndex, err := store(superIndexBytes, keyHex)
		if err != nil {
			fmt.Println("failed to store the index:", err)
			return
		}

		fmt.Println("super index:", storedSuperIndex.ID)

		// TODO: superIndexLink
		responseJSONBytes, err := json.Marshal(
			map[string]interface{}{
				"id":   storedSuperIndex.ID,
				"node": storedSuperIndex.Node,
				"key":  storedSuperIndex.Key,
			},
		)
		if err != nil {
			fmt.Println("failed to marshal json: ", err)
			return
		}

		w.Write(responseJSONBytes)
	}
}

func load(id string) ([]byte, error) {

	node := "127.0.0.1:3000"

	req, _ := http.NewRequest("GET", "http://"+node+"/file/"+id, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform the req: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	return bodyBytes, nil
}

func download() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		superIndexID := chi.URLParam(r, "id")

		superIndexKeyBase64 := chi.URLParam(r, "key")

		superIndexKey, err := base64.URLEncoding.DecodeString(superIndexKeyBase64)
		if err != nil {
			fmt.Println("failed to decode the key:", err)
			return
		}

		superIndexBytes, err := load(superIndexID)
		if err != nil {
			fmt.Println("failed to load the super index:", err)
			return
		}

		superIndexBytes, err = decrypt(superIndexKey, superIndexBytes)
		if err != nil {
			fmt.Println("failed to decrypt the super index:", err)
			return
		}

		superIndex := SuperIndex{}

		err = json.Unmarshal(superIndexBytes, &superIndex)
		if err != nil {
			fmt.Println("failed to unmarshal the super index:", err)
			return
		}

		buf := []byte{}
		for _, part := range superIndex.Parts {
			for i, replica := range part {

				replicaBytes, err := load(replica.ID)
				if err != nil {
					if i == len(part)-1 {
						fmt.Println("lost the part:", err)
						return
					}
					fmt.Println("failed to load the part:", err)
					continue
				}

				replicaKeyDecoded, err := base64.URLEncoding.DecodeString(replica.Key)
				if err != nil {
					fmt.Println("failed to decode the part key:", err)
					return
				}

				decryptedReplicaBytes, err := decrypt(replicaKeyDecoded, replicaBytes)
				if err != nil {
					fmt.Println("failed to decrypt the part:", err)
					return
				}

				buf = append(buf, decryptedReplicaBytes...)
				break
			}

		}

		// ioutil.WriteFile("test.mp4", buf, 0666)

		// f, _ := os.Open("test.mp4")

		// sizeFunc := func(f io.Seeker) (int64, error) {
		// 	size, err := f.Seek(0, io.SeekEnd)
		// 	if err != nil {
		// 		return 0, fmt.Errorf("failed to seek: %w", err)
		// 	}
		// 	_, err = f.Seek(0, io.SeekStart)
		// 	if err != nil {
		// 		return 0, fmt.Errorf("failed to seek: %w", err)
		// 	}
		// 	return size, nil
		// }

		// fmt.Println(sizeFunc(f))

		http.ServeContent(w, r, "1", time.Now(), bytes.NewReader(buf))
	}
}

func Streaming() error {
	r := chi.NewRouter()

	r.Post("/upload", upload())
	r.Get("/{id}/{key}", download())

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}
	return nil
}
