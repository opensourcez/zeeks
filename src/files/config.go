package files

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

func LoadConfig() {

	RuntimeConfig = new(RunConfig)
	configPath, ok := ArgMap["--config"]
	if !ok {
		log.Println("--config not specified")
		os.Exit(1)
	}

	var dir string
	var fn string
	dir, fn = filepath.Split(configPath)
	log.Println("in config", dir, fn)
	file, err := os.Open(configPath)
	if err != nil {
		log.Println("Could not open config file ...", err)
		os.Exit(1)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		file.Close()
		log.Println("Could not read/parse the config file ...", err)
		os.Exit(1)
	}
	NewConfig := new(RunConfig)
	err = json.Unmarshal(bytes, NewConfig)
	if err != nil {
		file.Close()
		log.Println("Could not read/parse the config file ...", err)
		os.Exit(1)
	}
	if NewConfig.Configs != nil && len(NewConfig.Configs) > 0 {
		RuntimeConfig = NewConfig
	} else {
		NewConfig := new(SearchConfig)
		err = json.Unmarshal(bytes, NewConfig)
		if err != nil {
			file.Close()
			log.Println("Could not read/parse the config file ...", err)
			os.Exit(1)
		}
		RuntimeConfig.ParsedConfigs = append(RuntimeConfig.ParsedConfigs, NewConfig)
	}

	file.Close()
	// parse all search configs
	for _, v := range RuntimeConfig.Configs {

		file, err := os.Open(dir + v + ".json")
		if err != nil {
			log.Println("Could not open config file ...", err)
			os.Exit(1)
		}

		bytes, err := io.ReadAll(file)
		if err != nil {
			log.Println("Could not read/parse the config file ...", err)
			file.Close()
			os.Exit(1)
		}
		NewConfig := new(SearchConfig)
		err = json.Unmarshal(bytes, NewConfig)
		if err != nil {
			file.Close()
			log.Println("Could not read/parse the config file ...", err)
			os.Exit(1)
		}

		for _, v := range NewConfig.Bytes {
			NewConfig.ByteSlice = append(NewConfig.ByteSlice, byte(v))
		}
		RuntimeConfig.ParsedConfigs = append(RuntimeConfig.ParsedConfigs, NewConfig)
		file.Close()
	}

	log.Println(RuntimeConfig)
	for i, v := range RuntimeConfig.ParsedConfigs {
		log.Println(i, v)
	}
	// user, err := user.Current()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// data, err := ioutil.ReadFile(user.HomeDir + "/search.json")
	// if err != nil {
	// 	log.Println("Could not find config file ...")
	// 	os.Exit(1)
	// }

}
