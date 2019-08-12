package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Tile struct {
	GameObject
	tileSize pixel.Vec

	sprite   *pixel.Sprite
	walkable bool
	tileController TileController
}

func (t *Tile) Sprite() *pixel.Sprite {
	return t.sprite
}

func (t *Tile) SetSprite(sprite *pixel.Sprite) {
	t.sprite = sprite
}

func (t *Tile) TileMapPosition() pixel.Vec {
	return t.mapPosition
}

// Updates both the pixel post and tilemap pos
func (t *Tile) SetTileMapPosition(tileMapPosition pixel.Vec) {
	t.mapPosition = tileMapPosition
	t.pixelPosition = pixel.V(tileMapPosition.X*t.tileSize.X, tileMapPosition.Y*t.tileSize.Y)
}

func (t *Tile) TileSize() pixel.Vec {
	return t.tileSize
}

func (t *Tile) SetTileSize(tileSize pixel.Vec) {
	t.tileSize = tileSize
}

func (t *Tile) TilePixelPosition() pixel.Vec {
	return t.pixelPosition
}

func (t *Tile) Draw(window *pixelgl.Window) {
	t.sprite.Draw(window, pixel.IM.Moved(t.pixelPosition))
}

func (t *Tile) SetTileController(window *pixelgl.Window) {
	t.sprite.Draw(window, pixel.IM.Moved(t.pixelPosition))
}

func (t *Tile) Tick() {
	t.tileController.Tick(t.mapPosition)
}

func (t *Tile) IsControlledBy(controllerId int) bool {
	return t.tileController.getControllingType() == controllerId
}