2:(file)BASE64 >> 
3:(file)BASE64 >> import (
4:(file)BASE64 >> 	"encoding/json"
8:(file)BASE64 >> 	"path/filepath"
10:(file)BASE64 >> 
12:(file)BASE64 >> 
17:(file)BASE64 >> 		os.Exit(1)
19:(file)BASE64 >> 
25:(file)BASE64 >> 	if err != nil {
26:(file)BASE64 >> 		log.Println("Could not open config file ...", err)
27:(file)BASE64 >> 		os.Exit(1)
29:(file)BASE64 >> 
31:(file)BASE64 >> 	if err != nil {
34:(file)BASE64 >> 		os.Exit(1)
36:(file)BASE64 >> 	NewConfig := new(RunConfig)
38:(file)BASE64 >> 	if err != nil {
41:(file)BASE64 >> 		os.Exit(1)
43:(file)BASE64 >> 	if NewConfig.Configs != nil && len(NewConfig.Configs) > 0 {
46:(file)BASE64 >> 		NewConfig := new(SearchConfig)
47:(file)BASE64 >> 		err = json.Unmarshal(bytes, NewConfig)
55:(file)BASE64 >> 
57:(file)BASE64 >> 	// parse all search configs
59:(file)BASE64 >> 
65:(file)BASE64 >> 
66:(file)BASE64 >> 		bytes, err := io.ReadAll(file)
72:(file)BASE64 >> 		NewConfig := new(SearchConfig)
73:(file)BASE64 >> 		err = json.Unmarshal(bytes, NewConfig)
79:(file)BASE64 >> 
86:(file)BASE64 >> 
88:(file)BASE64 >> 	for i, v := range RuntimeConfig.ParsedConfigs {
93:(file)BASE64 >> 	// 	log.Fatalf(err.Error())
95:(file)BASE64 >> 
101:(file)BASE64 >> 
