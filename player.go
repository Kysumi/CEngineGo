package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	position pixel.Vec
	sprite *pixel.Sprite
	moveSpeed float64
}

func (p *Player) Init() {
	p.moveSpeed = 1.1
	p.position = pixel.V(0,0)
}

func (p *Player) SetSprite(sprite *pixel.Sprite) {
	p.sprite = sprite
}

func (p *Player) Draw(window *pixelgl.Window) {
	p.sprite.Draw(window, pixel.IM.Moved(p.position))
}


func (p *Player) Update(window *pixelgl.Window, deltaTime float64) {
	if window.JustPressed(pixelgl.KeyW) {
		p.position.Y += (p.position.Y * p.moveSpeed) * deltaTime
	}

	if window.JustPressed(pixelgl.KeyD) {
		p.position.X += (p.position.X * p.moveSpeed) * deltaTime
	}

	if window.JustPressed(pixelgl.KeyS) {
		p.position.Y -= (p.position.Y * p.moveSpeed) * deltaTime
	}

	if window.JustPressed(pixelgl.KeyA) {
		window.SetClosed(true)
	}
}