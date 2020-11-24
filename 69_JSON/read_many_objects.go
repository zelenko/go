package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

var j = []byte(`
{"a":1}
{"a":2}
{"a":3}
`)

func main() {
	dec := json.NewDecoder(bytes.NewBuffer(j))

	for {

		var v interface{}
		err := dec.Decode(&v)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}

		fmt.Println(v)
	}
}
