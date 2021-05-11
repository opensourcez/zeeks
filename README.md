# ZEEKS
A tool for searching for keywords, regexp and more inside of large directories. This tool is still a work in progress and anyone that wants to contribue can fork and PR.

## Slow mode
This tool has the ability to slowly walk directories in order not to spike network traffic on network mounted volumes. This option is meant to enabled a stealth sreach.

## Concurrency
....

## 

# Notes
2. IP6REGXP: https://stackoverflow.com/questions/53497/regular-expression-that-matches-valid-ipv6-addresses


# todo
4. Create network drive and test if file is fetched more then once on multiple open.
4.1. if file is not opened more then once, we can run all kinds of cli things on it. Even in slow mode.
6. Run cli stuff like b64 on matches..
7. move meta data to sqlite ? https://github.com/volatiletech/sqlboiler

# Config Format
## JSON Runtime Config
This is the main configuration file that will be referenced in the --config flag
``` RUNTIME.conf
{
    // Search configs that will be used.
    // THESE CONFIGS SHOULD HAVE A RELATIVE PATH FROM THE CONFIG FILE
    "configs":    
        [
            "Example",
            "jwt",
        ],

    // Files with these strings in the name will be ignore by the search
    "ignore":[".exe", ".gitignore", ".git", "etc...]

    // The maximum file size in MB
    "maxFileSize":1000,
    
    // If you want to search inside the local copy of the file when possible.
    // This only works if the output path matches your previous searches
    "preferLocalFiles": true,

     // Pre parse the file with a cli.
     // This parsing step will REPLACE the normal step where the file is opened and read line by line.
     // can run any tool from the command line that does not require arguments and outputs text
     // f.x: strings, hd, etc..
    "parse": "[tool name]"
}
```
## JSON Search Config
This kind of config file is a "search config" one or more of these configs can be listed in the MAIN config file which is shown above this example.
``` Example.conf
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

    // A prefix that will be added to each match, we recommend making it short but descriptive.  
    "prefix": "SEARCH TAG"

    // SEE NOTES ABOUT PARSING IN THE MAIN CONFIG EXAMPLE ABOVE
    "parse": "[tool name]"
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
