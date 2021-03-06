// +build ignore

// Run this command: go run recreate.go info2.txt
// The output will be in the "info2.txt" file.
// This program will check every file in the "samples" folder.

// fields:
// https://github.com/rwcarlsen/goexif/blob/709fab3d192d7c62f86043caff1e7e3fb0f42bd8/exif/fields.go#L52:2
// https://github.com/rwcarlsen/goexif/blob/709fab3d192d7c62f86043caff1e7e3fb0f42bd8/mknote/fields.go
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

func main() {
	flag.Parse()
	fname := flag.Arg(0)

	dst, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	dir, err := os.Open("samples")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	names, err := dir.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	for i, name := range names {
		names[i] = filepath.Join("samples", name)
	}
	makeExpected(names, dst)
}

func makeExpected(files []string, w io.Writer) {
	fmt.Fprintf(w, "package exif\n\n")
	fmt.Fprintf(w, "var regressExpected = map[string]map[FieldName]string{\n")

	for _, name := range files {
		f, err := os.Open(name)
		if err != nil {
			continue
		}

		x, err := exif.Decode(f)
		if err != nil {
			f.Close()
			continue
		}

		fmt.Fprintf(w, "\"%v\": map[FieldName]string{\n", filepath.Base(name))
		x.Walk(&regresswalk{w})
		jdata, _ := x.MarshalJSON()
		fmt.Println(string(jdata))

		//fmt.Println(x.String())
		fmt.Fprintf(w, "},\n\n")
		f.Close()
	}
	fmt.Fprintf(w, "}")
}

type regresswalk struct {
	wr io.Writer
}

func (w *regresswalk) Walk(name exif.FieldName, tag *tiff.Tag) error {
	if strings.HasPrefix(string(name), exif.UnknownPrefix) {
		fmt.Fprintf(w.wr, "\"%v\": `%v`,\n", name, tag.String())
	} else {
		fmt.Fprintf(w.wr, "%v: `%v`,\n", name, tag.String())
	}
	return nil
}
