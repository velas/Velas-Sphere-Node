package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func NewPostFileHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// initContract := config.ContractInitializer
		// if initContract == nil {
		// 	return
		// }

		randomID := make([]byte, 16)
		_, err := io.ReadFull(rand.Reader, randomID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to read all from rand reader:", err)
			return
		}

		id := base64.URLEncoding.EncodeToString(randomID)

		target, err := os.Create(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to open the file:", err)
			return
		}

		err = r.ParseMultipartForm(1000 << 20)
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to parse the form:", err)
			// return
		}

		// rawKey := r.FormValue("key")

		// keyBytes, err := hex.DecodeString(rawKey)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	log.Println("failed to decode the key hex:", err)
		// 	return
		// }

		// privateKey, err := crypto.ToECDSA(keyBytes)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	log.Println("failed to decode the key:", err)
		// 	return
		// }

		// ethDepositContract, err := initContract(privateKey)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	log.Println("failed to init the contract:", err)
		// 	return
		// }

		// publicKey := privateKey.Public()
		// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		// if !ok {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	log.Println("got invalid public key")
		// 	return
		// }

		// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		// TODO: check node registration

		// registerTx, err := ethDepositContract.RegisterNode(nil, fromAddress, nil, nil, nil, nil)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	log.Println("failed to register the node:", err)
		// 	return
		// }
		// _ = registerTx

		source, _, err := r.FormFile("file")

		defer source.Close()

		_, err = io.Copy(target, source)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to copy the file:", err)
			return
		}

		// tx, err := ethDepositContract.CreateInvoice(nil, nil, nil, fromAddress, nil, nil, nil, nil)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	log.Println("failed to create the invoice:", err)
		// 	return
		// }

		responseBytes, err := json.Marshal(
			map[string]interface{}{
				"id": id,
				// "invoice": tx.Hash,
			},
		)

		w.Write(responseBytes)
	}
}
