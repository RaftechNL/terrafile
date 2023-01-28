package tf

import (
	"github.com/RaftechNL/terrafile/internal/providers"
	"github.com/RaftechNL/terrafile/internal/providers/config"
	_ "github.com/RaftechNL/terrafile/internal/providers/gh"
)

type Module struct {
	Source           string `yaml:"source"`
	Version          string `yaml:"version"`
	ProviderAliasRef string `yaml:"providerAliasRef"`
	Provider         providers.ProviderIface
	ProviderConfig   *config.ProviderConfig
}

// func (tm *TerraformModule) Download(modulePath string) error {

// 	err := gh.DownloadRepository(tm.Source, tm.Version, modulePath)

// 	if err != nil {
// 		return err
// 	}

// 	tm.Provider.DownloadModule()
// 	return nil
// }
