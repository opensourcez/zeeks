package files

import (
	"fmt"

	"github.com/fatih/color"
)

var printBufferMap = make(map[int]chan File)

func InitPrintBuffers() {
	for i := 0; i < 1; i++ {
		printBufferMap[i] = make(chan File, 50000)
		go processPrintBuffer(i)
	}
}

func processPrintBuffer(index int) {
	// log.Println("Starting print buffer nr:", index)
	var file File
	for {
		file = <-printBufferMap[index]
		// log.Println("WE FOUND THE WORD IN FILE:", file.Name)
		color.Green("FILE: " + file.Name)
		for i, v := range file.Results.Hits {

			fmt.Println(color.GreenString("("+i+"): ") + v)
		}
		GlobalWaitGroup.Done()
	}
}
