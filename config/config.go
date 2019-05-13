package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	RepUrl string `json: "repUrl"`
	BranchName string `json: "branchName"`
	FileName string `json: "fileName"`
}

func LoadConfiguration (filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}