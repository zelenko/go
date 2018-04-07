package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	httpServer := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/zip", zipFile)
	http.HandleFunc("/zipfolder", zipFolder)
	http.HandleFunc("/", index)

	// Run the server
	log.Println(httpServer.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, `<a href="/zip">zip</a> | <a href="/zipfolder">zip folder</a>`)
	w.Write([]byte(`<a href="/zip">zip</a> | <a href="/zipfolder">zip folder</a>`))
}

// download zipfile
func zipFile(w http.ResponseWriter, r *http.Request) {

	// list of files to zip
	files := []string{"IMG_6869.JPG", "IMG_6911.JPG"}

	filename := "images_" + time.Now().Format("2006-03-02_03-04-05pm") + ".zip"
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	zipWriter := zip.NewWriter(w)

	// Add files to zip
	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			log.Fatal(err)
		}

		/*
			header := &zip.FileHeader{
				Name:         a.filename,
				Method:       zip.Store,
				ModifiedTime: uint16(time.Now().UnixNano()),
				ModifiedDate: uint16(time.Now().UnixNano()),
			}
			fw, err := zw.CreateHeader(header)
		*/

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Fatal(err)
		}

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(writer, zipfile)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := zipWriter.Close(); err != nil {
		log.Fatal(err)
	}
}

// download zip folder
func zipFolder(w http.ResponseWriter, r *http.Request) {
	// folder to zip and download
	source := "images"

	info, err := os.Stat(source)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	filename := "images_" + time.Now().Format("2006-03-02_03-04-05pm") + ".zip"
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	archive := zip.NewWriter(w)
	defer archive.Close()

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Println(err)
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			log.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
}
