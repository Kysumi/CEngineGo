package main

import (
	"example.com/m/loaders"
	"github.com/faiface/pixel"
	"math/rand"
	"time"
)

type TileDetails struct {
	spriteName string
	maxTileCount int
	walkable bool
	consumed int
}

func randomPosition(maxSize pixel.Vec) pixel.Vec {
	return pixel.V(float64(randomInstance.Intn(int(maxSize.X) + 1)), float64(randomInstance.Intn(int(maxSize.Y)) + 1))
}

func getTiles() map[string][]TileDetails {

	tilesSets := make(map[string][]TileDetails)

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_0.png", 150, true, 0})

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_1.png", 100, true, 0})

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_2.png", 200, true, 0})

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_3.png", 100, true, 0})

	tilesSets["dirt"] = append(
		tilesSets["dirt"],
		TileDetails{"img_dirt_0.png", 150, true, 0})

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_4.png", 300, true, 0})

	tilesSets["barron"] = append(
		tilesSets["barron"],
		TileDetails{"img_barron_0.png", 80, true, 0})

	tilesSets["sand"] = append(
		tilesSets["sand"],
		TileDetails{"img_sand_0.png", 200, true, 0})

	tilesSets["stone"] = append(
		tilesSets["stone"],
		TileDetails{"img_stone_0.png", 30, true, 0})

	tilesSets["water"] = append(
		tilesSets["water"],
		TileDetails{"img_water_0.png", 0, false, 0})

	return tilesSets
}

func generate(newMap *Map) {
	tiles := getTiles()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for key := range tiles {
		for len(tiles[key]) > 0 {
			position := randomPosition(newMap.size)
			neighbours := newMap.getNeighbourTiles(position, true)

			for _, neighbour := range neighbours {
				if rand.Intn(10) > 2 {
					continue
				}
				tileOptions := len(tiles[key])

				if tileOptions <= 0 {
					break
				}

				arrayLoc := r.Intn(tileOptions)
				currentTileDetails := &tiles[key][arrayLoc]

				for _, conway := range conwayManager.TileDetails {
					if conway.BiomeType == key {
						neighbour.tileType = conway
					}
				}

				neighbour.SetSprite(loaders.GetSprite(currentTileDetails.spriteName))

				currentTileDetails.consumed += 1

				if currentTileDetails.consumed >= currentTileDetails.maxTileCount {
					tiles[key] = append(tiles[key][:arrayLoc], tiles[key][arrayLoc+1:]...)
				}
			}
		}
	}
}