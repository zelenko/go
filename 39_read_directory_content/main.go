package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// main is the entry point for the program.
func main() {
	searchFiles("../")
}

func searchFiles(dir string) { // dir is the parent directory you what to search
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if(!file.IsDir()){
			fmt.Println(" --",file.Name())

		} else{
			fmt.Println(file.Name())
		}
		
	}
}

/*
https://golang.org/pkg/os/#FileInfo

type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}
*/
