2:(file)BASE64 >> 
3:(file)BASE64 >> import (
6:(file)BASE64 >> 
7:(file)BASE64 >> var printBufferMap = make(map[int]chan File)
8:(file)BASE64 >> 
11:(file)BASE64 >> 		printBufferMap[i] = make(chan File, 50000)
15:(file)BASE64 >> 
16:(file)BASE64 >> func processPrintBuffer(index int) {
20:(file)BASE64 >> 		file = <-printBufferMap[index]
24:(file)BASE64 >> 
27:(file)BASE64 >> 		GlobalWaitGroup.Done()
