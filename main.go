package main

import (
	"fmt"
	"gg.gov.revenue.gonfluence/files"
	"github.com/spf13/afero"
	"os"
)

func main() {

	basePath := os.Getenv("WORKSPACE")
	var fs = afero.NewOsFs()
	result, err := files.Search(basePath, afero.NewReadOnlyFs(fs))

	if err != nil {
		panic(fmt.Errorf("failed to initialise Gonfluence [%w]", err))
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	for _, location := range result.Locations {
		fmt.Printf("found the following files [%s]\n", location)
	}
}
