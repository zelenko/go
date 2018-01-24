package main

import (
	"./records"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	fmt.Println("HTTP port :3000")
	r := httprouter.New() // methods GET, POST, PUT, PATCH and DELETE

	r.GET("/", index)
	r.GET("/items", records.Index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	r.GET("/items/show", records.Show)
	r.GET("/items/create", records.Create)
	r.POST("/items/create/process", records.CreateProcess)
	r.GET("/items/update", records.Update)
	r.POST("/items/update/process", records.UpdateProcess)
	r.GET("/items/delete/process", records.DeleteProcess)
	http.ListenAndServe(":3000", r)
}

// redirect
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/items", http.StatusSeeOther)
}
