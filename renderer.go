package main

import (
	"github.com/jackyb/go-sdl2/sdl"
	"github.com/pimms/suckbot/agent"
	"github.com/pimms/suckbot/agent/tile"
	"github.com/pimms/suckbot/env"
)

const (
	TILE_SIZE = 50
)

type t_renderer struct {
	window     *sdl.Window
	surface    *sdl.Surface
	renderer   *sdl.Renderer
	shouldExit bool

	texQuestion *sdl.Texture
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

	// Draw the tiles
	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			if tiles[x][y] != nil {
				if tiles[x][y].GetState() == env.CLEAN {
					t.drawTile(x, env.MAX_SIZE-y-1, 100, 255, 100, 255)
				} else {
					t.drawTile(x, env.MAX_SIZE-y-1, 100, 100, 100, 255)
				}
			}
		}
	}

	// Draw the undiscovered areas
	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			var status tile.Status
			status = agent.CHEAT_GetTileStatus(x, y)

			if status == tile.TILE_UNKOWN {
				t.drawQuestionMark(x, y)
			}
		}
	}

	// Draw the agent
	// TODO: Use the agent's actual position
	x, y := agent.CHEAT_GetCurrentTile().GetIndices()
	t.drawAgent(x, env.MAX_SIZE-y-1)

	t.renderer.Present()
}

func (t *t_renderer) drawTile(x, y int, r, g, b, a uint8) {
	x = TILE_SIZE * x
	y = TILE_SIZE * y

	t.renderer.SetDrawColor(r, g, b, a)

	var rect sdl.Rect
	rect.X = int32(x)
	rect.Y = int32(y)
	rect.W = TILE_SIZE
	rect.H = TILE_SIZE

	t.renderer.FillRect(&rect)
	t.renderer.SetDrawColor(0, 0, 0, 255)
}

func (t *t_renderer) drawQuestionMark(x, y int) {
	points := []sdl.Point{
		{1, 2}, {3, 0},
		{3, 0}, {5, 0},
		{5, 0}, {7, 2},
		{7, 2}, {7, 3},
		{7, 3}, {4, 6},
		{4, 6}, {4, 7},
		{4, 9}, {4, 11},
		{3, 10}, {5, 10}}

	t.renderer.SetDrawColor(255, 0, 0, 255)

	for i := 0; i < len(points)/2; i++ {
		var ptA, ptB sdl.Point
		ptA = points[i*2]
		ptB = points[i*2+1]

		ptA.X = (ptA.X * (TILE_SIZE - 20) / 10) + 12 + int32(x*TILE_SIZE)
		ptA.Y = (ptA.Y * (TILE_SIZE - 20) / 10) + 5 + int32(y*TILE_SIZE)
		ptB.X = (ptB.X * (TILE_SIZE - 20) / 10) + 12 + int32(x*TILE_SIZE)
		ptB.Y = (ptB.Y * (TILE_SIZE - 20) / 10) + 5 + int32(y*TILE_SIZE)

		t.renderer.DrawLine(
			int(ptA.X), int(ptA.Y),
			int(ptB.X), int(ptB.Y))
	}

	t.renderer.SetDrawColor(0, 0, 0, 255)
}

func (t *t_renderer) drawAgent(x, y int) {
	var rect sdl.Rect
	rect.X = 10 + int32(x*TILE_SIZE)
	rect.Y = 10 + int32(y*TILE_SIZE)
	rect.W = 30
	rect.H = 30

	t.renderer.SetDrawColor(0, 0, 255, 255)
	t.renderer.FillRect(&rect)
	t.renderer.SetDrawColor(0, 0, 0, 255)
}
