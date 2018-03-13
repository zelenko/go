package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// slash := string(os.PathSeparator) //or,
	slash := string(filepath.Separator) //to be impervious to OS...
	dirname := "." + slash
	fmt.Printf("dirname is: %s\n", dirname)

	searchFiles(dirname)
}

func searchFiles(dirname string) { // dir is the parent directory you what to search
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filesinfo, err := d.Readdir(-1)
	d.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/* type FileInfo interface {           Name() string       // base name of the file
		       Size() int64        // in bytes for regular files; system-dependent for others
		       Mode() FileMode     // file mode bits
		       ModTime() time.Time // modification time
		       IsDir() bool        // abbreviation for Mode().IsDir()
		       Sys() interface{}   // underlying data source (can return nil)
	       }    */

	for _, fi := range filesinfo {
		if fi.Mode().IsRegular() {
			fmt.Printf("-----\nFileName: %s\nSize: \t%d bytes\n",
				fi.Name(), fi.Size())
		} else {
			fmt.Printf("Name:%s Directory:%t", fi.Name(), fi.Mode().IsDir())
		}
	}
}
