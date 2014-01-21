package agent

import (
	"github.com/pimms/suckbot/agent/tile"
	"github.com/pimms/suckbot/env"
)

type Agent struct {
	tileState   tile.TileState
	currentTile env.ITile
}

func (a *Agent) CHEAT_GetTileStatus(x, y int) tile.Status {
	return a.tileState.GetTileStatusAtCoord(x, y)
}

func (a *Agent) CHEAT_GetCurrentTile() env.ITile {
	return a.currentTile
}

func (a *Agent) Initialize(startTile env.ITile) {
	a.tileState.AddDiscovery(startTile)
	a.currentTile = startTile
}

func (a *Agent) Tick() {
	// TODO:
	// Return the ACTION from a function analyzing
	// the current state, and have another function
	// actually execute the action.
	if a.currentTile.GetState() == env.DIRTY {
		a.vacuumCurrent()
	}
}

func (a *Agent) vacuumCurrent() {
	a.currentTile.OnVacuum()
}
