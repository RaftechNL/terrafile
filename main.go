package main

import (
	"fmt"

	"github.com/RaftechNL/terrafile/internal/terrafilev2"
)

const (
	defaultModuleOutputPath = "./vendorski"
)

func main() {

	tfv2, errv2 := terrafilev2.InitTerrafileV2FromYAML("v2.yaml")
	if errv2 != nil {
		fmt.Println(errv2)
	}

	for alias, module := range tfv2.Modules {
		fmt.Println(module.Source)

		providerAliasRef := module.GetProviderAliasRef()

		//TODO: Handle nil pointer situations! Critical panic :)
		module.SetProvider(tfv2.GetProviderByAliasRef(providerAliasRef))
		module.Download(fmt.Sprintf("%s/%s", defaultModuleOutputPath, alias))
	}

}
