package main

import (
	"html/template"
	"net/http"
)

var tmpl = `<html>
<head>
    <title>{{ . }}</title>
</head>
<body>
    {{ . }}
    <p>
      <a href="/">main</a> |
      <a href="/view/">view</a>
    </p>

</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("main") //name of the template is main
	t, _ = t.Parse(tmpl)      // parsing of template string
	t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/view/", handler)
	http.HandleFunc("/", index)
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("main") //name of the template is main
	t, _ = t.Parse(tmpl)      // parsing of template string
	t.Execute(w, "Main page")
}
