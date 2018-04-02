package main

import "fmt"

func main() {
	reverse1("words")
	reverse1("12")
	reverse1("ok")

	reverse2("words")
	reverse2("12")
	reverse2("ok")

	reverse3("words")
	reverse3("12")
	reverse3("ok")
}

func reverse1(s string) {
	r := []byte(s)
	var n []byte
	for i := len(r) - 1; i >= 0; i-- {
		n = append(n, r[i])
	}
	fmt.Print(s + "\t" + string(n) + "\n")
}

func reverse2(s string) {
	r := ""
	for _, v := range s {
		r = string(v) + r
	}
	fmt.Print(s + "\t" + r + "\n")
}

func reverse3(s string) {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	fmt.Print(s + "\t" + r + "\n")
}
