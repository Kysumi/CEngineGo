package main

type TileConway struct {
	GameObject
	tileType 	 TileDetailsConway
}

// func to call when want to process game of life
func (t *TileConway) Tick() {
	neighbours := currentMap.getNeighbourTiles(t.mapPosition, false)
	t.checkUnderPopulation(neighbours)
	t.checkLiveToNextGeneration(neighbours)
	t.checkOverPopulation(neighbours)
	t.checkReproduction(neighbours)
}

// Rule 1
func (t *TileConway) checkUnderPopulation(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if element.tileType == t.tileType {
			count ++
		}
	}

	if t.tileType.UnderPop == -1 {
		return
	}

	if t.tileType.UnderPop <= count {
		conwayManager.swapTileType(t)
	}
}

// Rule 2
func (t *TileConway) checkLiveToNextGeneration(neighbours []*Tile) {

}

// Rule 3
func (t *TileConway) checkOverPopulation(neighbours []*Tile) {

}

// Rule 4
func (t *TileConway) checkReproduction(neighbours []*Tile) {

}



