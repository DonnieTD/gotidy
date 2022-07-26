package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	Walk := func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// POINT TO THE FOLDER I NEED TO SORT
		// WALK THROUGH IT AND EXTRACT ALL FILES INTO FOLDERS OF THE SAME EXTENSION

		if !d.IsDir() {
			println("FILE: " + s)
			filePathStrArray := strings.Split(s, "/")
			fileName := filePathStrArray[len(filePathStrArray)-1]
			fileNameStrArray := strings.Split(fileName, ".")
			fileExtension := fileNameStrArray[len(fileNameStrArray)-1]
			// Does the extension folder exist in the output folder?
			_, existErr := os.Stat("./Output/" + fileExtension)
			if existErr != nil {
				// if nah create it
				if err := os.MkdirAll("./Output/"+fileExtension, os.ModePerm); err != nil {
					log.Fatal(err)
				}
			}
			// Move the file into that folder
			oldLocation := "./" + s
			newLocation := "./Output/" + fileExtension + "/" + fileName
			err := os.Rename(oldLocation, newLocation)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	}
	filepath.WalkDir(os.Args[1], Walk)
}
