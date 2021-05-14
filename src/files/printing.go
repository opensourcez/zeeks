package files

import "log"

var printBufferMap = make(map[int]chan File)

func InitPrintBuffers() {
	for i := 0; i < 1; i++ {
		printBufferMap[i] = make(chan File, 50000)
		go processPrintBuffer(i)
	}
}

func processPrintBuffer(index int) {
	var file File
	for {
		file = <-printBufferMap[index]
		log.Println(file)
		GlobalWaitGroup.Done()
	}
}
