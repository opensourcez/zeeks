package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/opensourcez/zeeks/src/files"
)

func parseArguments(args []string) map[string]string {

	argMap := make(map[string]string)
	argNumber := 0
	for i, v := range args {
		log.Println(i, v)
		split := strings.Split(v, "=")
		if len(split) < 2 {
			log.Println("Argument", v, "is invalid")
			os.Exit(1)
		}
		argMap[split[0]] = split[1]
		argNumber++
	}
	if argNumber == 0 {
		log.Println("you need to specify some arguments")
		os.Exit(1)
	}
	return argMap
}

func main() {

	files.LoadConfig(parseArguments(os.Args[1:]))
	rand.Seed(time.Now().UTC().UnixNano())
	files.InitSearchBuffers()
	files.InitPrintBuffers()
	files.WalkDirectories(".")

	// RUN: post search analyzis

	// Wait group
	files.GlobalWaitGroup.Wait()

	// comments used for searching
	// meow
	// meow
	// meow
	// meow
	// meow
}

func something() (error, chan int) {
	return nil, make(chan int)
}
