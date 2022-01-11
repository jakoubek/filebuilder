package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jakoubek/filebuilder"
)

func main() {

	fileindexFile := flag.String("fileindex", "", "path of the fileindex file (mandatory)")
	outputFile := flag.String("output", "", "path of the output file")
	workingDirectory := flag.String("wd", "", "working directory")
	force := flag.Bool("force", false, "force overwriting output file")
	flag.Parse()

	if *fileindexFile == "" {
		fmt.Fprintln(flag.CommandLine.Output(), "Error: no fileindex given!")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	app := filebuilder.NewApplication(
		filebuilder.WithWorkingDirectory(*workingDirectory),
		filebuilder.WithFilelist(*fileindexFile),
		filebuilder.WithOutputFile(*outputFile),
		filebuilder.WithForce(*force),
	)

	log.Println("Filelist file:", app.FileindexFile)
	log.Println("Working directory:", app.WorkingDirectory)
	log.Println("Output to:", app.GetOutput())

	if app.CheckOutput() == false {
		log.Fatalln("Output file already exists - aborting!")
	}

	filesread, err := app.ReadFilelist()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Files to process:", filesread)

	filesprocessed, err := app.ProcessFiles()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Files processed:", filesprocessed)

	app.Output()

}
