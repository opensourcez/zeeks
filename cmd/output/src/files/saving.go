2:(file)BASE64 >> 
3:(file)BASE64 >> import (
6:(file)BASE64 >> 	"path/filepath"
11:(file)BASE64 >> 
13:(file)BASE64 >> 
15:(file)BASE64 >> 	number, err := strconv.Atoi(ArgMap["--concurrent"])
16:(file)BASE64 >> 	if err != nil {
18:(file)BASE64 >> 		os.Exit(1)
21:(file)BASE64 >> 		fileBufferMap[i] = make(chan File, 100000)
25:(file)BASE64 >> 
27:(file)BASE64 >> 	log.Println("Starting print buffer nr:", index)
28:(file)BASE64 >> 	outDir, ok := ArgMap["--outputDir"]
37:(file)BASE64 >> 
49:(file)BASE64 >> 		// log.Println(outDir)
51:(file)BASE64 >> 		// log.Println(fn)
58:(file)BASE64 >> 
63:(file)BASE64 >> 		GlobalWaitGroup.Done()
