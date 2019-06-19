package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type Player struct {
	GameObject
	sprite *pixel.Sprite
	moveSpeed float64
}

func (p *Player) Init() {
	p.moveSpeed = 500
	p.pixelPosition = pixel.V(0,0)
}

func (p *Player) GetPosition() pixel.Vec {
	return p.pixelPosition
}

func (p *Player) SetSprite(sprite *pixel.Sprite) {
	p.sprite = sprite
}

func (p *Player) Draw(window *pixelgl.Window) {
	p.sprite.Draw(window, pixel.IM.Moved(p.pixelPosition))
}

func (p *Player) Update(window *pixelgl.Window, deltaTime float64) {
	if window.Pressed(pixelgl.KeyW) {
		p.pixelPosition.Y += p.moveSpeed * deltaTime
	}

	if window.Pressed(pixelgl.KeyD) {
		p.pixelPosition.X += p.moveSpeed * deltaTime
	}

	if window.Pressed(pixelgl.KeyS) {
		p.pixelPosition.Y -= p.moveSpeed * deltaTime
	}

	if window.Pressed(pixelgl.KeyA) {
		p.pixelPosition.X -= + p.moveSpeed * deltaTime
	}

	p.mapPosition.X = math.Round(p.pixelPosition.X / currentMap.size.X)
	p.mapPosition.Y = math.Round(p.pixelPosition.Y / currentMap.size.Y)
}