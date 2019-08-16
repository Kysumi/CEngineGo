package main

import (
	"github.com/faiface/pixel"
)

// Handle the randomized placing of biomes

// Handle other biome related...

type BiomeController struct {
	Biomes     []Biome
	currentMap *Map
}

func NewBiomeController(biomes []Biome, currentMap *Map) *BiomeController {
	return &BiomeController{biomes, currentMap}
}

func randomPosition(maxSize pixel.Vec) pixel.Vec {
	return pixel.V(float64(randomInstance.Intn(int(maxSize.X) + 1)), float64(randomInstance.Intn(int(maxSize.Y)) + 1))
}

func (bc *BiomeController) generateBiomes() {

	// for each biome
	for _, biome := range bc.Biomes {

		// for each config in the biome
		for _, biomeConfigTile := range biome.TileConfig {

			position := randomPosition(currentMap.size)
			neighbours := currentMap.getNeighbourTiles(position, true, 1)

			// Shuffle the array around so you can't see a pattern as easy
			randomInstance.Shuffle(
				len(neighbours),
				func(i, j int) { neighbours[i], neighbours[j] = neighbours[j], neighbours[i] })

			bc.applyBiomeToSet(biomeConfigTile, neighbours, biome)
		}
	}

	bc.fullRemaining()
}

func (bc *BiomeController) applyBiomeToSet(biomeConfig BiomeTileConfig, tiles []*Tile, biome Biome) {
	for _, tile := range tiles {

		newController := NewBiomeTileController(biome, biomeConfig)
		newController.sprite = GetSprite(biome.getRandomSpriteString(biomeConfig.Tile))

		tile.tileController = newController
	}
}

func (bc *BiomeController) fullRemaining() {
	waterBiome := loadBiomes([]string{ "waterBiome"})[0]
	tileConfig := waterBiome.TileConfig[0]

	for x := 0; x < int(currentMap.size.X); x++ {
		for y := 0; y < int(currentMap.size.Y); y++ {

			tile := currentMap.getTileFromGridPosition(pixel.V(float64(x), float64(y)))

			if tile.tileController != nil {
				continue
			}

			newController := NewBiomeTileController(waterBiome, tileConfig)
			newController.sprite = GetSprite(waterBiome.getRandomSpriteString(tileConfig.Tile))

			tile.tileController = newController
		}
	}
}