1:(file)BASE64 >> package main
2:(file)BASE64 >> 
3:(file)BASE64 >> import (
4:(file)BASE64 >> 	"io/ioutil"
7:(file)BASE64 >> 
9:(file)BASE64 >> 	files, err := ioutil.ReadDir("C:\\Users\\Notandi\\go\\src\\github.com\\zveinn\\golang-lessons-for-beginners\\")
10:(file)BASE64 >> 	if err != nil {
11:(file)BASE64 >> 		log.Fatal(err)
13:(file)BASE64 >> 
