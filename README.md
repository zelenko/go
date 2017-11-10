# Explore Golang features
![Go](https://raw.githubusercontent.com/zelenko/go/master/00_web/pub/img/toby.jpg)

## Go Resorces
* http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
* https://www.youtube.com/playlist?list=PLSak_q1UXfPrI6D67NF8ajfeJ6f7MH83S Learn To Code - Golang Training
* http://exercism.io/languages/go/about
* https://golang.org/pkg/html/template/
* https://www.youtube.com/watch?v=PZTnp8rDnl0 Web Programming with the Go Programming Language

## Download GO Packages
Packages are downloaded into folder specified in the `$GOPATH` system varible:

`
go get -u gopkg.in/mgo.v2
`

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
After [downloading](https://golang.org/dl/) and installing, specify where the code is stored: `set GOPATH=F:\GoCode`

Check existing variables `go env`.  Check current version `go version`.
