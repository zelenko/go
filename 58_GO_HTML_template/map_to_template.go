package main

import (
	"html/template"
	"io"
	"net/http"
)

type HTMLfile struct {
	Folder string
	File   string
}

func (e HTMLfile) Display(wr io.Writer, data interface{}) error {
	t := template.Must(template.ParseFiles(e.Folder + e.File))
	return t.ExecuteTemplate(wr, e.File, data)
}

func main() {
	http.HandleFunc("/test", index)
	http.HandleFunc("/", home)
	println("port 8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]interface{})
	body["one"] = template.HTML("<strong>template</strong>")
	body["two"] = template.HTML("<u>this is a test</u>")
	HTMLfile{"templates/", "test.html"}.Display(w, body)
}

func home(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Message string
		Desc    string
	}

	var rs []Output
	rs = append(rs, Output{"one -", "second field, description ...1"})
	rs = append(rs, Output{"two-", "second field, description ...2"})
	rs = append(rs, Output{"three-", "second field, description ...3"})
	rs = append(rs, Output{"four-", "second field, description ...4"})

	rs4 := make(map[string]string)
	rs4["one"] = "111111"
	rs4["two"] = "22222222222"

	body := make(map[string]interface{})
	body["one"] = "this is one test"
	body["two"] = "two section"
	body["list"] = rs
	rs2 := []string{"one", "two222", "three"}
	body["list2"] = rs2
	body["list3"] = rs4
	HTMLfile{"templates/", "home.html"}.Display(w, body)
}
