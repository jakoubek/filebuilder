package filebuilder

import (
	"fmt"
	"log"
	"strings"
)

func (app *Application) ProcessFiles() (int, error) {

	var sb strings.Builder
	var filecount int
	for i, file := range app.filenames {

		fmt.Println(i, "Process file "+file)
		filecontent, err := readFile(file)
		if err != nil {
			log.Println(err)
		} else {
			sb.WriteString(filecontent)
			filecount++
		}
	}
	app.allfilescontent = sb.String()

	return filecount, nil
}
