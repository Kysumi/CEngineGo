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
	count := 0

	for _, element := range neighbours {
		if element.tileType == t.tileType {
			count ++
		}
	}

	if t.tileType.Death == -1 {
		return
	}

	//if t.tileType.Death <= count {
	//	conwayManager.swapTileType(t)
	//}
}

// Rule 3
func (t *TileConway) checkOverPopulation(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if element.tileType == t.tileType {
			count ++
		}
	}

	if t.tileType.OverPop == -1 {
		return
	}

	if t.tileType.OverPop <= count {
		conwayManager.swapTileType(t)
	}
}

// Rule 4
func (t *TileConway) checkReproduction(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if element.tileType == t.tileType {
			count ++
		}
	}

	if t.tileType.Reproduction == -1 {
		return
	}

	if t.tileType.Reproduction > count {
		targetCount := t.tileType.Reproduction
		if len(neighbours) < targetCount {
			targetCount = len(neighbours)
		}

		countChanged := 0

		for (countChanged + count) < targetCount {
			randomInt := randomInstance.Intn(len(neighbours))
			tile := neighbours[randomInt]

			if tile.tileType.BiomeType != t.tileType.BiomeType {
				conwayManager.swapTileTypeUnderPop(&tile.TileConway, t.tileType)
				countChanged++
			}
		}
	}
}



