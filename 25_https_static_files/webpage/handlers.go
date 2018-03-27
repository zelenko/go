package webpage

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// func "getBuffer" runs before main starts.
func init() {
	getBuffer()
}

// buffer has the HTML content
var buffer = &bytes.Buffer{}

// getBuffer: Open file then copy content to buffer.
func getBuffer() {
	f, err := os.Open(`public\index.gohtml`)
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	defer f.Close()
	_, err = buffer.ReadFrom(f)
	if err != nil {
		log.Fatalln("Error reading from file:", err)
	}
}

// Secure is for https
func Secure(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/v", http.StatusSeeOther)
}

// Test function
func Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("The test function is working. \n"))
}

// List function
func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var tmpl = []byte(`<html>
<head>
<title>List</title>
</head>
<body>
	<h1>List in gohtml</h1>
	<p>
		<a href="/">main</a> |
		<a href="/test/">/test/</a> |
		<a href="/test">/test</a> |
		<a href="index.gohtml">index.gohtml</a> |
		<a href="/html">html</a> |
		
	</p>
</body>
</html>
`)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(tmpl)
}

// HTMLPage displays content from HTML file
func HTMLPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buffer.Bytes())
}

// Test2 function
func Test2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("The test2 function is working. \n"))
}

// Redirect user from HTTP to same page on HTTPS
func Redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusTemporaryRedirect)
}
