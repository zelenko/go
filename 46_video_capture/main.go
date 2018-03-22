// What it does:
//
// This example uses the VideoCapture class to capture frames from a connected webcam,
// and displays the video in a Window class.
//
// How to run:
//
// 		go run ./cmd/capwindow/main.go
//
// +build example

package main

import (
	"fmt"
	"os"
	"strconv"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tcapwindow [camera ID]")
		return
	}

	// parse args
	deviceID, _ := strconv.Atoi(os.Args[1])

	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Capture Window")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("Start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("Error cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}

/*

	more info:
	https://github.com/go-opencv/go-opencv <== Windows insturctions
	https://github.com/go-opencv/go-opencv
	https://gocv.io/writing-code/more-examples/
	https://github.com/hybridgroup/gocv/blob/master/cmd/capwindow/main.go   <== this file

*/
