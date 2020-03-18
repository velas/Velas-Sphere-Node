package command

type ProviderConfig struct {
	PluginTarget string `json:"plugin_target"`
}

type Config struct {
	Provider ProviderConfig `json:"provider"`
}
