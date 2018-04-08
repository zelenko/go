# Listing directory in HTTP
Taking notes before possible rewrite.

dirList is the function on line 101 or this page: https://golang.org/src/net/http/fs.go#L705
that lists directories and files on html page.

example:

```go
r.NotFound = http.FileServer(http.Dir("public"))

func FileServer(root FileSystem) Handler {
    return &fileHandler{root}
}

Open(name string) (File, error) // is an interface
```

func `serveFile` calls dirList function

`ServeFile` calls `serveFile`

`Dir` is type; see https://golang.org/src/net/http/fs.go?h=Dir#L40 line 40

```go
func (d Dir) Open(name string) (File, error) {

func FileServer(root FileSystem) Handler

type FileSystem interface {
    Open(name string) (File, error)
}
```