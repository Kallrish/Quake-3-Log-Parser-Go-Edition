package report

import (
	"io/fs"
	"log"
	"os"

	"github.com/Kallrish/quakeLogParser/shared"
)

func checkIfPathAndDirectoryStillExists(path string) {
	if _, pathError := os.Stat(path); os.IsNotExist(pathError) {
		os.MkdirAll(shared.LogFilePath, shared.PATH_CREATION_PERMISSION)
	}
}

func checkIfFileExists(path string) {
	if _, fileError := os.Stat(path); os.IsExist(fileError) {
		os.Remove(path)
	}
}

func createJsonFileInTargetPath(path string, json []byte, permission fs.FileMode) {
	err := os.WriteFile(path, json, permission)
	if err != nil {
		log.Fatal(err)
	}
}
