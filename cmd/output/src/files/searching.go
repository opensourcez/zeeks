2:(file)BASE64 >> 
3:(file)BASE64 >> import (
4:(file)BASE64 >> 	"bufio"
5:(file)BASE64 >> 	"bytes"
7:(file)BASE64 >> 	"math/rand"
11:(file)BASE64 >> 	"runtime/debug"
15:(file)BASE64 >> 
18:(file)BASE64 >> 
20:(file)BASE64 >> 
22:(file)BASE64 >> 	number, err := strconv.Atoi(ArgMap["--concurrent"])
23:(file)BASE64 >> 	if err != nil {
25:(file)BASE64 >> 		os.Exit(1)
34:(file)BASE64 >> 
36:(file)BASE64 >> 	// log.Println("Starting search buffer nr:", index)
37:(file)BASE64 >> 
39:(file)BASE64 >> 	if err != nil {
41:(file)BASE64 >> 		os.Exit(1)
49:(file)BASE64 >> 
52:(file)BASE64 >> 	if err != nil {
57:(file)BASE64 >> 
65:(file)BASE64 >> 			}
72:(file)BASE64 >> 		if readyToUnlock {
75:(file)BASE64 >> 	}()
76:(file)BASE64 >> 
78:(file)BASE64 >> 	if err != nil {
81:(file)BASE64 >> 		return
83:(file)BASE64 >> 
87:(file)BASE64 >> 		return
89:(file)BASE64 >> 
98:(file)BASE64 >> 	// note: don't forget to disable the file open below if it's a binary..
99:(file)BASE64 >> 
101:(file)BASE64 >> 	var preProcessing = make(map[string]string)
105:(file)BASE64 >> 	for _, c := range RuntimeConfig.ParsedConfigs {
106:(file)BASE64 >> 		if c.Parse != "" {
110:(file)BASE64 >> 
113:(file)BASE64 >> 		preProcessing[i] = out
115:(file)BASE64 >> 
116:(file)BASE64 >> 	for _, c := range RuntimeConfig.ParsedConfigs {
117:(file)BASE64 >> 		if c.Parse != "" {
122:(file)BASE64 >> 					foundKeyword = true
124:(file)BASE64 >> 			}
127:(file)BASE64 >> 
128:(file)BASE64 >> 	file, err = os.Open(v.Name)
129:(file)BASE64 >> 	if err != nil {
132:(file)BASE64 >> 		return
134:(file)BASE64 >> 
136:(file)BASE64 >> 	var line string
138:(file)BASE64 >> 	lineNumber := 1
145:(file)BASE64 >> 				continue
146:(file)BASE64 >> 			}
150:(file)BASE64 >> 			}
154:(file)BASE64 >> 
157:(file)BASE64 >> 		log.Println("..one to buffer")
160:(file)BASE64 >> 		log.Println("unlocking..")
164:(file)BASE64 >> 
167:(file)BASE64 >> 	if c.String != "" && strings.Contains(line, c.String) {
171:(file)BASE64 >> 	if c.Regexp != "" {
173:(file)BASE64 >> 		if match {
174:(file)BASE64 >> 			v.Results.Hits = append(v.Results.Hits, finalPrefix+line)
176:(file)BASE64 >> 		} else if err != nil {
181:(file)BASE64 >> 	// if v.Name == "main.go" {
184:(file)BASE64 >> 	if len(c.ByteSlice) > 0 && bytes.Contains(lineBytes, c.ByteSlice) {
190:(file)BASE64 >> 
192:(file)BASE64 >> 	_ = godirwalk.Walk(dir, &godirwalk.Options{
194:(file)BASE64 >> 
202:(file)BASE64 >> 			}
203:(file)BASE64 >> 
205:(file)BASE64 >> 		},
208:(file)BASE64 >> 
