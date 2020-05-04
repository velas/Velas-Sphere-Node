package command

import "github.com/ethereum/go-ethereum/common"

type NodeConfig struct {
	PluginTarget              string         `json:"plugin_target"`
	EthereumNodeTarget        string         `json:"ethereum_node_target"`
	KeystoreFilePath          string         `json:"keystore_file_path"`
	KeystorePassword          []byte         `json:"keystore_password"`
	EthdepositcontractAddress common.Address `json:"ethdepositcontract_address"`
}

type Config struct {
	Node NodeConfig `json:"node"`
}
