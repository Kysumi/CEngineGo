package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"os"
)

var imagePath = "./assets/sprites/"

// https://github.com/faiface/pixel/wiki/Moving,-scaling-and-rotating-with-Matrix
// Using example for above
func loadPicture(name string) (pixel.Picture, error) {
	file, err := os.Open(imagePath + name)

	if err != nil {
		return nil, err
	}

	// close the file once function returns.
	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

// Loads the image file and creates a sprite the same size as the original.
func loadSprite(name string) *pixel.Sprite {
	pic, err := loadPicture(name)
	if err != nil {
		fmt.Println("Could not find the file.....")
		panic(err)
	}

	return pixel.NewSprite(pic, pic.Bounds())
}