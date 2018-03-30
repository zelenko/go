// web interface using Go standard library only
package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/image/", imageHandler)
	log.Fatal(http.ListenAndServe("localhost:80", nil))
}

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

// Index is a data structure used to populate an indexTemplate.
type Index struct {
	Title string
	Body  string
	Links []Link
}

type Link struct {
	URL, Title string
}

// indexHandler is an HTTP handler that serves the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title: "Image gallery",
		Body:  "Welcome to the image gallery.",
	}
	for name, img := range images {
		data.Links = append(data.Links, Link{
			URL:   "/image/" + name,
			Title: img.Title,
		})
	}
	sort.Slice(data.Links, func(i, j int) bool { return data.Links[i].Title < data.Links[j].Title })
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

// imageTemplate is a clone of indexTemplate that provides
// alternate "sidebar" and "content" templates.
var imageTemplate = template.Must(template.Must(indexTemplate.Clone()).Funcs(template.FuncMap{"nav": nav}).ParseFiles("image.tmpl"))

// Image is a data structure used to populate an imageTemplate.
type Image struct {
	Title string
	URL   string
}

// imageHandler is an HTTP handler that serves the image pages.
func imageHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := images[strings.TrimPrefix(r.URL.Path, "/image/")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := imageTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

// images is a map of images.
var images = map[string]*Image{
	"open":        {"Open Source", "http://twiki.org/p/pub/Blog/BlogEntry201207x1/opensource-400.png"},
	"code":        {"Code", "https://www.extremetech.com/wp-content/uploads/2017/07/485120-learn-to-code-640x360.jpg"},
	"programming": {"Programming", "http://fossbytes.com/wp-content/uploads/2016/02/learn-to-code-what-is-programming.jpg"},
	"media":       {"Media", "https://stores-assets.stackcommerce.com/assets/email_capture_modal/email_capture_modal_academy-779c0f5eb45e7367addf95feebaf708bd859e86e0e46c01fac01d90469c77bc9.jpg"},
	"developer":   {"Developer", "https://www.cadcorp.com/images/uploads/product-images/cadcorp_developer_200x200.png"},
}

// nav returns a list of links for navigation section
func nav() []Link {
	var nav []Link

	// populate the slice
	for name, img := range images {
		nav = append(nav, Link{
			URL:   "/image/" + name,
			Title: img.Title,
		})
	}
	// sort the slice
	sort.Slice(nav, func(i, j int) bool { return nav[i].Title < nav[j].Title })
	return nav
}
