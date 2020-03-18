package command

type NodeConfig struct {
	PluginTarget string `json:"plugin_target"`
}

type Config struct {
	Node NodeConfig `json:"node"`
}
