package main

import (
	"./records"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/items", records.Index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/items/show", records.Show)
	http.HandleFunc("/items/create", records.Create)
	http.HandleFunc("/items/create/process", records.CreateProcess)
	http.HandleFunc("/items/update", records.Update)
	http.HandleFunc("/items/update/process", records.UpdateProcess)
	http.HandleFunc("/items/delete/process", records.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/items", http.StatusSeeOther)
}
