// testing web page without "html/template", but using the "fmt.Fprintf" instead.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var page = map[string]string{

	"head": `<html><body>

<a href="/">page 1</a> |
<a href="/view">view</a> |
<a href="/view/one">view one</a> |
<a href="/view/two">view two</a> |
<a href="/todo/">todo</a> |
<a href="/?next=3">page 3</a> |
<a href="/?next=4">page 4</a> |
<a href="/?next=5">page 5</a>

`,

	"body":   ``,
	"nav":    ``,
	"footer": `</body></html>`,
}

var order = []string{"head", "body", "nav", "footer"}

func handler(w http.ResponseWriter, r *http.Request) {
	page["nav"] = ""
	for i := 0; i < 10; i++ {
		p := strconv.Itoa(i)
		page["nav"] += `<a href="?next=` + p + `">page ` + p + "</a><br>\n"
	}

	page["body"] = `<p>view</p>`
	page["footer"] = `</body></html>`
	fmt.Fprintf(w, buildPage())
}

func index(w http.ResponseWriter, r *http.Request) {
	/*
		next, err := strconv.Atoi(r.FormValue("next"))
		if err != nil {
			fmt.Println(err)
		}
	*/
	page["body"] = `<p>Main template</p>`
	page["footer"] = `the end</body></html>`

	fmt.Fprintf(w, buildPage())
}

func viewOne(w http.ResponseWriter, r *http.Request) {
	//m, err := url.ParseQuery(`x=1&y=2&y=3;z`)

	param := strings.Split(r.URL.Path, "/")
	out := ""
	for _, i := range param {
		out += i
	}
	page["body"] = out
	fmt.Fprintf(w, buildPage())
}

var todo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	page["body"] = `<p>todo</p>`
	page["footer"] = `</body></html>`
	fmt.Fprintf(w, buildPage())
})

// logger tracks every request on stdout
func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func main() {

	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/view/:one", viewOne)
	http.HandleFunc("/view", handler)
	http.HandleFunc("/", index)
	//http.Handle("/todo/", todo)
	http.Handle("/todo/", logger(todo, "todoPage")) // tracks every request on stdout
	s.ListenAndServe()
}

func buildPage() string {
	out := ""
	for i := range order {
		out += page[order[i]]
	}
	return out
}

/*
 notes:
 dirList is the function on line 101 or this page: https://golang.org/src/net/http/fs.go#L705
 that lists directories and files on html page.

 example:
 r.NotFound = http.FileServer(http.Dir("public"))

 func FileServer(root FileSystem) Handler {
	return &fileHandler{root}
 }

 Open(name string) (File, error) // is an interaface


 func serveFile calls dirList function

 ServeFile calls serveFile

 Dir is type; see https://golang.org/src/net/http/fs.go?h=Dir#L40 line 40
 func (d Dir) Open(name string) (File, error) {

 func FileServer(root FileSystem) Handler

 type FileSystem interface {
        Open(name string) (File, error)
}

*/
