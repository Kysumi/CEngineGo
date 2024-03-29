package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type Camera struct {
	GameObject
	matrixPosition pixel.Matrix
	zoom float64
	zoomSpeed float64
	maxZoom float64
	minZoom float64
	chaseObject *Player
}

// Initialize with default values
func (c *Camera) Init() {
	c.pixelPosition = pixel.ZV
	c.zoom = 1.0
	c.zoomSpeed = 1.2
	c.minZoom = 0.4
	c.maxZoom = 5
}

// Sets the window the current camera position
func (c *Camera) Draw(window *pixelgl.Window, deltaTime float64) {
	window.SetMatrix(c.matrixPosition)
}

// Calculates the new position for the camera
func (c *Camera) Update(window *pixelgl.Window) {
	c.pixelPosition.X = c.chaseObject.GetPosition().X - window.Bounds().Max.X / 2
	c.pixelPosition.Y = c.chaseObject.GetPosition().Y - window.Bounds().Max.Y / 2

	c.matrixPosition = pixel.IM.Moved(c.pixelPosition.Scaled(-1))
	window.SetMatrix(c.matrixPosition)

	c.zoom *= math.Pow(c.zoomSpeed, window.MouseScroll().Y)

	if c.zoom > c.maxZoom {
		c.zoom = c.maxZoom
	} else if c.zoom < c.minZoom {
		c.zoom = c.minZoom
	}
}