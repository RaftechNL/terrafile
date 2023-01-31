package terrafilev2

import (
	"io/ioutil"

	"github.com/RaftechNL/terrafile/internal/providers"
	"github.com/RaftechNL/terrafile/internal/providers/config"
	"github.com/RaftechNL/terrafile/internal/providers/gh"
	"github.com/RaftechNL/terrafile/internal/tf"
	"gopkg.in/yaml.v2"
)

type TerrafileV2 struct {
	ProvidersConfigs map[string]config.ProviderConfig `yaml:"providers"`
	ProvidersFactory map[string]providers.ProviderIface
	Modules          map[string]tf.Module `yaml:"modules"`
}

func (t *TerrafileV2) initProvidersFactory() {
	t.ProvidersFactory = make(map[string]providers.ProviderIface)

	for name, providerConfig := range t.ProvidersConfigs {
		switch providerConfig.ProviderType {
		case "github":
			t.ProvidersFactory[name] = gh.NewGithubProvider(&providerConfig)
		}
	}

}

func (t *TerrafileV2) GetProviderByAliasRef(key string) *providers.ProviderIface {
	value, ok := t.ProvidersFactory[key]
	if !ok {
		return nil
	}

	return &value
}

// InitTerrafileV2FromYAML reads the terrafile from the given path and returns a TerrafileV2 struct
func InitTerrafileV2FromYAML(filePath string) (*TerrafileV2, error) {
	terraform, err := readFromYAML(filePath)
	if err != nil {
		return nil, err
	}

	terraform.initProvidersFactory()

	return terraform, nil
}

func readFromYAML(filePath string) (*TerrafileV2, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var terraform TerrafileV2
	err = yaml.Unmarshal(data, &terraform)
	if err != nil {
		return nil, err
	}

	return &terraform, nil
}
