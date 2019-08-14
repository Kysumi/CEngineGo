package main

type Biome struct {
	AvailableTiles []AvailableTile   `json:available_tiles`
	TileConfig     []BiomeTileConfig `json:tile_config`
}

func (b *Biome) getNewConfig(config BiomeTileConfig, isDeath bool) BiomeTileConfig {

	var tilesAvailable []string

	for _, element := range b.AvailableTiles {
		if element.Type == config.Tile {
			if isDeath {
				tilesAvailable = element.Reproduce
			} else {
				tilesAvailable = element.Death
			}
			break
		}
	}

	options := make([]BiomeTileConfig, 0)

	for _, element := range b.TileConfig {
		for _, tileOption := range tilesAvailable {
			if element.Tile == tileOption {
				options = append(options, element)
			}
		}
	}

	recordToUse := randomInstance.Intn(len(options))

	return options[recordToUse]
}