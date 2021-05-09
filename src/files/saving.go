package files

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var fileBufferMap = make(map[int]chan File)

func InitFileBuffer() {
	number, err := strconv.Atoi(ArgMap["--concurrent"])
	if err != nil {
		log.Println("--concurrent needs to be a number")
		os.Exit(1)
	}
	for i := 0; i < number; i++ {
		fileBufferMap[i] = make(chan File, 100000)
		go processFileBuffer(i)
	}
}

func processFileBuffer(index int) {
	log.Println("Starting print buffer nr:", index)
	outDir, ok := ArgMap["--outputDir"]
	if !ok {
		outDir = time.Now().Format("01-02-06-15-04-05")
	}
	var file File
	var err error
	var dir string
	var fn string
	var cloneFile *os.File

	for {
		file = <-fileBufferMap[index]
		dir, fn = filepath.Split(file.Name)
		dir = strings.Replace(dir, "../", "", -1)
		err = os.MkdirAll(outDir+"/"+dir, 0777)
		if err != nil {
			GlobalWaitGroup.Done()
			log.Println(err)
			continue
		}
		log.Println("saving file", outDir+"/"+dir+fn)
		// log.Println(outDir)
		// log.Println(dir)
		// log.Println(fn)
		cloneFile, err = os.OpenFile(outDir+"/"+dir+fn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			GlobalWaitGroup.Done()
			log.Println(err)
			continue
		}

		for _, v := range file.Results.Hits {
			_, _ = cloneFile.WriteString(v + "\n")
		}
		cloneFile.Close()
		GlobalWaitGroup.Done()
	}
}
