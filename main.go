package main

import (
	"log"
	"os"
	"test-cet-wp-plugin/internal/operations"
	"test-cet-wp-plugin/internal/tui"
	"test-cet-wp-plugin/internal/utilities"
)

func main() {
	log.Println("Checking if config file exists in User's Profile.")
	file := initConfigFile()

	if file != nil {
		defer file.Close()
		confFile := operations.ReadFile(file)

		log.Println("Environments okay starting tui...")

		tui.StartTea(confFile)

		//Test write file
		// var (
		// 	app  structures.Application = structures.Application{Name: "test", Path: "/", PluginPath: "plugins"}
		// 	apps structures.Applications
		// 	conf structures.ConfFile
		// )

		// apps = append(apps, app)
		// conf = structures.ConfFile{
		// 	Apps:       apps,
		// 	PluginDir:  "/",
		// 	WorkingDir: "/",
		// }
		// operations.WriteFile(file, &conf)
	}
}

func initConfigFile() *os.File {
	//Installing is user home dir
	configRootPath, err := os.UserHomeDir()
	if err != nil {
		utilities.HandleFatalError(err, true, "")
	}

	//File property variables
	var configDir = "/.cet-wp-plugin/"
	var configFileName = "environments.json"
	var isFileCreated = false
	var file *os.File

	// Create config parent dir
	isFileCreated = operations.CreateDirectory(configRootPath+configDir, configFileName, true)
	// Todo: create own func
	if !isFileCreated {
		//check if file exists
		file = operations.CreateFile(configRootPath+configDir, configFileName)
	} else {
		file, err = os.OpenFile(configRootPath+configDir+configFileName, os.O_RDWR, os.ModeAppend)
		if err != nil {
			utilities.HandleFatalError(err, true, "File exists but could not open")
		}
	}
	return file
}
