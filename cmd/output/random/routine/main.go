1:(file)BASE64 >> package main
2:(file)BASE64 >> 
3:(file)BASE64 >> import (
8:(file)BASE64 >> 
9:(file)BASE64 >> var buffer1 = make(chan int, 31)
10:(file)BASE64 >> var buffer2 = make(chan int, 31)
11:(file)BASE64 >> 
15:(file)BASE64 >> 		go putInbuffer1(i)
17:(file)BASE64 >> 
25:(file)BASE64 >> 
30:(file)BASE64 >> 		fmt.Println("no message sent")
32:(file)BASE64 >> 
34:(file)BASE64 >> 
35:(file)BASE64 >> func process() {
46:(file)BASE64 >> 
