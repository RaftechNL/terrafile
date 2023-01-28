package terrafilev2

import (
	"io/ioutil"

	"github.com/RaftechNL/terrafile/internal/providers/config"
	"github.com/RaftechNL/terrafile/internal/tf"
	"gopkg.in/yaml.v2"
)

type TerrafileV2 struct {
	Providers map[string]config.ProviderConfig `yaml:"providers"`
	Modules   map[string]tf.Module             `yaml:"modules"`
}

func ReadFromYAML(filePath string) (*TerrafileV2, error) {
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
