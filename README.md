
# Explore Golang features [![Go Report Card](https://goreportcard.com/badge/zelenko/go)](https://goreportcard.com/report/zelenko/go)
This repository is viewable on [sourcegraph.com](https://sourcegraph.com/github.com/zelenko/go).

![Go](https://raw.githubusercontent.com/zelenko/go/master/37_html_template/pub/img/toby.jpg)

## Go Resorces
* http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
* https://www.youtube.com/playlist?list=PLSak_q1UXfPrI6D67NF8ajfeJ6f7MH83S Learn To Code - Golang Training
* http://exercism.io/languages/go/about
* https://golang.org/pkg/html/template/
* https://www.youtube.com/watch?v=PZTnp8rDnl0 Web Programming with the Go Programming Language
* https://devhints.io/go Go cheatsheet

## Download GO Packages
Packages are downloaded into folder specified in the `$GOPATH` system varible:

`go get -u gopkg.in/mgo.v2`

The 'go get' command requires that the `git` is [installed](https://git-scm.com/download/win) on Windows.

---
## Systemd
`cd /etc/systemd/system/`

`nano golang.service`

```
[Unit]
Description=Go Server

[Service]
ExecStart=/var/www/web
WorkingDirectory=/var/www/
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

### Start the service
* `systemctl enable golang.service`
* `systemctl start golang.service`
* `systemctl status golang.service`
* `systemctl stop golang.service`

## Build on Linux
* `export GOPATH=/var/go/web/`
* `echo $GOPATH`
```
GOOS=linux GOARCH=amd64 go build -o web
```

## Build on Windows
* `set GOARCH=amd64`
* `set GOARCH=386`
* `set GOOS=linux`
* `set GOOS=windows`
* `echo %GOROOT%`
```
go build -o hello.exe hello.go
```
[List of GOOS/GOARCH supported by Go](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63#go-golang-goos-and-goarch)

## Installing GO on Debian
```
cd /usr/local
curl -LO https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz
shasum -a 256 go1.9.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz
```

`vi ~/.profile`
```
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOPATH=/var/go
```
`source ~/.profile`

`echo $GOROOT`

## Installing GO on Windows
After [downloading](https://golang.org/dl/) and installing, specify where the code is stored in system variable.  CLI command: `set GOPATH=F:\GoCode`

Check existing variables `go env`.  Check current version `go version`.


## The proper way to copy a slice
```GO
package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}
	e := make([]string, len(a))
	copy(e, a)
	fmt.Println(e)
}
```
## How to find out the data type?
```GO
// Figure out what type it is: maps, slices, or arrays!

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Declaring local variables
	map1 := map[string]string{"name": "John", "desc": "Golang"}
	map2 := map[string]int{"apple": 23, "tomato": 13}
	slice1 := []int{1, 2, 3}
	array1 := [3]int{1, 2, 3}
	// var m map[string]int
	// m = make(map[string]int)
	// More info Here: https://blog.golang.org/go-maps-in-action

	// Type, such as map[string]string, []int, [3]int
	fmt.Println("map1:", reflect.TypeOf(map1))
	fmt.Println("map2:", reflect.TypeOf(map2))
	fmt.Println("slice1:", reflect.TypeOf(slice1))
	fmt.Println("array1:", reflect.TypeOf(array1))

	// Value, such as map, slice, array.
	fmt.Println("map1:", reflect.ValueOf(map1).Kind())
	fmt.Println("map2:", reflect.ValueOf(map2).Kind())
	fmt.Println("slice1:", reflect.ValueOf(slice1).Kind())
	fmt.Println("array1:", reflect.ValueOf(array1).Kind())

	// True/False statement inside Printf
	fmt.Printf("%v is a map? %v\n", map1, reflect.ValueOf(map1).Kind() == reflect.Map)
	fmt.Printf("%v is a map? %v\n", map2, reflect.ValueOf(map2).Kind() == reflect.Map)
	fmt.Printf("%v is a map? %v\n", slice1, reflect.ValueOf(slice1).Kind() == reflect.Map)

	/*  More about reflect package: https://golang.org/pkg/reflect/
		 	Invalid Kind = iota
	        Bool
	        Int
	        Int8
	        Int16
	        Int32
	        Int64
	        Uint
	        Uint8
	        Uint16
	        Uint32
	        Uint64
	        Uintptr
	        Float32
	        Float64
	        Complex64
	        Complex128
	        Array
	        Chan
	        Func
	        Interface
	        Map
	        Ptr
	        Slice
	        String
	        Struct
	        UnsafePointer
	*/
}
```
![Go](http://farm4.staticflickr.com/3774/11740822616_e435d02a54_o.gif)

### Go Programs and Apps
* **Go Programs and Apps"** [http://go-lang.cat-v.org/go-code](http://go-lang.cat-v.org/go-code)
* **Who is using Go?** https://github.com/golang/go/wiki/GoUsers
* **Stream music using Go:** [nf/goplayer](https://github.com/nf/goplayer/blob/master/player.go)
* **Reading Excel files:** [https://github.com/tealeg/xlsx](https://github.com/tealeg/xlsx)
*  **Generate PDF:** [https://github.com/signintech/gopdf](https://github.com/signintech/gopdf)
* **RSS:** [https://github.com/mmcdole/gofeed](https://github.com/mmcdole/gofeed)
* **File Upload:** [https://github.com/astaxie/build-web-application-with-golang/blob/master/en/04.5.md](https://github.com/astaxie/build-web-application-with-golang/blob/master/en/04.5.md)
* **LDAP:** [https://github.com/mmitton/ldap](https://github.com/mmitton/ldap)
* **Web Frameworks:** [https://golanglibs.com/category/web-framework?sort=top](https://golanglibs.com/category/web-framework?sort=top)
* **Go syntax and features:** [https://github.com/a8m/go-lang-cheat-sheet](https://github.com/a8m/go-lang-cheat-sheet)

# Idiomatic Go
Use _established conventions_ for programming in Go, such as **naming**, **formatting**, **program construction**, and so on, so that programs you write will be _easy_ for other Go programmers to understand.
* https://about.sourcegraph.com/go/idiomatic-go/
* http://idiomaticgo.com/post/testing/idiomatic-go-tests/
* https://pocketgophers.com/idiomatic-go/
* https://golang.org/doc/effective_go.html

## Libraries
Save yourself some time.  Do not reinvent the wheel.  Use existing open source libraries in your projects.
* **Go Libraries:** [https://golanglibs.com/top?q=cms](https://golanglibs.com/top?q=cms)
* **Go Libraries:** [http://go-search.org/](http://go-search.org/search?q=content)
* **Awesome Go:** [https://awesome-go.com/](https://awesome-go.com/)
