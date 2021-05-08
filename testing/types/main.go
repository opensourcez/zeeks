package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/h2non/filetype"
)

func main() {
	buf, _ := ioutil.ReadFile("zeeks")

	kind, _ := filetype.Match(buf)
	log.Println(kind)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
		return
	}

	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
}
