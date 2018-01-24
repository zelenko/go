package config

import "html/template"

// TPL are all the html templates in the "templates" folder
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
