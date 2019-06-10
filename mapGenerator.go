package main

import (
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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return pixel.V(float64(r.Intn(int(maxSize.X) + 1)), float64(r.Intn(int(maxSize.Y)) + 1))
}

func getTiles() map[string][]TileDetails {

	tilesSets := make(map[string][]TileDetails)

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_0.png", 20, true, 0})

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_1.png", 100, true, 0})

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_2.png", 60, true, 0})

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_3.png", 20, true, 0})

	tilesSets["grass"] = append(
		tilesSets["grass"],
		TileDetails{"img_grass_4.png", 300, true, 0})

	return tilesSets
}

func generate(newMap *Map) {
	tiles := getTiles()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(tiles["grass"]) > 0 {
		position := randomPosition(newMap.size)
		neighbours := newMap.getNeighbourTiles(position, true)

		for _, neighbour := range neighbours {

			tileOptions := len(tiles["grass"])

			if tileOptions <= 0 {
				break
			}

			arrayLoc := r.Intn(tileOptions)
			currentTileDetails := &tiles["grass"][arrayLoc]

			neighbour.SetSprite(loadSprite(currentTileDetails.spriteName))

			currentTileDetails.consumed += 1

			if currentTileDetails.consumed >= currentTileDetails.maxTileCount {
				tiles["grass"] = append(tiles["grass"][:arrayLoc], tiles["grass"][arrayLoc+1:]...)
			}
		}
	}
}