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
	"gopkg.in/mgo.v2/bson"
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
	defer s.Close()

	fmt.Println("HTTP port :80")

	r := httprouter.New()
	r.GET("/", indexHandler)
	r.GET("/view/", viewHandler)
	r.GET("/list/", listHandler)
	r.GET("/products/", productsHandler)
	r.GET("/aggregate/", aggregateHandler)
	r.GET("/aggregate/:br/", aggregateHandler)
	r.POST("/aggregate/:br/", aggregateHandler)
	r.NotFound = http.FileServer(http.Dir("public"))

	http.ListenAndServe(":80", r)

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
	itemTemplate := `
        {{ define "body" }}
           {{- range listFunction }}
			{{ .Message}}<br />
			{{- end}}
        {{ end }}
     	`

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

// Anything below here requires database connection.  See db.go for connection details

// productHandler is executing html template
func productsHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// Data Type
	type Prod struct {
		// add ID and tags if you need them
		// ID     bson.ObjectId // `json:"id" bson:"_id"`
		Pline    string  // `json:"isbn" bson:"isbn"`
		Bline    string  // `json:"title" bson:"title"`
		Category string  // `json:"author" bson:"author"`
		Price    float32 // `json:"price" bson:"price"`
	}

	// Data
	listFunction := func() []Prod {
		prods := []Prod{{}}

		// same query in mongodb compass:
		//{$and:[{'sales.BR01':{$gt:30000}},{'sales.BR01':{$lt:60000}}]}

		err := Products3.Find(bson.M{
			"$and": []bson.M{
				{"sales.BR01": bson.M{"$gt": 30000}},
				{"sales.BR01": bson.M{"$lt": 60000}},
			},
		}).All(&prods)

		if err != nil {
			return nil
		}
		return prods
	}

	// View
	itemTemplate := `{{ define "body" }}
{{- range listFunction }}
{{ .Pline}} => {{ .Bline}} ==> {{ .Category}} ==> {{ .Price}}<br />
{{- end}}
{{ end }}`

	// Template, alternate way
	t := template.New("template.html")
	t = t.Funcs(template.FuncMap{"listFunction": listFunction})
	t, err := t.ParseFiles("template.html")
	t = template.Must(t, err)

	t = t.New("body")
	t, err = t.Parse(itemTemplate)
	t = template.Must(t, err)

	t.ExecuteTemplate(w, "template.html", "Products")
}

// aggregateHandler is executing html template
func aggregateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	branchID := ps.ByName("br")
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	selector := "Nothing"
	if r.Method == "POST" {
		//selector = r.Form["username"]
		selector = r.PostFormValue("msg")

		//for v := range r.Form["username"] {
		//		selector =+ selector + v
		//	}
	}

	// Data Type
	type ItemData struct {
		// ID     bson.ObjectId // `json:"id" bson:"_id"`
		ID       uint `bson:"_id"` // tags required
		Value    float32
		SalesAll float32 `bson:"sales_all"`
		Branch   string
		MaxSale  float32 `bson:"maxSale"`
	}
	// Sample output from mongodb shell:
	// { "_id" : 42973, "value" : 11.36, "sales_all" : 598.23, "branch" : "BR19", "maxSale" : 303.96 }

	// SomeProducts records from products3 collection.  Get the slice of TYPES!
	listFunction := func() []ItemData {
		/*	Pipe command for mongodb shell:
			db.products3.aggregate([
			  {$match: {'onhand.BR03':{"$gt":0},'sales.BR03':{'$lt':2}, sales_all:{'$gt':0}}},
			  {$unwind: '$branch'},
			  {$project: {"_id":1, "value":"$onhand_value.BR03", "sales_all":1,"sales":"$branch.sales","branch":"$branch.id"}},
			  {$sort:{'value':-1,'sales':-1}},
			  {$group: {_id:'$_id', value: { $first: "$value" }, sales_all: { $first: "$sales_all" }, branch: { $first: "$branch" }, "maxSale": { $max: '$sales'} }},
			  {$match: {'branch':'BR19'}},
			  {$sort:{'value':-1}},
			  {$limit: 300}
			])
		*/

		pipeline := []bson.M{
			{
				"$match": bson.M{
					"onhand.BR03": bson.M{"$gt": 0},
					"sales.BR03":  bson.M{"$lt": 2},
					"sales_all":   bson.M{"$gt": 0},
				},
			},
			{
				"$unwind": "$branch",
			},
			{
				"$project": bson.M{
					"_id":       1,
					"value":     "$onhand_value.BR03",
					"sales_all": 1,
					"sales":     "$branch.sales",
					"branch":    "$branch.id",
				},
			},
			{
				"$sort": bson.M{"value": -1, "sales": -1},
			},
			{
				"$group": bson.M{
					"_id":       "$_id",
					"value":     bson.M{"$first": "$value"},
					"sales_all": bson.M{"$first": "$sales_all"},
					"branch":    bson.M{"$first": "$branch"},
					"maxSale":   bson.M{"$max": "$sales"},
				},
			},
			{
				"$match": bson.M{"branch": branchID}, // "BR19"
			},
			{
				"$sort": bson.M{"value": -1},
			},
			{
				"$limit": 30,
			},
		}

		pipe := Products3.Pipe(pipeline)
		prods := []ItemData{{}} // Declaring type
		err := pipe.All(&prods) // Get results into memory
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return prods
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

<form action="/aggregate/BR05/" method="POST">
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
	{{ .ID}} => {{ .Value}} ==> {{ .SalesAll}} ==> {{ .Branch}} ==> {{ .MaxSale}}<br />
	{{- end}}
{{- end}}`

	// Template
	t := template.Must(template.New("template.html").Funcs(template.FuncMap{"listFunction": listFunction}).
		ParseFiles("template.html"))
	t = template.Must(t.New("body").Parse(itemTemplate))
	t.ExecuteTemplate(w, "template.html", "Aggregation: "+branchID+" "+selector)
}
