package main

// Handles the config per tile per biome
// from the biome config
type BiomeTileConfig struct {
	UnderPop        int    `json:"under_pop"`
	OverPop         int    `json:"over_pop"`
	Reproduction    int    `json:"reproduction"`
	CanDie          bool   `json:"can_die"`
	Tile            string `json:"tile"`
	NeighboursReach int    `json:"neighbours_reach"`
}
