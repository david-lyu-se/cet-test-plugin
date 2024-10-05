package main

import (
	"log"
	"os"
	"test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/operations"
	"test-cet-wp-plugin/internal/utilities"
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
	}
	return file
}

func handleEnvironments(environments structs.Environments) {

}

func initializePlugins() {
	// var plugin = nil

}
