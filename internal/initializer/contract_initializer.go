package initializer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ethdepositcontract "github.com/velas/Velas-Sphere-Contracts/ethdepositcontract"
)

func TransactOptions(ethClient *ethclient.Client) func(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	return func(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
		publicKey := key.PublicKey

		fromAddress := crypto.PubkeyToAddress(publicKey)
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get pending nonce: %w", err)
		}

		txC, err := ethClient.PendingTransactionCount(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get pending tx count: %w", err)
		}

		gasPrice, err := ethClient.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get suggested gas price: %w", err)
		}

		fmt.Println("got suggested gas price:", gasPrice)

		gasPrice.Set(big.NewInt(1000000000 + int64(rand.Int()%1000))) // using randomization for avoiding same-hash transactions.

		fmt.Println("used gas price:", gasPrice)

		auth := bind.NewKeyedTransactor(key)
		auth.Nonce = big.NewInt(int64(nonce) + int64(txC))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		return auth, nil
	}
}

func ContractInitializer(ethClient *ethclient.Client, address common.Address) (*ethdepositcontract.Ethdepositcontract, error) {
	c, err := ethdepositcontract.NewEthdepositcontract(address, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new deposit contract: %w", err)
	}

	return c, nil
}
