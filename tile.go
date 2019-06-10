package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Tile struct {
	tileSize pixel.Vec
	tileMapPosition pixel.Vec
	tilePixelPosition pixel.Vec

	sprite *pixel.Sprite
	walkable bool
}

func (t *Tile) Sprite() *pixel.Sprite {
	return t.sprite
}

func (t *Tile) SetSprite(sprite *pixel.Sprite) {
	t.sprite = sprite
}

func (t *Tile) TileMapPosition() pixel.Vec {
	return t.tileMapPosition
}

// Updates both the pixel post and tilemap pos
func (t *Tile) SetTileMapPosition(tileMapPosition pixel.Vec) {
	t.tileMapPosition = tileMapPosition
	t.tilePixelPosition = pixel.V(tileMapPosition.X * t.tileSize.X, tileMapPosition.Y * t.tileSize.Y)
}

func (t *Tile) TileSize() pixel.Vec {
	return t.tileSize
}

func (t *Tile) SetTileSize(tileSize pixel.Vec) {
	t.tileSize = tileSize
}

func (t *Tile) TilePixelPosition() pixel.Vec {
	return t.tilePixelPosition
}

func (t *Tile) Draw(window *pixelgl.Window) {
	t.sprite.Draw(window, pixel.IM.Moved(t.tilePixelPosition))
}