package operations

import (
	"encoding/json"
	"io"
	"log"
	"os"
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/utilities"
)

func CreateFile(dirPath string, fileName string) *os.File {
	if fileName == "" {
		fileName = "environments.json"
	}

	if file, err := os.OpenFile(dirPath+fileName, os.O_RDWR, os.ModeType); err != nil {
		// If file does not exists, create file
		if newFile, err := os.Create(dirPath + fileName); err != nil {
			utilities.HandleFatalError(err, true, "File "+dirPath+fileName+" not created")
		} else {
			// newFile.WriteString("{ \"Apps\": [],\"WorkingDir\":\"\",\"PluginDir\":\"\" }")

			// //file pointer reset
			// if _, err := newFile.Seek(0, io.SeekStart); err != nil {
			// 	utilities.HandleFatalError(err, true, "Unable to reset file pointer")
			// }
			WriteFile(newFile, nil)

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
		utilities.HandleFatalError(err, true, "")
	}

	return &confFile
}

func WriteFile(file *os.File, conf *structures.ConfFile) {
	var bytes []byte
	var err error
	if conf == nil {
		file.WriteString("{ \"Apps\": [],\"WorkingDir\":\"\",\"PluginDir\":\"\" }")
	} else {
		log.Println("inside write file")
		bytes, err = json.Marshal(conf)
		if err != nil {
			utilities.HandleFatalError(err, true, "Convert to bytes")
			return
		}

		log.Println(string(bytes))
		_, err = file.WriteAt(bytes, io.SeekStart)
		if err != nil {
			utilities.HandleFatalError(err, true, "")
		}
	}

	//file pointer reset
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		utilities.HandleFatalError(err, true, "Unable to reset file pointe")
	}
}
