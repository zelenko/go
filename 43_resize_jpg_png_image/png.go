package main

import (
	"github.com/nfnt/resize"
	"image/png"
	"log"
	"os"
)

func main() {
	// open file
	file, err := os.Open("logo.png")
	if err != nil {
		log.Fatal(err)
	}

	// decode png into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(0, 700, img, resize.Lanczos3)

	out, err := os.Create("logo_h700.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
