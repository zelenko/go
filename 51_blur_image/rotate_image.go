package main

import (
	"./graphics"
	"image"
	"image/jpeg"
	"math"
	"os"
)

func main() {
	imagePath, _ := os.Open("original.jpg")
	defer imagePath.Close()
	srcImage, _, _ := image.Decode(imagePath)

	srcDim := srcImage.Bounds()
	dstImage := image.NewRGBA(image.Rect(0, 0, srcDim.Dy(), srcDim.Dx()))
	graphics.Rotate(dstImage, srcImage, &graphics.RotateOptions{math.Pi / 2.0})

	newImage, _ := os.Create("original_rotated.jpg")
	defer newImage.Close()
	jpeg.Encode(newImage, dstImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
}
