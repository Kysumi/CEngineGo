package main

import (
	"encoding/json"
	"io/ioutil"
)

func loadBiomes(biomesNames []string) []Biome {

	var loadedBiomes []Biome

	for _, biomeName := range biomesNames {
		data, err := ioutil.ReadFile("./assets/biomes/" + biomeName + ".json")

		if err != nil {
			panic(err)
		}
		var newBiome = new(Biome)
		err = json.Unmarshal([]byte(data), &newBiome)
		if err != nil {
			panic(err)
		}

		loadedBiomes = append(loadedBiomes, *newBiome)
	}

	return loadedBiomes
}

//func (c *ConwayManager) Init() {
//	// read file
//	data, err := ioutil.ReadFile("./tileDetails.json")
//	if err != nil {
//		panic(err)
//	}
//
//	err = json.Unmarshal([]byte(data), &c)
//	if err != nil {
//		panic(err)
//	}
//}
