package config

import "html/template"

// TPL is all the HTML templates in the "templates" folder
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
