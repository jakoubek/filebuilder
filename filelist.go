package filebuilder

import (
	"bufio"
	"os"
)

func (app *Application) ReadFilelist() (int, error) {

	file, err := os.Open(app.FilelistFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var filecount int
	var files []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filename := OptionalPathPrepend(scanner.Text(), app.WorkingDirectory)
		files = append(files, filename)
		filecount++
	}
	app.filenames = files
	return filecount, scanner.Err()

}

func (app *Application) ReadFiles() {

}
