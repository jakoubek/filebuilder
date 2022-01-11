package filebuilder

import (
	"fmt"
	"os"
)

func (app *Application) Output() {

	if app.OutputFile != "" {

		f, err := os.Create(app.OutputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString(app.allfilescontent)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

	} else {
		fmt.Println(app.allfilescontent)
	}

}

func (app *Application) CheckOutput() bool {

	if app.OutputFile != "" {
		if doesFileExist(app.OutputFile) == true && app.force == false {
			return false
		}
		return true
	}
	return true

}
