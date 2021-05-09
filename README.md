# ZEEKS
A tool for searching for keywords, regexp and more inside of large directories. This tool is still a work in progress and anyone that wants to contribue can fork and PR.

### Slow mode
This tool has the ability to slowly walk directories in order not to spike network traffic on network mounted volumes. This option is meant to enabled a stealth sreach.


# Notes
1. Hex is actually string, use string search to find hex
2. IP6REGXP: https://stackoverflow.com/questions/53497/regular-expression-that-matches-valid-ipv6-addresses


# todo
3. SLOW MODE
4. Create network drive and test if file is fetched more then once on multiple open.
4.1. if file is not opened more then once, we can run all kinds of cli things on it. Even in slow mode.
6. Run cli stuff like b64 on matches..


# what to keep for later
1. printing

# what to change
2. Search function needs to take into account all configs.





# What we want
1. quickly parse a large number of files for word matching
1.1. we also need to be able to apply regepx patterns instead of strings.contains
1.2 - we might want to run strings on the file first, depending on the file type
1.3 - we might want to be able to disable running strings
2. we want to output the matches into their new files
3. we want to output the strings return into files as well
4. clone the directory tree we are walking, this will give us an indication of which files we can access
5. Inject into the cloned directory tree how many subdirectories or files a directory has. 

# Config Format
## JSON Runtime Config
``` base-and-token.conf
{
    // Search configs that will be used.
    "configs":    
        [
            "base64",
            "jwt",
        ],

    // Files with these strings in the name will be ignore by the search
    "ignore":[".exe",".gitignore","etc.."]

    // The maximum file size in megabytes
    "maxFileSize":1000,
}
```
## JSON Search Config
``` base64.conf
{
    // Searching for a string
    "string":"string to search for",

    // Searching for a byte sequence
    "byte":[32,24,52,23,255,0],

    // Matching with a regexp
    "regexp":"\\beyJhbGciOi.*\\b",

    // Files to ignore for this specific search
    "ignore":[".exe",".gitignore","etc.."]

    // The maximum file size for this specific search
    "maxFileSize":1000,

    // A prefix that will be added to each match, we recommend making it very descriptive so that you have an easier time understanding the results. 
    "prefix": "SEARCH TAG"
}
```

# flags
```
// only config
$ zeeks --config=[file].conf [directory]

// config with flags
// flags don't overwrite, they add extra checks
$ zeeks --config=[file].conf --strings=memory [directory]
$ zeeks --config=[file].conf --contains=[string] [directory]
$ zeeks --config=[file].conf --bytes=[string] [directory]
$ zeeks --config=[file].conf --regexp=[string] [directory]

// no config
$ zeeks --strings=disk --contains="meow" --regexp="" --bytes=0x10

// Configuring search speed
// --concurrent controls the number of files we can open at a time
// --timeout control the time in MILLISECONDS each concurrent reader will wait between opening files
$ zeeks --concurrent=10 --timeout=200 --config=[file].conf [directory]
```
