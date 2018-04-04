package main

import (
	"fmt"
	"html"
)

func main() {
	a := `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(a))

	b := `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(b))
}
