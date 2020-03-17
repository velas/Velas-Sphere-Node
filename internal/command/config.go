package command

type ProviderConfig struct {
	PluginTarget string `json:"plugin_target"`
}

type RequesterConfig struct {
	Target string `json:"target"`
}

type Config struct {
	Provider        ProviderConfig  `json:"provider"`
	RequesterConfig RequesterConfig `json:"requester"`
}
