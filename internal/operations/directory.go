package operations

import (
	"fmt"
	"os"
	"test-cet-wp-plugin/internal/utilities"
)

/*
Description: Creates a directory base on path with read write execute rights for everyone

Params:

	1: dirPath - directory path to be created.
	2: fileName - ONLY use is this is the ParentDir; otherwise will not work
	3: isParentDir - A bool to let the function now this is the parent directory of the target file
*/
func CreateDirectory(dirPath string, fileName string, isParentDir bool) bool {
	isFileCreated := false
	isDirectoryCreated := false

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Println(dirPath)
		mkdirErr := os.MkdirAll(dirPath, 0777)
		// Error creating parent dir
		if mkdirErr != nil {
			fmt.Println("mkdirErr")
			utilities.HandleFatalError(err, true, "")
		}
		isDirectoryCreated = true
	}

	if isParentDir && isDirectoryCreated {
		//create file
		CreateFile(dirPath, fileName)
	}

	return isFileCreated
}
