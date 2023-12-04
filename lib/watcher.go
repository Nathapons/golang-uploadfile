package lib

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func WactherFile(watcher *fsnotify.Watcher) {
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
