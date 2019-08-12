package main

import "github.com/faiface/pixel"

type BiomeTile struct {
	TileController
	biomeTileConfig BiomeTileConfig
	controllingType int
}

func NewBiomeTile(tileController TileController) *BiomeTile {
	return &BiomeTile{TileController: tileController, controllingType: TcBiome}
}

// func to call when want to process game of life
func (bt *BiomeTile) Tick(vec pixel.Vec) {
	neighbours := currentMap.getNeighbourTilesUnderController(vec, false, bt.biomeTileConfig.NeighboursReach, bt.controllingType)

	bt.checkUnderPopulation(neighbours)
	bt.checkLiveToNextGeneration(neighbours)
	bt.checkOverPopulation(neighbours)
	bt.checkReproduction(neighbours)

	bt.forceGrouping()
}

// Rule 1
func (bt *BiomeTile) checkUnderPopulation(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if element.tileType == bt.tileType {
			count++
		}
	}

	if bt.tileType.UnderPop == -1 {
		return
	}

	if bt.tileType.UnderPop <= count {
		//conwayManager.swapTileType(bt)
	}
}

// Rule 2
func (bt *BiomeTile) checkLiveToNextGeneration(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if element.tileType == bt.tileType {
			count++
		}
	}

	if bt.tileType.CanDie == false {
		return
	}

	//if bt.tileType.Death <= count {
	//
	//	randomInt := randomInstance.Intn(len(neighbours))
	//	tile := neighbours[randomInt]
	//
	//	conwayManager.swapTileTypeUnderPop(bt, tile.tileType)
	//}
}

// Rule 3
func (bt *BiomeTile) checkOverPopulation(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if element.tileType == bt.tileType {
			count++
		}
	}

	if bt.tileType.OverPop == -1 {
		return
	}

	if bt.tileType.OverPop <= count {
		//conwayManager.swapTileType(bt)
	}
}

// Rule 4
func (bt *BiomeTile) checkReproduction(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if element.tileType == bt.tileType {
			count++
		}
	}

	if bt.tileType.Reproduction == -1 {
		return
	}

	if bt.tileType.Reproduction > count {
		targetCount := bt.tileType.Reproduction
		if len(neighbours) < targetCount {
			targetCount = len(neighbours)
		}

		countChanged := 0

		for (countChanged + count) < targetCount {
			randomInt := randomInstance.Intn(len(neighbours))
			tile := neighbours[randomInt]

			if tile.tileType.BiomeType != bt.tileType.BiomeType {
				conwayManager.swapTileTypeUnderPop(&tile.BiomeTile, bt.tileType)
				countChanged++
			}
		}
	}
}

// Helping function to make sure we don'bt have random stray nodes.
func (bt *BiomeTile) forceGrouping() {

	directNeighbours := currentMap.getStraightNeighbourTiles(bt.mapPosition)
	count := 0
	for _, neighbour := range directNeighbours {
		if neighbour.tileType.BiomeType == bt.tileType.BiomeType {
			count++
		}
	}

	if count == 0 {
		loc := randomInstance.Intn(len(directNeighbours))
		temp := directNeighbours[loc]
		conwayManager.swapTileTypeUnderPop(bt, temp.tileType)
	}
}