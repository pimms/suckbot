package main

import (
	"github.com/jackyb/go-sdl2/sdl"
	"github.com/pimms/suckbot/agent"
	"github.com/pimms/suckbot/env"
)

type t_renderer struct {
	window     *sdl.Window
	surface    *sdl.Surface
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
	}
}

func (t *t_renderer) destroyWindow() {
	if t.window != nil {
		t.window.Destroy()
	}

	t.window = nil
	t.surface = nil
}

func (t *t_renderer) renderFrame(cont *env.Controller, agent *agent.Agent) {

}
