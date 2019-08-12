package main

type Biome struct {
	AvailableTiles []AvailableTile   `json:available_tiles`
	TileConfig     []BiomeTileConfig `json:tile_config`
}

func (b *Biome) getNewConfig(config BiomeTileConfig, isDeath bool) BiomeTileConfig {

	var newTile []string

	for _, element := range b.AvailableTiles {
		if element.Type == config.Tile {
			if isDeath {
				newTile = element.Reproduce
			} else {
				newTile = element.Death
			}
			break
		}
	}

	options := make([]BiomeTileConfig, 0)

	for _, element := range b.TileConfig {
		if element.Tile == config.Tile {
			options = append(options, element)
		}
	}

	recordToUse := randomInstance.Intn(len(options));



	config.Tile
}