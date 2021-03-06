package files

import (
	"bufio"
	"bytes"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var searchBufferMap = make(map[int]chan File)

func InitSearchBuffers() {

	for i := 0; i < RuntimeConfig.Buffers; i++ {
		log.Println("Strating concurrent buffer number:", i)
		searchBufferMap[i] = make(chan File, 100000)
		go processSearchBuffer(i)
	}
}

func processSearchBuffer(index int) {
	duration := time.Duration(RuntimeConfig.Timeout / 2)
	for {
		Process(<-searchBufferMap[index])
		time.Sleep(duration * time.Millisecond)
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

	statFile, err := os.Stat(v.Name)
	if err != nil {
		log.Println("could not stat file", v.Name, err)
		readyToUnlock = true
		return
	}

	if statFile.Size() > RuntimeConfig.MaxFileSize*1000000 {
		readyToUnlock = true
		return
	}

	for _, x := range RuntimeConfig.IgnoreFiles {
		if strings.Contains(v.Name, x) {
			readyToUnlock = true
			return
		}
	}

	v.OutputPath = MakePath(v.Name)
	// if local files are not prefered, then clear the current local file.
	if !RuntimeConfig.PreferLocalFiles {
		log.Println("removing:", v.OutputPath)
		_ = os.Remove(v.OutputPath)
		// Re-make the path to ensure it's existance
		v.OutputPath = MakePath(v.Name)
	}
	localFile = OpenFile(v.OutputPath)
	if localFile == nil {
		// The aboce method handles all the error printing
		// we exit here because this should never happen
		os.Exit(1)
	}

	stat, statErr := localFile.Stat()
	if statErr != nil {
		log.Println("could not stat local file..", localFile.Name())
		// we exit here because this should never happen
		os.Exit(1)
	}

	if RuntimeConfig.PreferLocalFiles && stat.Size() == statFile.Size() {
		log.Println("prefering local file:", localFile.Name())
		// The entire file is already here, no need to open it from the remote drive.
	} else {
		file, err = os.Open(v.Name)
		if err != nil {
			log.Println("Can not open file", v.Name, err)
			readyToUnlock = true
			return
		}
		// Always make a local copy, even if we don't want to keep it
		SaveFile(file, localFile)
		// Close the remote file, it is not needed anymore
		file.Close()
	}

	// Close the local file to flush all buffers
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

	// Open the local file again, we closed it in case we needed to do pre=processing
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

	if len(c.ByteSlice) > 0 && bytes.Contains(lineBytes, c.ByteSlice) {
		v.Results.Hits = append(v.Results.Hits, finalPrefix+line)
		return true
	}
	return false
}

func WalkDirectories(dir string) {

	duration := time.Duration(RuntimeConfig.Timeout)

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		log.Println(d)
		if !d.IsDir() {
			GlobalWaitGroup.Add(1)
			searchBufferMap[rand.Intn(len(searchBufferMap))] <- File{
				Name:    path,
				IsDir:   d.IsDir(),
				Results: SearchResults{},
			}
			time.Sleep(duration * time.Millisecond)
		} else {
			for _, x := range RuntimeConfig.IgnoreFolders {
				if strings.Contains(path, x) {
					return filepath.SkipDir
				}
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

}
