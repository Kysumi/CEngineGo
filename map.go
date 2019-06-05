package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Map struct {
	size pixel.Vec
	grid [][]Tile

}

// Inits the grid for the map
func (m *Map) CreateGrid(size pixel.Vec) {
	m.size = size
	m.resizeGrid()
	m.loadMap()
}

// TODO make this take in a config file??
func (m *Map) loadMap() {
	for x := 0; x < int(m.size.X); x++ {
		for y := 0; y < int(m.size.Y); y++ {
			m.grid[x][y].SetSprite(loadSprite("img_grass2_3.png"))
			m.grid[x][y].SetTileSize(pixel.V(32,32))
			m.grid[x][y].SetTileMapPosition(pixel.V(float64(x), float64(y)))
		}
	}
}

// re-sizes the array to be the current size set
func (m *Map) resizeGrid() {
	m.grid = make([][]Tile, int(m.size.X))

	for x := 0; x < len(m.grid); x++ {
		m.grid[x] = make([]Tile, int(m.size.Y))
	}
}

func (m *Map) Draw(window *pixelgl.Window) {

	for x := 0; x < int(m.size.X); x++ {
		for y := 0; y < int(m.size.Y); y++  {
			m.grid[x][y].Draw(window)
		}
	}
}