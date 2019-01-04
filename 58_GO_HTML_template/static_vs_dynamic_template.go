package main

import (
	"html/template"
	"io"
	"net/http"
)

// TemplateExecutor is of type interface
type TemplateExecutor interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

// DebugTemplateExecutor is
// dynamic
type DebugTemplateExecutor struct {
	Glob string
}

// ExecuteTemplate is
// dynamic
func (e DebugTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t := template.Must(template.ParseGlob(e.Glob))
	return t.ExecuteTemplate(wr, name, data)
}

// ReleaseTemplateExecutor is
// static
type ReleaseTemplateExecutor struct {
	Template *template.Template
}

// ExecuteTemplate is
// static
func (e ReleaseTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return e.Template.ExecuteTemplate(wr, name, data)
}

//const templateGlob = "templates/*.html"
const debug = true

var executor TemplateExecutor

func main() {
	if debug {
		executor = DebugTemplateExecutor{"templates/*.html"}

	} else {
		executor = ReleaseTemplateExecutor{
			template.Must(template.ParseGlob("templates/*.html")),
		}
	}

	http.HandleFunc("/test", index)
	println("port 8080")

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	executor.ExecuteTemplate(w, "test.html", nil)
}
