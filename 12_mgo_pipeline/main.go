package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
	"time"
)

var tmpl = `<html>
<head>
    <title>List</title>
</head>
<body>
    <p>
      <a href="/">main</a> |
      <a href="/view/">view</a>
	</p>
	
	<h1>Members</h1>
	<ul>
		{{range .List}}
				<li>{{.Name}} - {{.Phone}} - {{.ID}} - {{.Timestamp}}</li>
		{{end}}
	</ul>

</body>
</html>
`

//{{.Timestamp}}

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

var Results []Person

func main() {
	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Collection People
	c := session.DB("test").C("people")

	// Query All using pipeline
	pipe := c.Pipe([]bson.M{
		{"$project": bson.M{
			"name":      true, // 1 or true will work, but false is not working
			"_id":       1,
			"phone":     "0", // zero must be in quotations
			"timestamp": 1,
		}},
		{"$match": bson.M{"name": bson.M{"$ne": "Alex1"}}},
		{"$sort": bson.M{"name": 1}},
		{"$limit": 300}},
	)
	err = pipe.All(&Results)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total Results: %d\n", len(Results))

	// start the server
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("main") //name of the template is main
	t, _ = t.Parse(tmpl)      // parsing of template string
	t.Execute(w, struct{ List []Person }{Results})
}
