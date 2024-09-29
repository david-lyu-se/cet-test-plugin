package main

import (
	"log"
	"os"
	"runtime"
	"test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/operations"
	// tea "github.com/charmbracelet/bubbletea"
)

func main() {

	// tea.NewProgram();
	// Use bubble tea here
	log.Println("Checking if config file exists in User's Profile.")

	file := initConfigFile()

	if file != nil {
		//File read writes here
		// var userEnvChoice = 0
		environments := operations.ReadFile(file)

		handleEnvironments(*environments)

		//bubble tea choose environment
		// userEnvChoice = 1

	}

}

func initConfigFile() *os.File {
	//File property variables
	var configRootPath = ""
	var configDir = "cet-wp-plugin/"
	var configFileName = "environments.json"
	var isFileCreated = false
	var file *os.File

	if runtime.GOOS == "windows" {
		configRootPath = ""
	} else {
		configRootPath = "~/"
	}

	// Create config parent dir
	isFileCreated = operations.CreateDirectory(configRootPath+configDir, configFileName, true)
	// Todo: create own func
	if !isFileCreated {
		//check if file exists
		file = operations.CreateFile(configRootPath+configDir, configFileName)
	}
	return file
}

func handleEnvironments(environments structs.Environments) {

}

func initializePlugins() {
	// var plugin = nil

}
