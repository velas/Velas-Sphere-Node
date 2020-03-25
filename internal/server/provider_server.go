package server

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
)

type ProviderServer struct {
	PluginClient resources.PluginClient
	// EthClient    *ethclient.Client
}

func (p ProviderServer) RequestTaskExecution(ctx context.Context, req *resources.TaskExecutionRequest) (*resources.TaskExecutionRequestResponse, error) {
	if p.PluginClient == nil {
		return nil, errors.New("no plugin provided")
	}
	pluginClient := p.PluginClient

	// if p.EthClient == nil {
	// 	return nil, errors.New("no eth client provided")
	// }
	// ethClient := p.EthClient

	resp, err := pluginClient.RequestJobExecution(ctx, &resources.JobExecutionRequest{
		Id:    req.GetId(),
		Input: req.GetInput(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not request job execution from plugin")
	}

	// // TODO: inject the private key
	// privateKey, err := crypto.HexToECDSA("")
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to parse key hex")
	// }

	// publicKey := privateKey.PublicKey

	// fromAddress := crypto.PubkeyToAddress(publicKey)
	// nonce, err := p.EthClient.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)     // in wei
	// auth.GasLimit = uint64(300000) // in units
	// auth.GasPrice = gasPrice

	// address := crypto.PubkeyToAddress(publicKey)

	// // TODO: provide real numbers
	// c, err := contract.NewEthdepositcontract(address, ethClient)
	// invoiceTx, err := c.CreateInvoice(nil, nil, nil, address, nil, nil, nil, nil)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to create tx")
	// }

	// // TODO: use invoiceTx?
	// _ = invoiceTx

	return &resources.TaskExecutionRequestResponse{
		Id:     resp.GetId(),
		Output: resp.GetOutput(),
	}, nil
}
