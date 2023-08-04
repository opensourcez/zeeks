# ZEEKS
A tool for searching for keywords, regexp and more in files on network or local drives. 
This tool is still a work in progress and anyone that wants to contribue can fork and PR.

## Big shoutout to Niels who made Yar, wich is where we got most of the regxp rules https://github.com/nielsing/yar

## Slow mode
This tool has the ability to slowly walk directories in order not to spike network traffic on network mounted volumes. This option is meant to enabled a stealth search which is preffered by cyber security professionals.

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
    "ignoreFiles":[".exe", ".gitignore"]
    "ignoreFolders":[".git","vendor"]

    // The maximum file size in MB
    "maxFileSize":1000,
    
    // Save all files to the output directory
    "saveAllFiles": true,

    // Save files to the output directory that match your search criteria
    "saveMatchedFiles": true,
    
    // If you want to search inside the local copy of the file when possible.
    // This only works if the output path matches your previous searches
    "preferLocalFiles": true,

    // How many files do we want to open at a time
    "buffers": 5,

    // How many MILLISECONDS do we want to wait between reading files
    // This timer is applied to each buffer individually
    "timeout": 1000,

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

    // A prefix that will be added to each match, we recommend making it short but descriptive.  
    "prefix": "SEARCH TAG"

    // SEE NOTES ABOUT PARSING IN THE MAIN CONFIG EXAMPLE ABOVE
    "parse": "[tool name]"
}
```
# flags
```
// Running only with a config
$ zeeks --config=[file].conf --outputDir=[directory]  --dir=[directory]
```
