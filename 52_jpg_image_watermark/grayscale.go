package main

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	//"image/png"
	"image/jpeg"
)

func main() {
	grayScale("original.jpg", "original_grayScale2.jpg")
}

func grayScale(filename, newFileName string) {
	infile, err := os.Open(filename)

	if err != nil {
		log.Printf("failed opening %s: %s", filename, err)
		panic(err.Error())
	}
	defer infile.Close()

	imgSrc, _, err := image.Decode(infile)
	if err != nil {
		log.Println("Cannot decode image.")
		panic(err.Error())
	}

	// Create a new grayScale image
	bounds := imgSrc.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := imgSrc.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	// Encode the grayScale image to the new file
	newFile, err := os.Create(newFileName)
	if err != nil {
		log.Printf("failed creating %s: %s", newFile.Name(), err)
		panic(err.Error())
	}
	defer newFile.Close()

	// png.Encode(newFile,grayScale)
	jpeg.Encode(newFile, grayScale, &jpeg.Options{jpeg.DefaultQuality})
}
