package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	httpServer := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/img", img)
	http.HandleFunc("/", index)

	// Run the server
	log.Println(httpServer.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, `<a href="/zip">zip</a> | <a href="/zipfolder">zip folder</a>`)
	w.Write([]byte(`<a href="/img">img</a>`))

}

func img(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, `<a href="/zip">zip</a> | <a href="/zipfolder">zip folder</a>`)
	//w.Write([]byte(`<a href="/zip">zip</a> | <a href="/zipfolder">zip folder</a>`))

	imgFile, err := os.Open("original.jpg")
	if err != nil {
		log.Println(err)
	}
	defer imgFile.Close()

	io.Copy(w, imgFile)
}
