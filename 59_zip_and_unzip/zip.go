package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Files to Zip
	files := []string{"IMG_6869.JPG", "IMG_6911.JPG"}
	output := "images.zip"

	err := ZipFiles(output, files)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Zipped File: " + output)
}

// ZipFiles compresses one or many files into a single zip archive file
func ZipFiles(filename string, files []string) error {

	newfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, zipfile)
		if err != nil {
			return err
		}
	}
	return nil
}
