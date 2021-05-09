1:(file)BASE64 >> package main
2:(file)BASE64 >> 
3:(file)BASE64 >> import (
5:(file)BASE64 >> 	"math/rand"
10:(file)BASE64 >> 
13:(file)BASE64 >> 
15:(file)BASE64 >> 
17:(file)BASE64 >> 	parseArguments(os.Args[1:])
18:(file)BASE64 >> 
28:(file)BASE64 >> 
32:(file)BASE64 >> 	// files.InitPrintBuffers()
35:(file)BASE64 >> 
39:(file)BASE64 >> 
41:(file)BASE64 >> 
45:(file)BASE64 >> 		split := strings.Split(v, "=")
53:(file)BASE64 >> 	if argNumber == 0 {
55:(file)BASE64 >> 		os.Exit(1)
57:(file)BASE64 >> 	return files.ArgMap
