package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	GameObject
	sprite *pixel.Sprite
	moveSpeed float64
}

func (p *Player) Init() {
	p.moveSpeed = 500
	p.position = pixel.V(0,0)
}

func (p *Player) GetPosition() pixel.Vec {
	return p.position
}

func (p *Player) SetSprite(sprite *pixel.Sprite) {
	p.sprite = sprite
}

func (p *Player) Draw(window *pixelgl.Window) {
	p.sprite.Draw(window, pixel.IM.Moved(p.position))
}

func (p *Player) Update(window *pixelgl.Window, deltaTime float64) {
	if window.Pressed(pixelgl.KeyW) {
		p.position.Y += p.moveSpeed * deltaTime
	}

	if window.Pressed(pixelgl.KeyD) {
		p.position.X += p.moveSpeed * deltaTime
	}

	if window.Pressed(pixelgl.KeyS) {
		p.position.Y -= p.moveSpeed * deltaTime
	}

	if window.Pressed(pixelgl.KeyA) {
		p.position.X -= + p.moveSpeed * deltaTime
	}
}