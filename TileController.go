package main

const (
	TC_BIOME = iota
	TC_OTHER = iota
)

// This exists so in future i can have
// different objects controlling the tiles without
// issues. :pray:
type TileController struct {
	ControllingType int
}

func (tc *TileController) getControllingType() int {
	return tc.ControllingType
}

func (tc *TileController) setControllingType() *TileController {
	return tc
}