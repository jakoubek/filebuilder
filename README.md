# filebuilder

**filebuilder** is a litte tool (written in Go) that lets you concatenate a list of source code files in a given order into a *build* file.

I use this myself to create a build out of a bunch of SQL files.

## How to

1. create a text file with all file names - in the order you want them to appear in the build file

```
create_table.sql
insert_dummy_data.sql
create_procedure.sql
```

2. run the program

```bash
filebuilder --filelist=filelist.txt --output=build.sql
```

The resulting *build.sql* file has all the contents of the aforementioned files.


## Installation

```bash
git clone https://github.com/jakoubek/filebuilder.git
cd cmd/filebuilder
go build
```

## Flags

- filelist: name (and - optionally - path) of the filelist file
- wd: working directory (optionally; the path of the filelist file is the default working directory, if none given)
- output: name (and - optionally - path) of the output file (defaults to Stdout if none given)
- force: overwrites the output file if it already exists (otherwise aborts) 
