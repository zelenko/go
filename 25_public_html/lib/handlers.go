package zelenko

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Secure is for https
func Secure(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/v", http.StatusSeeOther)
}

// Test function
func Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("The test function is working. \n"))
}

// Test2 function
func Test2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("The test2 function is working. \n"))
}

// Redirect user from HTTP to same page on HTTPS
func Redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusTemporaryRedirect)
}
