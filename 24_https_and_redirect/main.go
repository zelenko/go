package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTPS port :443")
	fmt.Println("HTTP port :80")

	r := httprouter.New()
	r.GET("/", secure)

	n := httprouter.New()
	n.GET("/", notSecure)

	//  Start HTTP
	go func() {
		// err := http.ListenAndServe(":80", n)
		err := http.ListenAndServe(":80", http.HandlerFunc(redirect))
		if err != nil {
			log.Fatalln("Web server (HTTP): ", err)
		}
	}()

	//go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	//  Start HTTPS
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", r)
	if err != nil {
		log.Fatal("Web server (HTTPS): ", err)
	}
}

// secure is for https
func secure(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello HTTPS ===========> world.\n"))
}

// notSecure is for HTTP
func notSecure(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Transport layer is NOT secure.\n"))
}

// Redicrect from HTTP to HTTPS
func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusMovedPermanently)
}

// Go to https://localhost:443/ or https://127.0.0.1:443/
// list of TCP ports:
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

// Generate unsigned certificate
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
// for example
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost

// WINDOWS
// windows may have issues with go env GOROOT
// go run %(go env GOROOT)%/src/crypto/tls/generate_cert.go --host=localhost

// instead of go env GOROOT
// you can just use the path to the GO SDK
// wherever it is on your computer

// debian 8 (jessie)
// sudo apt-get install certbot -t jessie-backports

// sudo certbot certonly --webroot -w /var/www/example -d example.com -d www.example.com -w /var/www/thing -d thing.is -d m.thing.is

// to generate self-signed:
// openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem
