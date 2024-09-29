package operations

import (
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

	if _, err := os.Stat(dirPath); os.IsExist(err) {
		mkdirErr := os.Mkdir(dirPath, 0777)
		// Error creating parent dir
		if mkdirErr != nil {
			utilities.HandleFatalError(err, true, "An error creating config"+dirPath)
		}
		isDirectoryCreated = true
	}

	if isParentDir && isDirectoryCreated {
		//create file
		CreateFile(dirPath, fileName)
	}

	return isFileCreated
}
