package main

import (
	// "example.com/m/loaders"
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
			//m.grid[x][y].SetSprite(GetSprite("img_water_0.png"))
			m.grid[x][y].SetTileSize(pixel.V(32, 32))
			m.grid[x][y].SetTileMapPosition(pixel.V(float64(x), float64(y)))
			//m.grid[x][y].tileType = conwayManager.TileDetails[0]
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
		for y := 0; y < int(m.size.Y); y++ {
			m.grid[x][y].Draw(window)
		}
	}
}

// Checks if the grid location is valid.
func (m *Map) WithinMapBounds(position pixel.Vec) bool {

	if int(position.Y) >= len(m.grid) || int(position.Y) < 0 {
		return false
	}

	if int(position.X) >= len(m.grid) || int(position.X) < 0 {
		return false
	}

	return true
}

func (m *Map) getTileFromGridPosition(position pixel.Vec) *Tile {
	return &m.grid[int(position.X)][int(position.Y)]
}

func (m *Map) getNeighbourTiles(position pixel.Vec, withParent bool, reach int) []*Tile {

	tiles := make([]*Tile, 0)

	for x := -reach; x <= reach; x++ {
		for y := -1; y <= 1; y++ {
			newPosition := pixel.V(position.X+float64(x), position.Y+float64(y))

			// If it is the current tile ignore
			if withParent == false {
				if position.X == newPosition.X && position.Y == newPosition.Y {
					continue
				}
			}

			if !m.WithinMapBounds(newPosition) {
				continue
			}

			tiles = append(tiles, m.getTileFromGridPosition(newPosition))
		}
	}

	return tiles
}

func (m *Map) getNeighbourTilesUnderController(position pixel.Vec, withParent bool, reach int, controller int) []*Tile {

	tiles := make([]*Tile, 0)

	for x := -reach; x <= reach; x++ {
		for y := -1; y <= 1; y++ {
			newPosition := pixel.V(position.X+float64(x), position.Y+float64(y))

			// If it is the current tile ignore
			if withParent == false {
				if position.X == newPosition.X && position.Y == newPosition.Y {
					continue
				}
			}

			if !m.WithinMapBounds(newPosition) {
				continue
			}

			tile := m.getTileFromGridPosition(newPosition)

			if !tile.IsControlledBy(controller) {
				continue
			}

			tiles = append(tiles, tile)
		}
	}

	return tiles
}


func (m *Map) getStraightNeighbourTiles(position pixel.Vec) []*Tile {

	tiles := make([]*Tile, 0)

	positions := []pixel.Vec{
		pixel.V(position.X-1, position.Y),
		pixel.V(position.X+1, position.Y),
		pixel.V(position.X, position.Y-1),
		pixel.V(position.X, position.Y+1)}

	for _, element := range positions {
		if !m.WithinMapBounds(element) {
			continue
		}

		tiles = append(tiles, m.getTileFromGridPosition(element))
	}

	return tiles
}