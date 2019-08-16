package main

import (
	"github.com/faiface/pixel"
)

type BiomeTileController struct {
	TileController
	biome Biome
	biomeTileConfig BiomeTileConfig
	controllingType int

	sprite *pixel.Sprite
}

func NewBiomeTileController(biome Biome, config BiomeTileConfig) *BiomeTileController {
	biomeTileController := &BiomeTileController{controllingType: TcBiome, biome: biome, biomeTileConfig: config}
	biomeTileController.setNewSprite()

	return biomeTileController
}

func (bt *BiomeTileController) setNewSprite() {
	bt.sprite =  GetSprite(bt.biome.getRandomSpriteString(bt.biomeTileConfig.Tile))
}

func (bt *BiomeTileController) getControllingType() int {
	return bt.controllingType
}

func (bt *BiomeTileController) getSprite() *pixel.Sprite {
	return bt.sprite
}

// func to call when want to process game of life
func (bt *BiomeTileController) Tick(vec pixel.Vec) {
	neighbours := currentMap.getNeighbourTilesUnderController(
		vec,
		false,
		bt.biomeTileConfig.NeighboursReach,
		bt.controllingType)

	bt.checkUnderPopulation(neighbours)
	bt.checkLiveToNextGeneration(neighbours)
	bt.checkOverPopulation(neighbours)
	bt.checkReproduction(neighbours)

	//bt.forceGrouping()
}

// Rule 1
func (bt *BiomeTileController) checkUnderPopulation(neighbours []*Tile) {

	if bt.biomeTileConfig.UnderPop == -1 {
		return
	}

	if bt.biomeTileConfig.UnderPop <= len(neighbours) && bt.biomeTileConfig.CanDie {
		bt.biomeTileConfig = bt.biome.getNewConfig(bt.biomeTileConfig, true)
		bt.setNewSprite()
	}
}

// Rule 2
func (bt *BiomeTileController) checkLiveToNextGeneration(neighbours []*Tile) {
	//count := 0

	//for _, element := range neighbours {
	//	if element.tileType == bt.tileType {
	//		count++
	//	}
	//}
	//
	//if bt.tileType.CanDie == false {
	//	return
	//}
	//
	//if bt.tileType.Death <= count {
	//
	//	randomInt := randomInstance.Intn(len(neighbours))
	//	tile := neighbours[randomInt]
	//
	//	conwayManager.swapTileTypeUnderPop(bt, tile.tileType)
	//}
}

// Rule 3
func (bt *BiomeTileController) checkOverPopulation(neighbours []*Tile) {
	//count := 0

	//for _, element := range neighbours {
	//	if element.tileType == bt.tileType {
	//		count++
	//	}
	//}
	//
	//if bt.tileType.OverPop == -1 {
	//	return
	//}
	//
	//if bt.tileType.OverPop <= count {
	//	//conwayManager.swapTileType(bt)
	//}
}

// Rule 4
func (bt *BiomeTileController) checkReproduction(neighbours []*Tile) {
	//count := 0
	//
	//for _, element := range neighbours {
	//	if element.tileType == bt.tileType {
	//		count++
	//	}
	//}
	//
	//if bt.tileType.Reproduction == -1 {
	//	return
	//}
	//
	//if bt.tileType.Reproduction > count {
	//	targetCount := bt.tileType.Reproduction
	//	if len(neighbours) < targetCount {
	//		targetCount = len(neighbours)
	//	}
	//
	//	countChanged := 0
	//
	//	for (countChanged + count) < targetCount {
	//		randomInt := randomInstance.Intn(len(neighbours))
	//		tile := neighbours[randomInt]
	//
	//		if tile.tileType.BiomeType != bt.tileType.BiomeType {
	//			conwayManager.swapTileTypeUnderPop(&tile.BiomeTileController, bt.tileType)
	//			countChanged++
	//		}
	//	}
	//}
}

// Helping function to make sure we don'bt have random stray nodes.
func (bt *BiomeTileController) forceGrouping() {

	//directNeighbours := currentMap.getStraightNeighbourTiles(bt.mapPosition)
	//count := 0
	//for _, neighbour := range directNeighbours {
	//	if neighbour.tileType.BiomeType == bt.tileType.BiomeType {
	//		count++
	//	}
	//}
	//
	//if count == 0 {
	//	loc := randomInstance.Intn(len(directNeighbours))
	//	temp := directNeighbours[loc]
	//	conwayManager.swapTileTypeUnderPop(bt, temp.tileType)
	//}
}