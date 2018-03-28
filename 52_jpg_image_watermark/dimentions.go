package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {

	w, h, err := getImageSize("icon.png")
	if err != nil {
		fmt.Println("Cannot get size, error: ", err.Error())
	}
	fmt.Print("width: ", w, " height: ", h, "\n")

	w, h, err = getImageSize("original.jpg")
	if err != nil {
		fmt.Println("Cannot get size, error: ", err.Error())
	}
	fmt.Print("width: ", w, " height: ", h, "\n")

	w, h, err = getImageSize("result.jpg")
	if err != nil {
		fmt.Println("Cannot get size, error: ", err.Error())
	}
	fmt.Print("width: ", w, " height: ", h, "\n")

}

func getImageSize(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("Cannot open file.  Error: %s", err.Error())
	}

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, fmt.Errorf("Cannot decode.  Error: %s", err.Error())
	}
	return img.Width, img.Height, nil
}
