package handler

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

	ethdepositcontract "github.com/velas/Velas-Sphere-Contracts"
)

func NewPostFileHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		initOptions := config.TransactOptionsInitializer
		if initOptions == nil {
			return
		}

		ethDepositContract := config.Contract

		_ = ethdepositcontract.EthdepositcontractABI

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

		rawKey := r.FormValue("key")

		keyBytes, err := hex.DecodeString(rawKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to decode the key hex:", err)
			return
		}

		privateKey, err := crypto.ToECDSA(keyBytes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to decode the key:", err)
			return
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("got invalid public key")
			return
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		// TODO: check node registration

		registerOpts, err := initOptions(privateKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to init opts")
			return
		}
		membershipFee := int64(100000000000)
		registerOpts.Value = big.NewInt(membershipFee)

		registerTx, err := ethDepositContract.RegisterNode(registerOpts, fromAddress, nil, nil, nil, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to register the node:", err)
			return
		}
		_ = registerTx

		source, _, err := r.FormFile("file")

		defer source.Close()

		_, err = io.Copy(target, source)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to copy the file:", err)
			return
		}

		invoiceOpts, err := initOptions(privateKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to init opts")
			return
		}

		tx, err := ethDepositContract.CreateInvoice(invoiceOpts, nil, nil, fromAddress, nil, nil, nil, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("failed to create the invoice:", err)
			return
		}

		responseBytes, err := json.Marshal(
			map[string]interface{}{
				"id":      id,
				"invoice": tx.Hash,
			},
		)

		w.Write(responseBytes)
	}
}
