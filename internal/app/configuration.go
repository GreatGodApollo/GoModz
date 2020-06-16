package app

import (
	"encoding/json"
	"github.com/GreatGodApollo/GoModz/pkg/api"
	"os"
)

type Configuration struct {
	Token 		string	 `json:"token"`
	Prefixes 	[]string `json:"prefixes"`
	Owners 		[]string `json:"owners"`
}

func LoadConfiguration(file string, log api.Log) Configuration {
	var config Configuration
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	_ = jsonParser.Decode(&config)
	return config
}