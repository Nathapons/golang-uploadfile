package lib

import "os"

func CreateFolder(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, os.ModePerm)
	}
}
