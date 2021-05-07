package files

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/karrick/godirwalk"
)

var searchBufferMap = make(map[int]chan File)

func InitSearchBuffers() {
	for i := 0; i < 5; i++ {
		searchBufferMap[i] = make(chan File, 5000)
		go processSearchBuffer(i)
	}
}

func processSearchBuffer(index int) {
	// log.Println("Starting search buffer nr:", index)
	for {
		Search(<-searchBufferMap[index])
	}
}

func Search(v File) {

	stat, err := os.Stat(v.Name)
	if err != nil {
		log.Println("could not stat file", v.Name, err)
		GlobalWaitGroup.Done()
		return
	}

	// if the file is 200mb or bigger, we continue
	if stat.Size() > RuntimeConfig.MaxFileSize*1000000 {
		GlobalWaitGroup.Done()
		return
	}

	// ...
	for _, x := range RuntimeConfig.Ignore {
		if strings.Contains(v.Name, x) {
			GlobalWaitGroup.Done()
			return
		}
	}

	file, err := os.Open(v.Name)
	if err != nil {
		log.Println("Can not open file", v.Name, err)
		GlobalWaitGroup.Done()
		return
	}

	scanner := bufio.NewScanner(file)
	var line string
	var foundKeyword bool
	lineNumber := 1
	for scanner.Scan() {
		line = scanner.Text()
		for _, c := range RuntimeConfig.ParsedConfigs {
			if c.String != "" && strings.Contains(line, c.String) {
				foundKeyword = true
				v.Results.Hits[lineNumber] = line
			}
		}

		lineNumber++
	}
	file.Close()

	if foundKeyword {
		// Do not add a waitgrouo done here because
		// we are not done with the file yet.
		printBufferMap[rand.Intn(len(printBufferMap))] <- v
	} else {
		GlobalWaitGroup.Done()
	}
}

func WalkDirectories(dir string) {
	_ = godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(osPathname string, info *godirwalk.Dirent) error {

			if !info.IsDir() {
				GlobalWaitGroup.Add(1)
				searchBufferMap[rand.Intn(len(searchBufferMap))] <- File{
					Name:  osPathname,
					IsDir: info.IsDir(),
					Results: SearchResults{
						Hits: make(map[int]string),
					},
				}
			}

			return nil
		},
		Unsorted: true,
	})

}
