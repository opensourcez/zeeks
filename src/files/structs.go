package files

import (
	"log"
	"sync"
	"time"
)

var GlobalWaitGroup = sync.WaitGroup{}
var RuntimeConfig *RunConfig

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
	Hits map[int]string
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
	Ignore        []string `json:"ignore"`
	MaxFileSize   int64    `json:"maxFileSize"`
	Configs       []string `json:"configs"`
	Strings       bool     `json:"strings"`
	ParsedConfigs []*SearchConfig
}

type SearchConfig struct {
	Ignore      []string `json:"ignore"`
	MaxFileSize int64    `json:"maxFileSize"`
	String      string   `json:"string"`
	Byte        byte     `json:"byte"`
	Regexp      string   `json:"regexp"`
	// runs the strings command on the file before searching
	Strings bool `json:"strings"`
}
