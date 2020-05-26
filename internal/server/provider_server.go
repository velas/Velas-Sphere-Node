package server

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/contract"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
)

type ProviderServer struct {
	PluginClient       resources.PluginClient
	Ethdepositcontract *contract.Ethdepositcontract
}

func (p ProviderServer) RequestTaskExecution(ctx context.Context, req *resources.TaskExecutionRequest) (*resources.TaskExecutionResponse, error) {
	if p.PluginClient == nil {
		return nil, errors.New("no plugin provided")
	}
	pluginClient := p.PluginClient

	// if p.Ethdepositcontract == nil {
	// 	return nil, errors.New("no contract instance injected")
	// }
	// ethDepositContract := p.Ethdepositcontract

	resp, err := pluginClient.RequestJobExecution(
		ctx,
		&resources.JobExecutionRequest{
			Id:    req.GetId(),
			Input: req.GetInput(),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not request job execution from plugin")
	}

	// invoiceTx, err := ethDepositContract.CreateInvoice(nil, nil, nil, common.Address{}, nil, nil, nil, nil)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to create tx")
	// }

	// TODO: use invoiceTx?
	// _ = invoiceTx

	return &resources.TaskExecutionResponse{
		// Took: diff
		// InvoiceID: invoiceTx.Hash().String()
		Id:     resp.GetId(),
		Output: resp.GetOutput(),
	}, nil
}
