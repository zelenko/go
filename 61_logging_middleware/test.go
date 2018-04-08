// saving logs to text file
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// The syscall package is platform dependent. The online docs are only showing the linux build.
	// fErr, _ := os.OpenFile("Errfile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	// syscall.Dup2(int(fErr.Fd()), 1) /* -- stdout */
	// syscall.Dup2(int(fErr.Fd()), 2) /* -- stderr */

	// Generate log file
	f, err := os.OpenFile("testlogfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/view/:one", viewOne)
	http.HandleFunc("/view", logger2(handler))
	http.HandleFunc("/", logger2(index))
	//http.Handle("/todo/", todo)
	http.Handle("/todo/", logger(todo, "todoPage")) // tracks every request on stdout
	log.Println(s.ListenAndServe())
}

var page = map[string]string{

	"head": `<!doctype html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
</head>
<body>
	<a href="/">page 1</a> |
	<a href="/view">view</a> |
	<a href="/view/one">view one</a> |
	<a href="/view/two">view two</a> |
	<a href="/todo/">todo</a> |
	<a href="/?next=3">page 3</a> |
	<a href="/?next=4">page 4</a> |
	<a href="/?next=5">page 5</a>
`,

	"body": ``,
	"nav":  ``,
	"footer": `
</body>
</html>
`,
}

var order = []string{"head", "body", "nav", "footer"}

// view
func handler(w http.ResponseWriter, r *http.Request) {
	page["nav"] = ""
	for i := 0; i < 10; i++ {
		p := strconv.Itoa(i)
		page["nav"] += `<a href="?next=` + p + `">page ` + p + "</a><br>\n"
	}
	page["body"] = `<p>view</p>`
	page["footer"] = `</body></html>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buildPageBytes())
}

// main page
func index(w http.ResponseWriter, r *http.Request) {
	page["body"] = `<p>Main template</p>`
	page["nav"] = ""
	page["footer"] = `the end</body></html>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buildPageBytes())
}

func viewOne(w http.ResponseWriter, r *http.Request) {
	//m, err := url.ParseQuery(`x=1&y=2&y=3;z`)

	param := strings.Split(r.URL.Path, "/")
	out := ""
	for _, i := range param {
		out += i
	}
	page["body"] = out
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buildPageBytes())
}

var todo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	page["body"] = `<p>todo</p>`
	page["footer"] = `</body></html>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buildPageBytes())
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

func logger2(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//start := time.Now()
		f(w, r)
		//log.Print(r.URL.Path + "\t" + strconv.FormatInt(time.Since(start).Nanoseconds(), 10)) //r.RequestURI
		log.Print(r.UserAgent(), " From: ", r.Referer(), " To: ", r.RequestURI, " IP: "+r.RemoteAddr+"\n")
		for k, v := range r.Header {
			log.Printf("Header field %q, Value %q\n", k, v)
		}
	}
}

func buildPageBytes() []byte {
	out := ""
	for i := range order {
		out += page[order[i]]
	}
	return []byte(out)
}
