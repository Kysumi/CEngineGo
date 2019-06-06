package main

type TileDetails struct {
	spriteName string
	maxTileCount int
}


func getTiles() map[string]TileDetails {

	tilesSets := make(map[string]TileDetails)


	tilesSets["grass"] = append(tilesSets["grass"], TileDetails{})
}