package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
)

func wactherFile(watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				if strings.HasSuffix(event.Name, ".csv") {
					log.Printf("New CSV file detected: %s\n", event.Name)
					// Add your logic to upload data in real-time here
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error watching folder:", err)
		}
	}
}

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
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, os.ModePerm)
	}

	err = watcher.Add(folderPath)
	if err != nil {
		fmt.Println("Error adding folder to watcher:", err)
		return
	}

	fmt.Printf("Watching folder: %s\n", folderPath)

	go wactherFile(watcher)
}
