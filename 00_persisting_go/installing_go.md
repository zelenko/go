
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

To download all package dependencies:

`go get ./...`

---

## Installing GO on Windows
After [downloading](https://golang.org/dl/) and installing, specify where the code is stored in system variable.  CLI command: `set GOPATH=F:\GoCode`

Check existing variables `go env`.  Check current version `go version`.

# Installing golint
To download it run:
```
go get -u github.com/golang/lint/golint
```

To build it run:
```
go build github.com/golang/lint/golint
```

Put it into `GOROOT/bin` where other executables are.

To use it run:
```
golint ./...
```


## Installing GO on Debian (old version 1.9.2)
These instructions were written when version 1.9.2 came out.

* `cd /usr/local`
* `curl -LO https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz`
* `shasum -a 256 go1.9.2.linux-amd64.tar.gz`
* `tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz`


Update path by editing profile file: `vi ~/.profile`
```
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOPATH=/var/go
```
Reload profile: `source ~/.profile`

Check where the variable points: `echo $GOROOT`
