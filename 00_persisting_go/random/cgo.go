package main

// #include <crypt.h>
//
// #cgo LDFLAGS: -L. -lcrypt
//
// char* mycrypt(char* s, char* salt) {
//   return crypt(s, salt);
// }
import "C"
import "fmt"

func main() {
	test()
}

func getHash(username, pass string) string {
	return C.GoString(C.mycrypt(C.CString(pass), C.CString("$1$"+username)))
}

func test() {
	fmt.Println("put:", "$1$CONTRACT$BD2CbhFWoCU4ug9pJ6LR.1")
	fmt.Println("got:", getHash("CONTRACTOR", "test"))
}

// tested on Linux
