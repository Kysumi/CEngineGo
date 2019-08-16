package main

type TileConfig struct {
	TileConfig []TileNameFile `json:"tile_config"`
}

type TileNameFile struct {
	Name string `json:"name"`
	Image string `json:"image"`
}