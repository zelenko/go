package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var i int

// ReadLine reads file faster
func ReadLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		// do something
		//fmt.Println(s)
		doSomething(s)
		line, isPrefix, err = r.ReadLine()
	}
	if isPrefix {
		fmt.Println("buffer size to small")
		return
	}
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

// ReadString reads file slower
func ReadString(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	for err == nil {
		// print line
		//fmt.Print(line)
		doSomething(line)
		line, err = r.ReadString('\n')
	}
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

// Scan reads file OK
func scan(filename string) {
	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		//fmt.Println(i, fileScanner.Text())
		doSomething(fileScanner.Text())
	}
}

func doSomething(input string) {
	if len(input) != 10000 {
		i++
	}
}

func main() {
	filename := `foo3.txt` // Large file

	start := time.Now()
	ReadLine(filename) // faster
	fmt.Println("Time:", time.Since(start), "lines:", i)

	i = 0
	start = time.Now()
	scan(filename) // ok
	fmt.Println("Time:", time.Since(start), "lines:", i)

	i = 0
	start = time.Now()
	ReadLine(filename) // faster
	fmt.Println("Time:", time.Since(start), "lines:", i)

	i = 0
	start = time.Now()
	ReadString(filename) // slow
	fmt.Println("Time:", time.Since(start), "lines:", i)

	i = 0
	start = time.Now()
	ReadLine(filename) // faster
	fmt.Println("Time:", time.Since(start), "lines:", i)

	i = 0
	start = time.Now()
	scan(filename) // ok
	fmt.Println("Time:", time.Since(start), "lines:", i)

}
