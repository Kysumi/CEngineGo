package main

import "github.com/faiface/pixel"

type BiomeTile struct {
	TileController
	biomeTileConfig BiomeTileConfig
}



// Sets the tile the BiomeTile should affect.
func (b *BiomeTile) SetPosition(pos pixel.Vec) {

}