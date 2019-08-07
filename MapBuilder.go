package main

import (
	"github.com/faiface/pixel"
)

type MapBuilder struct {
	Biomes []Biome
	Size   pixel.Vec
}

func (mb *MapBuilder) AddBiome(biome Biome) *MapBuilder {
	mb.Biomes = append(mb.Biomes, biome)
	return mb
}

func (mb *MapBuilder) SetSize(size pixel.Vec) {
	mb.Size = size
}

func (mb *MapBuilder) build() *Map {
	mappy := new(Map)
	mappy.CreateGrid(mb.Size)

	return mappy
}