package main

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

func main() {
	for i := 0; i < 5; i++ {
		println(randomGenerator())
	}

	for i := 0; i < 61; i++ {
		print(randomNumber(), " ")
	}
}

func randomGenerator() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func randomString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	// fmt.Println(b)
	return base64.StdEncoding.EncodeToString(b)
}

func randomNumber() string {
	min := big.NewInt(1)
	max := big.NewInt(5)
	nBig, _ := rand.Int(rand.Reader, max)
	nBig = nBig.Add(min, nBig)
	return nBig.String()
}
