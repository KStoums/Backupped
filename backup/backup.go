package backup

import (
	"Tolnkee-Backup-Test/messages"
	"Tolnkee-Backup-Test/utils"
	"archive/zip"
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Backup(pathToBackup string) string {
	fmt.Println(messages.FULL_BACKUP_START)

	zipFile, err := os.Create(fmt.Sprintf("./%d.zip", time.Now().Unix()))
	if err != nil {
		log.Fatalln(messages.ERROR_CREATE_ZIP_FILE, err)
	}

	newWriter := zip.NewWriter(zipFile)

	err = filepath.Walk(pathToBackup, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatalln(messages.ERROR_FILEPATH_WALK, err)
		}

		if info.IsDir() {
			relPath, err := filepath.Rel(pathToBackup, path)
			if err != nil {
				log.Fatalln(messages.ERROR_GET_RELPATH, err)
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				log.Fatalln(messages.ERROR_CREATE_ZIP_HEADER, err)
			}
			header.Name = relPath + "/"
			header.Method = zip.Deflate

			_, err = newWriter.CreateHeader(header)
			if err != nil {
				log.Fatalln(messages.ERROR_WRITE_HEADER, err)
			}
		} else {
			file, err := os.Open(path)
			if err != nil {
				log.Fatalln(messages.ERROR_OPEN_FILE, err)
			}
			defer file.Close()

			fileInfo, err := file.Stat()
			if err != nil {
				log.Fatalln(messages.ERROR_GET_FILE_INFO, err)
			}

			relPath, err := filepath.Rel(pathToBackup, path)
			if err != nil {
				log.Fatalln(messages.ERROR_GET_RELPATH, err)
			}

			header, err := zip.FileInfoHeader(fileInfo)
			if err != nil {
				log.Fatalln(messages.ERROR_CREATE_ZIP_HEADER, err)
			}
			header.Name = relPath
			header.Method = zip.Deflate

			writer, err := newWriter.CreateHeader(header)
			if err != nil {
				log.Fatalln(messages.ERROR_WRITE_HEADER, err)
			}

			bar := progressbar.NewOptions64(fileInfo.Size(),
				progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
				progressbar.OptionEnableColorCodes(true),
				progressbar.OptionShowBytes(true),
				progressbar.OptionSetDescription("Backup in progress"),
				progressbar.OptionFullWidth(),
				progressbar.OptionSetTheme(progressbar.Theme{
					Saucer:        "[yellow]•[reset]",
					SaucerHead:    "[yellow]●[reset]",
					SaucerPadding: " ",
					BarStart:      "[",
					BarEnd:        "]",
				}))

			_, err = io.Copy(io.MultiWriter(writer, bar), file)
			if err != nil {
				log.Fatalln(messages.ERROR_COPY_FILE, err)
			}
			return nil
		}
		return nil
	})
	newWriter.Close()
	zipFile.Close()
	if err != nil {
		log.Fatalln(messages.ERROR_FILEPATH_WALK, err)
	}

	utils.ClearTerminal()
	fmt.Println(messages.FULL_BACKUP_FINISH)
	time.Sleep(3 * time.Second)

	var finalInput string
	for finalInput == "" {
		utils.ClearTerminal()
		fmt.Print(messages.RETURN_MAIN_MENU)

		var input string
		fmt.Scanln(&input)

		if strings.EqualFold(input, "y") || strings.EqualFold(input, "n") {
			finalInput = input
			break
		} else {
			utils.ClearTerminal()
			fmt.Println(messages.ERROR_STRING_CONVERTION_SYNTAX)
			time.Sleep(3 * time.Second)
		}
	}

	return finalInput
}
