package initializer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sorenvonsarvort/velas-sphere/internal/contract"
)

func TransactOptions(ethClient *ethclient.Client) func(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	return func(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
		publicKey := key.PublicKey

		fromAddress := crypto.PubkeyToAddress(publicKey)
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get pending nonce: %w", err)
		}

		gasPrice, err := ethClient.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get suggested gas price: %w", err)
		}

		auth := bind.NewKeyedTransactor(key)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		return auth, nil
	}
}

func ContractInitializer(ethClient *ethclient.Client, address common.Address) (*contract.Ethdepositcontract, error) {
	c, err := contract.NewEthdepositcontract(address, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new deposit contract: %w", err)
	}

	return c, nil
}
