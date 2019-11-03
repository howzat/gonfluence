package main

import (
	"fmt"
	"gg.gov.revenue.gonfluence/configuration"
	"gg.gov.revenue.gonfluence/files"
	"github.com/spf13/afero"
)

func main() {

	config := configuration.ReadConfiguration("gonfluence.json")
	result, err := files.Search(config, afero.NewReadOnlyFs(afero.NewOsFs()))

	fmt.Printf("config [%v]\n", config)

	if err != nil {
		panic(fmt.Errorf("failed to initialise Gonfluence [%w]", err))
	}

	for _, location := range result.Locations {
		fmt.Printf("found the following files [%s]\n", location)
	}
}
