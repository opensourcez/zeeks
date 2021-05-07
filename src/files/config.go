package files

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func LoadConfig(argMap map[string]string) {

	RuntimeConfig = new(RunConfig)

	for i, v := range argMap {
		if i == "--config" {
			file, err := os.Open(v)
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
		}
		// parse all search configs
		for _, v := range RuntimeConfig.Configs {

			file, err := os.Open(v + ".json")
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
			RuntimeConfig.ParsedConfigs = append(RuntimeConfig.ParsedConfigs, NewConfig)
			file.Close()
		}
	}

	log.Println(RuntimeConfig)
	log.Println(RuntimeConfig.ParsedConfigs[0])
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
