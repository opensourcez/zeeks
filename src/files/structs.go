package files

import (
	"log"
	"sync"
	"time"
)

var GlobalWaitGroup = sync.WaitGroup{}
var RuntimeConfig *RunConfig
var ArgMap = make(map[string]string)
var MATCH_POSTFIX = "-matches"

func GetMatchPath(originalPath string) string {
	return originalPath + "-" + MATCH_POSTFIX
}

// File ...
type File struct {
	Name    string
	Dir     string
	IsDir   bool
	ModTime time.Time
	Mode    string
	Size    int64
	Results SearchResults
}
type SearchResults struct {
	Hits []string
}

// FullPath ...
func (f *File) FullPath() string {
	return f.Dir + "\\" + f.Name
}

// Print ...
func (f *File) Print() {
	log.Println("FILE INFO:")
	log.Println(f.Name)
}

// Config ...
type RunConfig struct {
	Ignore           []string `json:"ignore"`
	MaxFileSize      int64    `json:"maxFileSize"`
	Configs          []string `json:"configs"`
	Strings          bool     `json:"strings"`
	ParsedConfigs    []*SearchConfig
	Parse            string `json:"parse"`
	SaveAllFiles     bool   `json:"saveAllFiles"`
	SaveMatchedFiles bool   `json:"saveMatchedFiles"`
}

type SearchConfig struct {
	Ignore      []string `json:"ignore"`
	MaxFileSize int64    `json:"maxFileSize"`
	String      string   `json:"string"`
	Bytes       []int    `json:"bytes"`
	ByteSlice   []byte
	Regexp      string `json:"regexp"`
	Prefix      string `json:"prefix"`
	Parse       string `json:"parse"`
}
