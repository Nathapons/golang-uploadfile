package lib

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func WactherFile(watcher *fsnotify.Watcher, extension string) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				if strings.HasSuffix(event.Name, extension) {
					log.Printf("New %s file detected: %s\n", extension, event.Name)
					// Add your logic to upload data in real-time here
				}
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				if strings.HasSuffix(event.Name, extension) {
					log.Printf("New %s file update: %s\n", extension, event.Name)
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
