package main

import (
	"fmt"

	_ "github.com/RaftechNL/terrafile/internal/providers/gh"
	"github.com/RaftechNL/terrafile/internal/terrafilev2"
)

func main() {

	tfv2, errv2 := terrafilev2.InitTerrafileV2FromYAML("v2.yaml")
	if errv2 != nil {
		fmt.Println(errv2)
	}

	for _, module := range tfv2.Modules {
		fmt.Println(module.Source)
		// x := tfv2.Providers[module.ProviderAliasRef]

		// fmt.Println(x)

		// module.Provider.DownloadModule(&module.ProviderConfig)

		fmt.Println(module)
	}

}
