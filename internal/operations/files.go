package operations

import (
	"encoding/json"
	"io"
	"os"
	structures "test-cet-wp-plugin/internal/model/structs"
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
			newFile.WriteString("{ \"Apps\": [],\"WorkingDir\":\"\",\"PluginDir\":\"\" }")

			//file pointer reset
			if _, err := newFile.Seek(0, io.SeekStart); err != nil {
				utilities.HandleFatalError(err, true, "Unable to reset file pointer")
			}

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

func ReadFile(file *os.File) *structures.ConfFile {
	var confFile structures.ConfFile
	byte, err := io.ReadAll(file)
	if err != nil {
		utilities.HandleFatalError(err, true, "Could not convert file to bytes")
	}

	err = json.Unmarshal(byte, &confFile)

	if err != nil {
		utilities.HandleFatalError(err, true, "Could not parse json")
	}

	return &confFile
}

func WriteFile(file *os.File) {

}
