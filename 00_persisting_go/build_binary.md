# Building executable
Exporting binary file.

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

## to build on Windows for Linux
* `set GOOS=linux`
* `set GOARCH=amd64`
* `echo %GOARCH%`
* `go build -o web main.go mux.go`

To execute the file on linux in the directory, it needs to be executable, so chmod 744, then run: `./web`
