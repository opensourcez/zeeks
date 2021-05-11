package files

import (
	"io"
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

	var file File
	var cloneFile *os.File
	for {
		file = <-fileBufferMap[index]
		cloneFile = OpenFile(GetMatchPath(file.OutputPath))
		if cloneFile == nil {
			continue
		}
		for _, v := range file.Results.Hits {
			_, _ = cloneFile.WriteString(v + "\n\n")
		}
		cloneFile.Close()
		GlobalWaitGroup.Done()
	}
}

func MakePath(filePath string) string {
	outDir, ok := ArgMap["--outputDir"]
	if !ok {
		outDir = time.Now().Format("01-02-06-15-04-05")
	}
	dir, fn := filepath.Split(filePath)
	dir = strings.Replace(dir, "../", "", -1)
	err := os.MkdirAll(outDir+"/"+dir, 0777)
	if err != nil {
		return outDir + "/" + dir + fn
	}
	return outDir + "/" + dir + fn
}

func OpenFile(path string) (cloneFile *os.File) {
	var err error
	cloneFile, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		return nil
	}
	return
}

func SaveFile(source *os.File, destination *os.File) {
	n, err := io.Copy(destination, source)
	if err != nil {
		log.Println("could not copy file:", err)
	}
	log.Println(n, "S:", source.Name(), "D:", destination.Name())
}
