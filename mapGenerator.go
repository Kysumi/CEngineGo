package main

type TileDetails struct {
	spriteName string
	maxTileCount int
}


func getTiles() map[string][]TileDetails {

	tilesSets := make(map[string][]TileDetails)

	tilesSets["grass"] = append(tilesSets["grass"], TileDetails{"", 20})
	tilesSets["grass"] = append(tilesSets["grass"], TileDetails{"", 10})


	return tilesSets
}

func generate(newMap *Map) {

	//tiles := getTiles()
	//
	//newMap.getNeighbourTiles(pixel.V(0,0))
	//
	//tiles["grass"]


}