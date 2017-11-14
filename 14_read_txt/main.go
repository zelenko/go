package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
    "strings"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func writeLines(lines []string, path string) (err error) {
    var (
        file *os.File
    )

    if file, err = os.Create(path); err != nil {
        return
    }
    defer file.Close()

    //writer := bufio.NewWriter(file)
    for _,item := range lines {
        //fmt.Println(item)
        _, err := file.WriteString(strings.TrimSpace(item) + "\n"); 
        //file.Write([]byte(item)); 
        if err != nil {
            //fmt.Println("debug")
            fmt.Println(err)
            break
        }
    }
    /*content := strings.Join(lines, "\n")
    _, err = writer.WriteString(content)*/
    return
}

func main() {
    lines, err := readLines("foo.txt")
    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
    }
    for _, line := range lines {
        fmt.Println(line)
    }
    //array := []string{"7.0", "8.5", "9.1"}
    err = writeLines(lines, "foo2.txt")
    fmt.Println(err)
}