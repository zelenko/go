package main

import (
	"./rotate"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {

	rotator("original.jpg", "original_180.jpg", 180)
	rotator("original.jpg", "original_10.jpg", 10)
	rotator("original.jpg", "original_290.jpg", 290)
	rotator("original.jpg", "original_360.jpg", 360)
	rotator("original.jpg", "original_350.jpg", 350)
	rotator("original.jpg", "original_10.jpg", 10)
	rotator("original.jpg", "original_45.jpg", 45)
	rotator("original.jpg", "original_370.jpg", 370)

	/*
		rotator("icon.png", "icon_10.png", 10)
		rotator("icon.png", "icon_110.png", 110)
		rotator("icon.png", "icon_80.png", 80)
		rotator("icon.png", "icon_70.png", 70)
		rotator("icon.png", "icon_45.png", 45)
		rotator("icon.png", "icon_90.png", 90)
		rotator("icon.png", "icon_350.png", 350)
	*/
}

func rotator(source, newFileName string, degree float64) {
	// Open a test image.
	infile, err := os.Open(source)

	if err != nil {
		log.Printf("failed opening %s: ", err)
		panic(err.Error())
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		log.Println("Cannot decode image.")
		panic(err.Error())
	}

	// Rotate
	src = rotate.Rotate(src, degree, color.NRGBA{0, 0, 0, 0})

	// Save new file
	newFile, err := os.Create(newFileName)
	if err != nil {
		log.Printf("failed creating %s: %s", newFile.Name(), err)
		panic(err.Error())
	}
	defer newFile.Close()

	//png.Encode(newFile,src)
	err = jpeg.Encode(newFile, src, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		log.Printf("failed encoding/savid file: %s: ", err)
	}
}
