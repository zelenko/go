package main

import (
	"fmt"
	"html/template"
	"os"
)

// Output is a type
type Output struct {
	Message string
}

func main() {

	functionMap := template.FuncMap{
		"signature": signature,
		"list":      list,
	}

	//tmpl := template.Must(template.New("body").Funcs(functionMap).Parse(`
	tmpl := template.Must(template.New("simple.html").Funcs(functionMap).ParseFiles("simple.html"))
	tmpl = template.Must(tmpl.New("body").Parse(`

        {{ define "body" }}
           Body
        {{ end }}

		{{ define "navigation" -}}
			one | two | three | four | five
        {{ end }}
     	`))

	tmpl = template.Must(tmpl.New("base").Parse(`
         Start of base template

         {{ template "body" }}

         End of base template
     	`))

	tmpl = template.Must(tmpl.New("baz").Parse(`
         Start of baz template

         {{ template "body" }}

         End of baz template
     	`))
	//tmpl = template.Must(tmpl.New("simple.html").Funcs(functionMap).ParseFiles("simple.html"))
	//tmpl = template.Must(tmpl.New("simple.html").ParseFiles("simple.html"))

	//tmpl.ExecuteTemplate(os.Stdout, "base", nil)
	//tmpl.ExecuteTemplate(os.Stdout, "baz", nil)
	tmpl.ExecuteTemplate(os.Stdout, "simple.html", nil) // works
	//tmpl.Execute(os.Stdout, nil)
}

// Passing single string to template
func signature() string {
	return fmt.Sprintf("www.example.com")
}

// Passing a slice of structures to template as a list
func list() (r []Output) {
	//r := []Output{}
	r = append(r, Output{Message: fmt.Sprint("one")})
	r = append(r, Output{Message: fmt.Sprint("two")})
	r = append(r, Output{Message: fmt.Sprint("three")})
	r = append(r, Output{Message: fmt.Sprint("four")})

	return
}
