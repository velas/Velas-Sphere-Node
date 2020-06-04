package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/velas/Velas-Sphere-Contracts/ethdepositcontract"
)

func contractDeploymentDemo() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	k, err := keystore.DecryptKey(
		[]byte(
			`{"version":3,"id":"0605e0f8-6053-4ce8-b980-c6c0ac618515","address":"dd3332013e1d12885ab6a3d970acd83314e4b39b","Crypto":{"ciphertext":"7617ab91d05bb2b940de3a0c5daf8f29206153f9382fcaa64d143dccd759066a","cipherparams":{"iv":"9d9369f7d242c2b107d68ee450f8d34c"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"f7cd65947598e2cb7d96eb21bffc17041543b477102f4f34c8cf8788c7b61044","n":8192,"r":8,"p":1},"mac":"13a16e18236508fb7f8706628e0deab8e3950669810cc961b9ffad1050211737"}}`,
		),
		"wowshit111",
	)
	privateKey := k.PrivateKey

	// keyBytes := crypto.FromECDSA(privateKey)

	// keyHex := hex.EncodeToString(keyBytes)

	fmt.Println(crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
	// return

	// fmt.Println(privateKey, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.TODO(), fromAddress, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice = big.NewInt(10)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := ethdepositcontract.DeployEthdepositcontract(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	_ = instance
}
