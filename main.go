package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	InitFile       = "setup.json"
	AdditionalFile = "config.json"
)

func Load(configPath string, structure interface{}) {
	err := readFile(structure, path.Join(configPath, InitFile))
	if err != nil {
		panic(fmt.Sprintf("Could not load initial configuration: %v", err))
	}
	readFile(structure, path.Join(configPath, AdditionalFile))
}

func readFile(c interface{}, fileName string) (err error) {
	configFile := path.Join(getPath(), fileName)
	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(raw, &c)
	return
}

func getPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Could not load workdir: %v", err))
	}
	return dir
}
