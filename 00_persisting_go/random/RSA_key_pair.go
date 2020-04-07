package main

import (
	"bytes"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	key, b, err := generatePrivateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

	public2, err := generateSSHPublicKey(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(public2))
}

func generatePrivateKey() (*rsa.PrivateKey, []byte, error) {
	// keyBytes, err := rsa.GenerateKey(rand.Reader, 2048)
	keyBytes, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	certBytes := new(bytes.Buffer)

	err = pem.Encode(certBytes, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(keyBytes)})
	if err != nil {
		return nil, nil, err
	}

	return keyBytes, certBytes.Bytes(), nil
}

func generateSSHPublicKey(key *rsa.PrivateKey) ([]byte, error) {

	public, err := ssh.NewPublicKey(&key.PublicKey) // *rsaPublicKey
	if err != nil {
		return nil, err
	}

	certBytes := new(bytes.Buffer)

	err = pem.Encode(certBytes, &pem.Block{Type: "PUBLIC KEY", Bytes: public.Marshal()})
	if err != nil {
		return nil, err
	}

	return certBytes.Bytes(), nil
}
