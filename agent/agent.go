package agent

import (
	"fmt"
	"github.com/pimms/suckbot/agent/tile"
	"github.com/pimms/suckbot/env"
)

const (
	SUCK = 4
	NOOP = 5

	// The directional constants are defined
	// in the "env" package, under the names
	// env.UP, env.LEFT and so on. They are
	// used as actions when the agent should
	// move.
)

type Agent struct {
	tileState     tile.TileState
	fullyExplored bool
	currentTile   *tile.TileWrapper
}

func (a *Agent) CHEAT_GetTileStatus(x, y int) tile.Status {
	return a.tileState.GetTileStatusAtCoord(x, y)
}

func (a *Agent) CHEAT_GetCurrentTile() env.ITile {
	return a.currentTile.GetITile()
}

func (a *Agent) Initialize(startTile env.ITile) {
	a.tileState.AddDiscovery(startTile)

	//fock teh ploice
	a.currentTile = a.tileState.GetWrapper(startTile)
}

func (a *Agent) Tick() {
	var action int

	action = a.getAction()

	a.performAction(action)

	a.printAction(action)
}

func (a *Agent) printAction(action int) {
	switch action {
	case 0:
		fmt.Print("UP")
	case 1:
		fmt.Print("RIGHT")
	case 2:
		fmt.Print("DOWN")
	case 3:
		fmt.Print("LEFT")
	case 4:
		fmt.Print("SUCK")
	case 5:
		fmt.Print("NoOP")
	}
}

func (a *Agent) getAction() int {
	if a.currentTile.GetITile().GetState() == env.DIRTY {
		return SUCK
	}

	if !a.fullyExplored {
		return int(a.getSearchDirection())
	} else if a.fullyExplored {

	}

	return NOOP
}

func (a *Agent) performAction(action int) {
	switch action {
	case NOOP:

	case SUCK:
		a.vacuumCurrent()
	}
}

func (a *Agent) getSearchDirection() env.Direction {
	return tile.TileFind(a.currentTile, tile.TILE_UNKOWN, &a.tileState)
}

func (a *Agent) vacuumCurrent() {
	// Clean the tile I'm currently standing on
	a.currentTile.GetITile().OnVacuum()
}
