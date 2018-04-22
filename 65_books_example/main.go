package main

import (
	"./config"
	"./records"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	defer config.CloseSession()

	r := httprouter.New()
	r.ServeFiles("/static/*filepath", http.Dir("static"))
	r.GET("/", mainPage)
	r.GET("/items", records.Index)
	r.GET("/items/show", records.Show)
	r.GET("/items/create", records.Create)
	r.POST("/items/create/process", records.CreateProcess)
	r.GET("/items/update", records.Update)
	r.POST("/items/update/process", records.UpdateProcess)
	r.GET("/items/delete/process", records.DeleteProcess)
	r.NotFound = http.FileServer(http.Dir("public"))

	//  Start HTTP
	go func() {
		err := http.ListenAndServe(":80", http.HandlerFunc(redirect))
		if err != nil {
			log.Fatalln("Web server (HTTP): ", err)
		}
	}()

	//  Start HTTPS
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", r)
	if err != nil {
		log.Fatal("Web server (HTTPS): ", err)
	}
}

// redirect used to move traffic from HTTP to same page on HTTPS
func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusTemporaryRedirect)
}

// mainPage is a redirect to the main page
func mainPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/v", http.StatusSeeOther)
}
