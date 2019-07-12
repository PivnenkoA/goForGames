package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	w     int32  = 800
	h     int32  = 600
	title string = "Some game"
)

type colorRGBA struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func errorer(e error, c int) int {
	if e != nil {
		panic(e)
	}
	return c
}

func main() {
	//white := colorRGBA{255, 255, 255, 255}
	//green := colorRGBA{0, 155, 0, 255}
	black := colorRGBA{0, 0, 0, 255}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_OPENGL)
	errorer(err, 1)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	errorer(err, 2)
	defer renderer.Destroy()

	plr, err := newPlayer(renderer)

	errorer(err, 3)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(black.r, black.g, black.b, black.a)
		renderer.Clear()
		plr.updatePlayer()

		plr.draw(renderer)

		renderer.Present()
	}

}
