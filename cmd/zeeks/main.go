package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/opensourcez/zeeks/src/files"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	parseArguments(os.Args[1:])

	// set argument defaults
	_, ok := files.ArgMap["--outputDir"]
	if !ok {
		files.ArgMap["--outputDir"] = time.Now().Format("01-02-06")
	}

	log.Println("Running with arg map:")
	log.Println(files.ArgMap)
	files.LoadConfig()

	// Default buffer size is 5
	if files.RuntimeConfig.Buffers == 0 {
		files.RuntimeConfig.Buffers = 3
	}

	files.InitSearchBuffers()
	// files.InitPrintBuffers()
	files.InitFileBuffer()
	files.WalkDirectories(files.ArgMap["--dir"])

	// Wait group
	files.GlobalWaitGroup.Wait()
}

func parseArguments(args []string) map[string]string {

	argNumber := 0
	for i, v := range args {
		log.Println(i, v)
		split := strings.Split(v, "=")
		if len(split) < 2 {
			log.Println("Argument", v, "is invalid")
			os.Exit(1)
		}
		files.ArgMap[split[0]] = split[1]
		argNumber++
	}
	if argNumber == 0 {
		log.Println("you need to specify some arguments")
		os.Exit(1)
	}
	return files.ArgMap
}
