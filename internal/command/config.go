package command

type NodeConfig struct {
	PluginTarget string `json:"plugin_target"`
}

type Config struct {
	Provider NodeConfig `json:"provider"`
}
