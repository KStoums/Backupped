package utils

import (
	"Tolnkee-Backup-Test/messages"
	"log"
	"os"
)

func RemoveBackupFileIfError(zipFile string) {
	err := os.Remove(zipFile)
	if err != nil {
		log.Fatalln(messages.ERROR_REMOVE_ZIP_FILE)
	}

	return
}
