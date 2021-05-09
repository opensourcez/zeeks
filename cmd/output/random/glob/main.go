1:(file)BASE64 >> package main
2:(file)BASE64 >> 
3:(file)BASE64 >> import (
6:(file)BASE64 >> 	"path/filepath"
8:(file)BASE64 >> 
10:(file)BASE64 >> 	files, err := filepath.Glob("C:\\Users\\Notandi\\go\\src\\github.com\\zveinn\\golang-lessons-for-beginners\\*")
11:(file)BASE64 >> 	if err != nil {
12:(file)BASE64 >> 		log.Fatal(err)
14:(file)BASE64 >> 	fmt.Println(files) // contains a list of all files in the current directory
