// One page MGO aggregation
//
// Copyright (c) 2018 - Zelenko <https://github.com/zelenko>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

// itemTemplate can be overridden in each function locally
var itemTemplate = `
        {{ define "body" }}
           {{- range listFunction }}
			{{ .Message}}<br />
			{{- end}}
        {{ end }}
     	`

// main is the entry point for the program.
func main() {

	fmt.Println("HTTP port :80")

	r := httprouter.New()
	r.GET("/", indexHandler)
	r.GET("/view/", viewHandler)
	r.GET("/list/", listHandler)
	r.GET("/products/", postHandler)
	r.POST("/products/", postHandler)
	r.NotFound = http.FileServer(http.Dir("public"))

	http.ListenAndServe(":80", r)

}

// postHandler uses append to populate slice
func postHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Data Type
	type Output struct {
		Message string
	}

	selector := r.FormValue("from")
	selector2 := r.FormValue("to")
	method := r.Method

	options := []string{"BR01", "BR02", "BR03", "BR04", "BR05", "BR06"}

	if selector == "" {
		selector = options[0]
	}
	if selector2 == "" {
		selector2 = options[1]
	}

	// Data
	listFunction := func() (r []Output) {
		r = append(r, Output{Message: "selector: " + selector})
		r = append(r, Output{Message: "selector2: " + selector2})
		r = append(r, Output{Message: "method: " + method})
		r = append(r, Output{Message: "three"})
		r = append(r, Output{Message: "four"})
		return
	}

	// pass functions (passing data)
	functionMap := template.FuncMap{
		"listFunction": listFunction,
		//"listFunction3": listFunction3,
	}

	// View
	itemTemplate := `{{ define "body" }}
<form name="form1" action="/products/" method="POST">
<p>From: ` +

		generateSelector("from", options, selector, selector2) +
		`To: ` +
		generateSelector("to", options, selector2, selector) +

		`</form>
</p>

   {{- range listFunction }}
	{{ .Message}}<br />
	{{- end}}
{{- end}}`

	// Template
	t := template.Must(template.New("template.html").Funcs(functionMap).ParseFiles("template.html"))
	t = template.Must(t.New("body").Parse(itemTemplate))
	t.ExecuteTemplate(w, "template.html", "Post")
}

// viewHandler uses append to populate slice
func viewHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// Data Type
	type Output struct {
		Message string
	}

	// Data
	listFunction := func() (r []Output) {
		r = append(r, Output{Message: fmt.Sprint("one")})
		r = append(r, Output{Message: fmt.Sprint("two")})
		r = append(r, Output{Message: fmt.Sprint("three")})
		r = append(r, Output{Message: fmt.Sprint("four")})
		return
	}

	// pass functions (passing data)
	functionMap := template.FuncMap{
		"listFunction": listFunction,
		//"listFunction3": listFunction3,
	}

	// View
	itemTemplate := `{{ define "body" }}
<p>
<a href="/aggregate/BR01/">01</a> |
<a href="/aggregate/BR02/">02</a> |
<a href="/aggregate/BR03/">03</a> |
<a href="/aggregate/BR04/">04</a> |
<a href="/aggregate/BR05/">05</a> |
<a href="/aggregate/BR06/">06</a> |

<form name="form1" action="/products/" method="POST">
<select name="t1" onchange="this.form.submit()">
<option value="ACCOUNTING">ACCOUNTING</option>
<option value="ACCOUNTING.AP">ACCOUNTING.AP</option>
<option value="OP.REG.MGR">OP.REG.MGR</option>
<option value="OPERATIONS.MGR">OPERATIONS.MGR</option>
<option value="PURCHASING">PURCHASING</option>
<option value="RECEIVING">RECEIVING</option>
<option value="RF.MANAGER">RF.MANAGER</option>
<option value="SALES.ACCOUNTING">SALES.ACCOUNTIN</option>
<option value="SUPER SALES">SUPER SALES</option>
<option value="TEST.OPR.M">TEST.OPR.M</option>
</select>
</form>

</p>


   {{- range listFunction }}
	{{ .Message}}<br />
	{{- end}} }}
{{- end}}`

	// Template
	t := template.Must(template.New("template.html").Funcs(functionMap).ParseFiles("template.html"))
	t = template.Must(t.New("body").Parse(itemTemplate))
	t.ExecuteTemplate(w, "template.html", "View")

}

// indexHandler uses append to populate slice
func indexHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// Data Type
	type Output struct {
		Message string
	}

	// Data
	listFunction := func() (r []Output) {
		r = append(r, Output{Message: fmt.Sprint("one -")})
		r = append(r, Output{Message: fmt.Sprint("two-")})
		r = append(r, Output{Message: fmt.Sprint("three-")})
		r = append(r, Output{Message: fmt.Sprint("four-")})
		return
	}

	// pass functions (passing data)
	functionMap := template.FuncMap{
		"listFunction": listFunction,
		//"listFunction3": listFunction3,
	}

	// Template
	t := template.Must(template.New("template.html").Funcs(functionMap).ParseFiles("template.html"))
	t = template.Must(t.New("body").Parse(itemTemplate))
	t.ExecuteTemplate(w, "template.html", "Index")
}

// listHandler is executing html template
func listHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// Data Type
	type branch struct {
		Message string
		Note    string
	}

	// Data
	// listFunction returns slice of branches.  Branches are declared by a literal.
	listFunction := func() []branch {

		// r is a slice of branch data types
		r := []branch{
			{Message: "column[2]", Note: "This is the note!"},
			{Message: "column[3]", Note: "This is the note!"},
			{Message: "column[1]", Note: "This is the note!"},
			{Message: "column[4]", Note: "This is the note!"},
			{Message: "column[5]", Note: "This is the note!"},
			{Message: "column[6]", Note: "This is the note!"},
			{Message: "column[8]", Note: "This is the note!"},
			{Message: "column[9]", Note: "This is the note!"},
			{Message: "column[7]", Note: "This is the note!"},
		}

		return r
	}

	// pass functions (passing data)
	functionMap := template.FuncMap{
		"listFunction": listFunction,
		//"listFunction3": listFunction3,
	}

	// View
	itemTemplate := `
{{- define "body" }}
	{{- range listFunction }}
	{{ .Message}} => {{ .Note}}<br />
	{{- end}}
 {{ end }}`

	// Template
	t := template.Must(template.New("template.html").Funcs(functionMap).ParseFiles("template.html"))
	t = template.Must(t.New("body").Parse(itemTemplate))
	t.ExecuteTemplate(w, "template.html", "List")
}

// generateSelector creates HTML for the drop down selection.
func generateSelector(name string, options []string, current string, current2 string) (output string) {
	output = "\n" + `<select name="` + name + `" onchange="this.form.submit()">` + "\n"
	for _, name := range options {
		if name == current2 {
			// skip this this
		} else if name != current {
			output += `	<option value="` + name + `">` + name + `</option>` + "\n"
		} else {
			output += `	<option value="` + name + `" selected>` + name + `</option>` + "\n"
		}
	}
	output += "</select>\n"
	return
}
