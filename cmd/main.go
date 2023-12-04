package main

import (
	"fmt"
	"main/lib"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
)

func main() {
	envMap, e := godotenv.Read()
	if e != nil {
		fmt.Println(e)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		return
	}
	defer watcher.Close()

	folderPath := envMap["FOLDER_PATH"]
	lib.CreateFolder(folderPath)

	err = watcher.Add(folderPath)
	if err != nil {
		fmt.Println("Error adding folder to watcher:", err)
		return
	}

	fmt.Printf("Watching folder: %s\n", folderPath)

	extension := envMap["EXTENSION"]
	lib.WactherFile(watcher, extension)
}
