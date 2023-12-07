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

			// Detect update file
			if event.Op&fsnotify.Write == fsnotify.Write {
				if strings.HasSuffix(event.Name, extension) {
					log.Printf("New %s file update: %s\n", strings.ToUpper(extension), event.Name)
					// Add your logic to upload data in real-time here
				}
			}

			// Detect create file
			// if event.Op&fsnotify.Create == fsnotify.Create {
			// 	if strings.HasSuffix(event.Name, extension) {
			// 		log.Printf("New %s file detected: %s\n", strings.ToUpper(extension), event.Name)
			// 	}
			// }
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error watching folder:", err)
		}
	}
}
