# ZEEKS

# todo
1. analyzie the current project and find out where we need to change things
2. start on the most basic version (1. command example)


# what to change
...



# What we want
1. quickly parse a large number of files for word matching
1.1. we also need to be able to apply regepx patterns instead of strings.contains
1.2 - we might want to run strings on the file first, depending on the file type
1.3 - we might want to be able to disable running strings
2. we want to output the matches into their new files
3. we want to output the strings return into files as well
4. .. 


# Formats
## Output files
### strings file pattern
- ./strings/[filename and path]
- ./data/[filename and path]

# Config format
## JSON runtime config
``` base-and-token.conf
{
    "configs":    
        [
            "base64.conf",
            "jwt.conf",
        ],
    "ignore":[".exe",".gitignore","etc.."]
    "maxFileSize":1000,
}
```
## JSON search config
``` base64.conf
{
    "contains":"",
    "byte":0x10,
    "regexp":"",
    "strings":true, // default true
    "ignore":[".exe",".gitignore","etc.."]
    "maxFileSize":1000,
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
```