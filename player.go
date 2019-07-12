package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const playerSpeed float64 = 0.09

type player struct {
	tex  *sdl.Texture
	x, y float64
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {

	img, err := sdl.LoadBMP("sprites/hero.bmp")
	if err != nil {
		return player{}, fmt.Errorf("not texture %v", err)
	}
	defer img.Free()

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("creating texture %v", err)
	}
	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex, &sdl.Rect{X: 0, Y: 0, W: 50, H: 50}, &sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 50, H: 50})
}

func (p *player) updatePlayer() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	}
}
