package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Creates a pixel window using static config.
// TODO Make config get loaded from a file.
func createWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "CEngine2",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		fmt.Println("Failed to create new window")
		panic(err)
	}

	return win
}