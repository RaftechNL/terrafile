package config

type ProviderConfig struct {
	ProviderType string       `yaml:"providerType,omitempty"`
	Auth         ProviderAuth `yaml:"auth"`
}

type ProviderAuth struct {
	Token    *string `yaml:"token"`
	SSHKey   *string `yaml:"ssh_key"`
	Username *string `yaml:"username"`
	Password *string `yaml:"password"`
}
