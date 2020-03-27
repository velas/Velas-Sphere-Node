package command

type NodeConfig struct {
	PluginTarget       string `json:"plugin_target"`
	EthereumNodeTarget string `json:"ethereum_node_target"`
}

type Config struct {
	Node NodeConfig `json:"node"`
}
