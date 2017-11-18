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

// Person fields
type person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

// Results
var results []person

func main() {
	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Collection People
	c := session.DB("test").C("people")

	// Query All
	err = c.Find(bson.M{"name": bson.M{"$ne": "Alex1"}}).Sort("-timestamp").All(&results)
	if err != nil {
		panic(err)
	}

	//for _, v := range Results {
	//fmt.Printf("%s -> %s\n", k, v)
	//t1, e := time.Parse(time.RFC3339,			"2012-08-11T22:08:41+00:00")
	//	p(v.Phone, "\t", v.Timestamp.Format("2006-01-02 3:04PM"), "\t", v.Name, "\t")
	//}
	fmt.Printf("Total Results: %d\n", len(results))

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
	t.Execute(w, struct{ List []person }{results})
}
