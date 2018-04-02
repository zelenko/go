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

	// Dimension of new thumbnail 80 X 80
	dstImage := image.NewRGBA(image.Rect(0, 0, 280, 80))
	// Thumbnail function of Graphics
	graphics.Thumbnail(dstImage, srcImage)

	newImage, _ := os.Create("thumbnail_thumbnail.jpg")
	defer newImage.Close()
	jpeg.Encode(newImage, dstImage, &jpeg.Options{jpeg.DefaultQuality})
}
