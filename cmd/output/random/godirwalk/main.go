1:(file)BASE64 >> package main
2:(file)BASE64 >> 
3:(file)BASE64 >> import (
6:(file)BASE64 >> 
9:(file)BASE64 >> 
13:(file)BASE64 >> 			if strings.Contains(osPathname, ".git") {
14:(file)BASE64 >> 				return errors.New("skipping: " + osPathname)
15:(file)BASE64 >> 			}
17:(file)BASE64 >> 		},
19:(file)BASE64 >> 			if err.Error() == "path filter" {
21:(file)BASE64 >> 			}
22:(file)BASE64 >> 			return godirwalk.Halt
23:(file)BASE64 >> 		},
