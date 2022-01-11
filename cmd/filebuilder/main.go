package main

import (
	"flag"
	"log"

	"github.com/jakoubek/filebuilder"
)

func main() {

	filelistFile := flag.String("filelist", "", "path of the filelist file")
	outputFile := flag.String("output", "", "path of the output file")
	workingDirectory := flag.String("wd", "", "working directory")
	force := flag.Bool("force", false, "force overwriting output file")
	flag.Parse()

	app := filebuilder.NewApplication(
		filebuilder.WithWorkingDirectory(*workingDirectory),
		filebuilder.WithFilelist(*filelistFile),
		filebuilder.WithOutputFile(*outputFile),
		filebuilder.WithForce(*force),
	)

	log.Println("Filelist file:", app.FilelistFile)
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
