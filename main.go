package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func Load(structure interface{}) {
	err := readFile(structure, "../configs/config.json")
	if err != nil {
		panic(fmt.Sprintf("Could not load initial configuration: %v", err))
	}
	readFile(structure, "../configs/config.json")
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
	var dir string
	var err error
	if len(os.Args) >= 2 {
		dir, err = filepath.Abs(filepath.Dir(os.Args[1]))
		if err != nil {
			panic(fmt.Sprintf("Could not load given filepath: %v", err))
		}
	}

	if dir == "" {
		//load root folder
		_, filename, _, _ := runtime.Caller(0)
		dir = path.Dir(path.Join(filename, ".."))
	}
	return dir
}