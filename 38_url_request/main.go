package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"io/ioutil"
)

func main() {
	apiUrl := "https://talaikis.com/api/quotes/random/"
	//resource := "/user/"
	data := url.Values{}
	data.Set("name", "foo")
	data.Add("surname", "bar")

	u, _ := url.ParseRequestURI(apiUrl)
	//u.Path = resource
	urlStr := u.String() // "https://talaikis.com/api/quotes/random/"

	client := &http.Client{}
	//	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // <-- URL-encoded payload
	r, _ := http.NewRequest("GET", urlStr, strings.NewReader(data.Encode())) // <-- URL-encoded payload
	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)

	fmt.Println(resp.ContentLength)

	fmt.Println(resp.Header)

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))
}
