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
	parseArguments(os.Args[1:])
	files.LoadConfig()
	rand.Seed(time.Now().UTC().UnixNano())
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
