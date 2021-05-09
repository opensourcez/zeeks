3:(file)BASE64 >> 
5:(file)BASE64 >> This tool has the ability to slowly walk directories in order not to spike network traffic on network mounted volumes. This option is meant to enabled a stealth sreach.
6:(file)BASE64 >> 
7:(file)BASE64 >> 
9:(file)BASE64 >> 1. Hex is actually string, use string search to find hex
11:(file)BASE64 >> 
12:(file)BASE64 >> 
14:(file)DRIVE >> 4. Create network drive and test if file is fetched more then once on multiple open.
14:(file)BASE64 >> 4. Create network drive and test if file is fetched more then once on multiple open.
17:(file)BASE64 >> 
18:(file)BASE64 >> 
19:(file)BASE64 >> # what to keep for later
21:(file)BASE64 >> 
22:(file)BASE64 >> # what to change
24:(file)BASE64 >> 
25:(file)BASE64 >> 
26:(file)BASE64 >> 
27:(file)BASE64 >> 
28:(file)BASE64 >> 
32:(file)BASE64 >> 1.2 - we might want to run strings on the file first, depending on the file type
36:(file)BASE64 >> 4. clone the directory tree we are walking, this will give us an indication of which files we can access
38:(file)BASE64 >> 
43:(file)BASE64 >>     // Search configs that will be used.
49:(file)BASE64 >> 
50:(file)BASE64 >>     // Files with these strings in the name will be ignore by the search
52:(file)BASE64 >> 
55:(file)BASE64 >> 
64:(file)BASE64 >>     "string":"string to search for",
65:(file)BASE64 >> 
66:(file)BASE64 >>     // Searching for a byte sequence
68:(file)BASE64 >> 
70:(file)JWT >>     "regexp":"\\beyJhbGciOi.*\\b",
71:(file)BASE64 >> 
74:(file)BASE64 >> 
77:(file)BASE64 >> 
80:(file)BASE64 >> 
81:(file)BASE64 >>     // Pre parse the file with a cli
88:(file)BASE64 >> 
92:(file)BASE64 >> $ zeeks --config=[file].conf [directory]
93:(file)BASE64 >> 
94:(file)BASE64 >> // config with flags
97:(file)BASE64 >> $ zeeks --config=[file].conf --contains=[string] [directory]
100:(file)BASE64 >> 
101:(file)BASE64 >> // no config
102:(file)MEOW >> $ zeeks --strings=disk --contains="meow" --regexp="" --bytes=0x10
103:(file)BASE64 >> 
106:(file)BASE64 >> // --timeout control the time in MILLISECONDS each concurrent reader will wait between opening files
