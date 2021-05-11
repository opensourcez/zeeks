package files

import (
	"bufio"
	"bytes"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/karrick/godirwalk"
)

var searchBufferMap = make(map[int]chan File)

func InitSearchBuffers() {
	number, err := strconv.Atoi(ArgMap["--concurrent"])
	if err != nil {
		log.Println("--concurrent needs to be a number")
		os.Exit(1)
	}
	// TODO flag to control the number of file buffers
	for i := 0; i < number; i++ {
		log.Println("Strating concurrent buffer number:", i)
		searchBufferMap[i] = make(chan File, 100000)
		go processSearchBuffer(i)
	}
}

func processSearchBuffer(index int) {
	// log.Println("Starting search buffer nr:", index)

	number, err := strconv.Atoi(ArgMap["--timeout"])
	if err != nil {
		log.Println("--timeout needs to be a number")
		os.Exit(1)
	}
	duration := time.Duration(number / 2)
	for {
		time.Sleep(duration * time.Millisecond)
		// TODO enable throttling for checks
		Process(<-searchBufferMap[index])
	}
}

func RunExec(cmd string, value string) string {
	out, err := exec.Command(cmd, value).Output()
	if err != nil {
		return ""
	}
	return string(out)
}

func Process(v File) {
	var file *os.File
	var localFile *os.File
	var readyToUnlock bool
	defer func() {
		if r := recover(); r != nil {
			if file != nil {
				log.Println("error in file", file.Name())
			}
			log.Println("Panic while parsing file", r)
			log.Println(string(debug.Stack()))
		}
		if file != nil {
			file.Close()
		}
		if localFile != nil {
			localFile.Close()
		}
		if readyToUnlock {
			GlobalWaitGroup.Done()
		}
	}()

	stat, err := os.Stat(v.Name)
	if err != nil {
		log.Println("could not stat file", v.Name, err)
		readyToUnlock = true
		return
	}

	if stat.Size() > RuntimeConfig.MaxFileSize*1000000 {
		readyToUnlock = true
		return
	}

	for _, x := range RuntimeConfig.Ignore {
		if strings.Contains(v.Name, x) {
			readyToUnlock = true
			return
		}
	}

	file, err = os.Open(v.Name)
	if err != nil {
		log.Println("Can not open file", v.Name, err)
		readyToUnlock = true
		return
	}

	v.OutputPath = MakePath(v.Name)
	// log.Println(v.OutputPath)
	localFile = OpenFile(v.OutputPath)
	if localFile == nil {
		GlobalWaitGroup.Done()
		os.Exit(1)
	}

	// Always make a local copy, even if we don't want to keep it
	SaveFile(file, localFile)
	// close both file to flush buffers
	// closing the file is also important incase we are running pre processing
	file.Close()
	localFile.Close()

	var foundKeyword = false
	var preProcessing = make(map[string]string)

	// Add the global parse setting
	if RuntimeConfig.Parse != "" {
		preProcessing[RuntimeConfig.Parse] = ""
	}

	// Add per file parsing settings
	for _, c := range RuntimeConfig.ParsedConfigs {
		if c.Parse != "" {
			preProcessing[c.Parse] = ""
		}
	}

	for i, _ := range preProcessing {
		out := RunExec(i, v.OutputPath)
		preProcessing[i] = out
	}

	for _, c := range RuntimeConfig.ParsedConfigs {
		if c.Parse != "" {
			splitOut := strings.Split(preProcessing[c.Parse], "\n")
			for _, x := range splitOut {
				match := FindMatch(c, &v, 0, x, []byte(x), c.Parse)
				if match {
					foundKeyword = true
				}
			}
		}
	}

	// Open the local file again, we closed it incase we needed to do pre=processing
	localFile = OpenFile(v.OutputPath)
	scanner := bufio.NewScanner(localFile)
	var line string
	var lineBytes []byte
	lineNumber := 1

	for scanner.Scan() {
		line = scanner.Text()
		lineBytes = scanner.Bytes()
		for _, c := range RuntimeConfig.ParsedConfigs {
			if c.Parse != "" {
				continue
			}
			match := FindMatch(c, &v, lineNumber, line, lineBytes, "file")
			if match {
				foundKeyword = true
			}

		}
		lineNumber++
	}

	// Cleaning up files incase we don't want them to be saved locally
	if !RuntimeConfig.SaveMatchedFiles && !RuntimeConfig.SaveAllFiles {
		err = os.Remove(localFile.Name())
		if err != nil {
			log.Println(err)
		}
	} else if !foundKeyword && !RuntimeConfig.SaveAllFiles {
		err = os.Remove(localFile.Name())
		if err != nil {
			log.Println(err)
		}
	}

	if foundKeyword {
		fileBufferMap[rand.Intn(len(fileBufferMap))] <- v
	} else {
		readyToUnlock = true
	}
}

func FindMatch(c *SearchConfig, v *File, lineNumber int, line string, lineBytes []byte, extraPrefix string) bool {
	finalPrefix := strconv.Itoa(lineNumber) + ":(" + extraPrefix + ")" + c.Prefix + " >> "
	if c.String != "" && strings.Contains(line, c.String) {
		v.Results.Hits = append(v.Results.Hits, finalPrefix+line)
		return true
	}
	if c.Regexp != "" {
		match, err := regexp.MatchString(c.Regexp, line)
		if match {
			v.Results.Hits = append(v.Results.Hits, finalPrefix+line)
			return true
		} else if err != nil {
			log.Println("REGEXP ERRR:", err)
			return false
		}
	}
	// if v.Name == "main.go" {
	// 	// log.Println("searching:", c.Bytes, c.ByteSlice, lineBytes)
	// }
	if len(c.ByteSlice) > 0 && bytes.Contains(lineBytes, c.ByteSlice) {
		v.Results.Hits = append(v.Results.Hits, finalPrefix+line)
		return true
	}
	return false
}

func WalkDirectories(dir string) {

	number, err := strconv.Atoi(ArgMap["--timeout"])
	if err != nil {
		log.Println("--timeout needs to be a number")
		os.Exit(1)
	}
	duration := time.Duration(number)

	_ = godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(osPathname string, info *godirwalk.Dirent) error {
			time.Sleep(duration * time.Millisecond)

			if !info.IsDir() {
				GlobalWaitGroup.Add(1)
				searchBufferMap[rand.Intn(len(searchBufferMap))] <- File{
					Name:    osPathname,
					IsDir:   info.IsDir(),
					Results: SearchResults{},
				}
			}

			return nil
		},
		Unsorted: true,
	})

}
