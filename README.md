[![Go Report Card](https://goreportcard.com/badge/zelenko/go)](https://goreportcard.com/report/zelenko/go)
# Explore Golang features
This repository is viewable on [sourcegraph.com](https://sourcegraph.com/github.com/zelenko/go).

![Go](https://raw.githubusercontent.com/zelenko/go/master/00_web/pub/img/toby.jpg)

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
