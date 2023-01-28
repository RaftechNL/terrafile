package main

import (
	"fmt"

	"github.com/RaftechNL/terrafile/internal/terrafilev2"
)

func main() {

	tfv2, errv2 := terrafilev2.ReadFromYAML("v2.yaml")
	if errv2 != nil {
		fmt.Println(errv2)
	}

	fmt.Println(tfv2)

}
