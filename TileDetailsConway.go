package main

type TileDetailsConway struct {
	DeathTileType int `json:"death_tile_type"`
	UnderPop     int `json:"under_pop"`
	OverPop      int `json:"over_pop"`
	Reproduction int `json:"reproduction"`
	Death        int `json:"death"`
	BiomeType   string `json:"biome_type"`
	NeighboursReach   int `json:"neighbours_reach"`
}
