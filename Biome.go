package main

type Biome struct {
	AvailableTiles []AvailableTile   `json:available_tiles`
	TileConfig     []BiomeTileConfig `json:tile_config`
}