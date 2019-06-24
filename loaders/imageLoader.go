package loaders

import (
	"fmt"
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"os"
)

var imagePath = "./assets/sprites/"
var loadedAssets = map[string]*pixel.Sprite{}

// Loads the picture from file
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

func initMapIfRequired() {
	if loadedAssets == nil {
		loadedAssets = make(map[string]*pixel.Sprite)
	}
}

// Loads the image file and creates a sprite the same size as the original.
func loadSprite(name string) *pixel.Sprite {
	pic, err := loadPicture(name)
	if err != nil {
		fmt.Println("Could not find the file.....")
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	loadedAssets[name] = sprite

	return sprite
}

// Checks if the asset exists in the array
func checkIfLoaded(name string) bool {
	if loadedAssets[name] == nil {
		return false
	} else {
		return true
	}
}

//
func getFromMemory(name string) *pixel.Sprite {
	return loadedAssets[name]
}

//
func GetSprite(name string) *pixel.Sprite {
	initMapIfRequired()

	if checkIfLoaded(name) {
		return getFromMemory(name)
	} else {
		return loadSprite(name)
	}
}