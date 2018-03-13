// testing web page without "html/template", but using the "fmt.Fprintf" instead.

package main

import (
	"fmt"
	"net/http"
	//"strconv"
)

var page = map[string]string{

	"head": `<html><body>

<a href="/">page 1</a>
<a href="/view/">page 2</a>
<a href="/?next=3">page 3</a>
<a href="/?next=4">page 4</a>
<a href="/?next=5">page 5</a>

`,

	"body":   ``,
	"nav":    ``,
	"footer": `</body></html>`,
}

var order = []string{"head", "body", "nav", "footer"}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, buildPage())
}

func index(w http.ResponseWriter, r *http.Request) {
	/*
		next, err := strconv.Atoi(r.FormValue("next"))
		if err != nil {
			fmt.Println(err)
		}
	*/
	page["footer"] = `the end</body></html>`

	fmt.Fprintf(w, buildPage())
}

func main() {

	s := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/view", handler)
	http.HandleFunc("/", index)
	s.ListenAndServe()
}

func buildPage() (out string) {

	for i := range order {
		out += page[order[i]]
	}

	return
}
