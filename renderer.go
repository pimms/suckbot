package main

import (
	"github.com/jackyb/go-sdl2/sdl"
	"github.com/pimms/suckbot/agent"
	"github.com/pimms/suckbot/env"
)

type t_renderer struct {
	window     *sdl.Window
	surface    *sdl.Surface
	renderer   *sdl.Renderer
	shouldExit bool
}

func (t *t_renderer) pollEvents() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {
		case *sdl.QuitEvent:
			t.shouldExit = true
		}
	}
}

func (t *t_renderer) createWindow() {
	if t.window == nil {
		t.window = sdl.CreateWindow("SuckBot",
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			800, 600, sdl.WINDOW_SHOWN)

		t.surface = t.window.GetSurface()

		t.renderer = sdl.CreateRenderer(t.window, -1,
			sdl.RENDERER_ACCELERATED)
	}
}

func (t *t_renderer) destroyWindow() {
	if t.renderer != nil {
		t.renderer.Destroy()
	}

	if t.window != nil {
		t.window.Destroy()
	}

	t.window = nil
	t.surface = nil
	t.renderer = nil
}

func (t *t_renderer) renderFrame(cont *env.Controller, agent *agent.Agent) {
	var tiles [env.MAX_SIZE][env.MAX_SIZE]env.ITile

	tiles = cont.CHEAT_GetTiles()

	t.renderer.Clear()

	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			if tiles[x][y] != nil {
				t.drawTile(tiles[x][y])
			}
		}
	}

	t.renderer.Present()
}

func (t *t_renderer) drawTile(tile env.ITile) {
	var state env.TileState
	var x, y int

	state = tile.GetState()
	x, y = tile.GetIndices()

	x = 60 + 50*x
	y = 60 + 50*y

	if state == env.CLEAN {
		t.renderer.SetDrawColor(100, 255, 100, 255)
	} else {
		t.renderer.SetDrawColor(100, 100, 100, 255)
	}

	var rect sdl.Rect
	rect.X = int32(x)
	rect.Y = int32(y)
	rect.W = 50
	rect.H = 50

	t.renderer.FillRect(&rect)
	t.renderer.SetDrawColor(0, 0, 0, 255)
}
