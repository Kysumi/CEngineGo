package main

import "github.com/faiface/pixel"

const (
	TcBiome = iota
	TcOther = iota
)

// This exists so in future i can have
// different objects controlling the tiles without
// issues. :pray:
type TileController interface {
	getControllingType(tc *TileController) int
	Tick(tc *TileController) (vec pixel.Vec)
}