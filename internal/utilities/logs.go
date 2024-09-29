package utilities

import (
	"log"
	"os"
)

/*
Description: Handles error and throws error code 1.

Params:

	1: Error message - The error message to print if userMessage is empty
	2: bool to print to terminal
	3: userMessage - Message to say to user, if empty will print err param
*/
func HandleFatalError(err error, shouldPrintErr bool, userMessage string) {

	if shouldPrintErr {
		if userMessage != "" {
			log.Println(userMessage)
		} else {
			log.Fatalln(err)
		}
	}
	os.Exit(1)
}
