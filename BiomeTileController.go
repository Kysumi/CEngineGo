package main

import (
	"github.com/faiface/pixel"
)

type BiomeTileController struct {
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

func (bt *BiomeTileController) checkIfMatching(controller *BiomeTileController) bool {
	return bt.biomeTileConfig.Tile == controller.biomeTileConfig.Tile
}

func (bt *BiomeTileController) setNewBiome(biome Biome, config BiomeTileConfig) {


	bt.setNewSprite()
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

	count := 0

	if bt.biomeTileConfig.UnderPop == -1 {
		return
	}

	for _, element := range neighbours {
		if element.tileController.(*BiomeTileController).biomeTileConfig.Tile == bt.biomeTileConfig.Tile {
			count++
		}
	}

	if bt.biomeTileConfig.UnderPop <= count && bt.biomeTileConfig.CanDie {
		bt.setNewBiome(bt.biome, bt.biome.getNewConfig(bt.biomeTileConfig, true))
	}
}

// Rule 2
func (bt *BiomeTileController) checkLiveToNextGeneration(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if bt.checkIfMatching(element.tileController.(*BiomeTileController)) {
			count++
		}
	}

	if bt.biomeTileConfig.CanDie == false {
		return
	}

	if bt.biomeTileConfig.Reproduction >= count {

		randomInt := randomInstance.Intn(len(neighbours))
		tile := neighbours[randomInt]

		tileController := tile.tileController.(*BiomeTileController)
		tileController.setNewBiome(bt.biome, bt.biome.getNewConfig(bt.biomeTileConfig, false))
	}
}

// Rule 3
func (bt *BiomeTileController) checkOverPopulation(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if bt.checkIfMatching(element.tileController.(*BiomeTileController)) {
			count++
		}
	}

	if bt.biomeTileConfig.OverPop == -1 {
		return
	}

	if bt.biomeTileConfig.OverPop <= count {
		bt.setNewBiome(bt.biome, bt.biome.getNewConfig(bt.biomeTileConfig, true))
	}
}

// Rule 4
func (bt *BiomeTileController) checkReproduction(neighbours []*Tile) {
	count := 0

	for _, element := range neighbours {
		if bt.checkIfMatching(element.tileController.(*BiomeTileController)) {
			count++
		}
	}

	if bt.biomeTileConfig.Reproduction == -1 {
		return
	}

	if bt.biomeTileConfig.Reproduction > count {
		targetCount := bt.biomeTileConfig.Reproduction
		if len(neighbours) < targetCount {
			targetCount = len(neighbours)
		}

		countChanged := 0

		for (countChanged + count) < targetCount {
			randomInt := randomInstance.Intn(len(neighbours))
			tile := neighbours[randomInt]
			tileController := tile.tileController.(*BiomeTileController)


			if tileController.biomeTileConfig.Tile != bt.biomeTileConfig.Tile {
				tileController.setNewBiome(bt.biome, bt.biome.getNewConfig(bt.biomeTileConfig, false))
				countChanged++
			}
		}
	}
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