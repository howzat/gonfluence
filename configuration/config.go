package configuration

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/afero"
)

type Configuration struct {
	Exclusions []string
	BaseDir    string
}

func ReadConfiguration(filePath string) Configuration {
	var fs = afero.NewOsFs()
	var config, _ = fs.Open(filePath)
	return parseConfig(config)
}

func parseConfig(file afero.File) Configuration {
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Printf(fmt.Sprintf("error: %e", err))
	}

	return conf
}
