package filebuilder

import (
	"log"
	"os"
	"path/filepath"
)

type Application struct {
	WorkingDirectory string
	FileindexFile    string
	OutputFile       string
	force            bool
	filenames        []string
	allfilescontent  string
}

type ApplicationConfig func(app *Application)

func NewApplication(opts ...ApplicationConfig) *Application {

	app := Application{
		WorkingDirectory: "",
		FileindexFile:    "",
		OutputFile:       "",
		force:            false,
	}

	for _, opt := range opts {
		opt(&app)
	}

	err := os.Chdir(app.WorkingDirectory)
	if err != nil {
		log.Println(err)
	}

	return &app
}

func WithFilelist(fileindexFile string) ApplicationConfig {
	return func(app *Application) {
		app.FileindexFile = fileindexFile

		if app.WorkingDirectory == "" {
			wd := filepath.Dir(app.FileindexFile)
			app.WorkingDirectory = wd
		}
	}
}

func WithOutputFile(outputFile string) ApplicationConfig {
	return func(app *Application) {
		if outputFile != "" {
			dir, file := filepath.Split(outputFile)
			if dir == "" {
				dir = app.WorkingDirectory
			}
			outputFile = filepath.Join(dir, file)
			app.OutputFile = outputFile
		}
	}
}

func WithWorkingDirectory(workingDirectory string) ApplicationConfig {
	return func(app *Application) {
		if workingDirectory != "" {
			app.WorkingDirectory = workingDirectory
		}
	}
}

func WithForce(force bool) ApplicationConfig {
	return func(app *Application) {
		app.force = force
	}
}

func (app *Application) GetOutput() string {
	if app.OutputFile != "" {
		return app.OutputFile
	}
	return "Stdout"
}
