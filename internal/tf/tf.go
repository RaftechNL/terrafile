package tf

import (
	"github.com/RaftechNL/terrafile/internal/providers"
)

type Module struct {
	providers.ModuleSpec `yaml:",inline"`
	ProviderAliasRef     string `yaml:"providerAliasRef"`

	Provider *providers.ProviderIface
}

func (m *Module) SetProvider(provider *providers.ProviderIface) {
	m.Provider = provider
}

func (m *Module) GetProviderAliasRef() string {
	return m.ProviderAliasRef
}

func (m *Module) Download(outputPath string) error {
	return (*m.Provider).DownloadModule(m.ModuleSpec, outputPath)
}
