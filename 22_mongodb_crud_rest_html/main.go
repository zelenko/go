package main

import (
	"./records"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	r.ServeFiles("/static/*filepath", http.Dir("static"))
	r.GET("/", index)
	r.GET("/items", records.Index)
	r.GET("/items/show", records.Show)
	r.GET("/items/create", records.Create)
	r.POST("/items/create/process", records.CreateProcess)
	r.GET("/items/update", records.Update)
	r.POST("/items/update/process", records.UpdateProcess)
	r.GET("/items/delete/process", records.DeleteProcess)
	r.GET("/old", records.OldDesign)
	http.ListenAndServe(":2080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/items", http.StatusSeeOther)
}
