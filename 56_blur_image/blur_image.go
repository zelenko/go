package main

import (
	"./graphics"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	imagePath, _ := os.Open("original.jpg")
	defer imagePath.Close()
	srcImage, _, _ := image.Decode(imagePath)

	dstImage := image.NewRGBA(srcImage.Bounds())
	// Blur Function
	graphics.Blur(dstImage, srcImage, &graphics.BlurOptions{StdDev: 5.5})

	newImage, _ := os.Create("original_blurred.jpg")
	defer newImage.Close()
	jpeg.Encode(newImage, dstImage, &jpeg.Options{jpeg.DefaultQuality})
}
