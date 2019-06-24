package main

import (
	"encoding/json"
	"example.com/m/loaders"
	"io/ioutil"
	"math/rand"
	"time"
)

type ConwayManager struct {
	TileDetails []TileDetailsConway `json:"tileDetails"`
}

func (c *ConwayManager) swapTileType(tile *TileConway) {

	tile.tileType = c.TileDetails[tile.tileType.DeathTileType]

	mapTile := currentMap.getTileFromGridPosition(tile.mapPosition)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tiles := getTiles()
	key := tile.tileType.BiomeType
	details := tiles[key][r.Intn(len(tiles[key]))]

	mapTile.SetSprite(loaders.GetSprite(details.spriteName))
}


func (c *ConwayManager) Init() {
	// read file
	data, err := ioutil.ReadFile("./tileDetails.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(data), &c)
	if err != nil {
		panic(err)
	}
}
