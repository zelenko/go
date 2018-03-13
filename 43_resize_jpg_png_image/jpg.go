package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	// open image
	file, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(0, 200, img, resize.Lanczos3)

	out, err := os.Create("test_resized_h200.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
