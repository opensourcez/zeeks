2:(file)BASE64 >> 
3:(file)BASE64 >> import (
8:(file)BASE64 >> 
10:(file)BASE64 >> var RuntimeConfig *RunConfig
11:(file)BASE64 >> var ArgMap = make(map[string]string)
12:(file)BASE64 >> 
26:(file)BASE64 >> 
31:(file)BASE64 >> 
32:(file)BASE64 >> // Print ...
33:(file)BASE64 >> func (f *File) Print() {
35:(file)BASE64 >> 	log.Println(f.Name)
37:(file)BASE64 >> 
41:(file)BASE64 >> 	MaxFileSize   int64    `json:"maxFileSize"`
42:(file)BASE64 >> 	Configs       []string `json:"configs"`
43:(file)BASE64 >> 	Strings       bool     `json:"strings"`
45:(file)BASE64 >> 	Parse         string `json:"parse"`
47:(file)BASE64 >> 
52:(file)BASE64 >> 	Bytes       []int    `json:"bytes"`
