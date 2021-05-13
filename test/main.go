package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	root := "./.."
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		log.Println(d)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// err := io.WalkDir(nil, ".", func(path string, d io.DirEntry, err error) error {
	// 	log.Println(d.Info())
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
