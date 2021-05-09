1:(file)BASE64 >> package main
2:(file)BASE64 >> 
3:(file)BASE64 >> import (
8:(file)BASE64 >> 
12:(file)BASE64 >> 	// BuildLinux()
14:(file)BASE64 >> 
19:(file)BASE64 >> 	if err := c.Run(); err != nil {
23:(file)BASE64 >> 
25:(file)BASE64 >> 	os.Setenv("GOOS", "darwin")
28:(file)BASE64 >> 	if err := c.Run(); err != nil {
37:(file)BASE64 >> 	if err := c.Run(); err != nil {
