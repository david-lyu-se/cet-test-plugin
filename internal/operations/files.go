package operations

import (
	"encoding/json"
	"io"
	"os"
	"test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/utilities"
)

func CreateFile(dirPath string, fileName string) *os.File {
	if fileName == "" {
		fileName = "environments.json"
	}

	if file, err := os.Open(dirPath + fileName); err != nil {
		// If file does not exists, create file
		if newFile, err := os.Create(dirPath + fileName); err != nil {
			utilities.HandleFatalError(err, true, "File "+dirPath+fileName+" not created")
		} else {
			return newFile
		}
	} else {
		return file
	}
	// Should not run. Will error out, have return nil so compiler will be happy
	utilities.HandleFatalError(nil, true, "Something went terrible wrong creating directory: "+dirPath)
	return nil
}

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		utilities.HandleFatalError(err, true, "Could not close File")
	}
}

func ReadFile(file *os.File) *structs.Environments {
	var environments structs.Environments
	byte, err := io.ReadAll(file)
	if err != nil {
		utilities.HandleFatalError(err, true, "Could not convert file to bytes")
	}

	err = json.Unmarshal(byte, &environments)

	if err != nil {
		utilities.HandleFatalError(err, true, "Could not parse file")
	}

	return &environments
}

func WriteFile(file *os.File) {

}
