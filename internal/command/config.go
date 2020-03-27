package command

type NodeConfig struct {
	PluginTarget        string `json:"plugin_target"`
	StoragePluginTarget string `json:"storage_plugin_target"`
	EthereumNodeTarget  string `json:"ethereum_node_target"`
}

type Config struct {
	Node NodeConfig `json:"node"`
}
