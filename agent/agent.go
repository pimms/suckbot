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
	fmt.Print("Selected action:\t")

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

	fmt.Print("\n")
}

func (a *Agent) getAction() int {
	if a.currentTile.GetITile().GetState() == env.DIRTY {
		return SUCK
	}

	if !a.fullyExplored {
		// If there are no tiles left to explore,
		// fall through to the "a.fullyExplored"-case.
		var dir = int(a.getSearchDirection())
		if dir != int(env.NONE) {
			return dir
		}
	}

	if a.fullyExplored {
		fmt.Println("Time to suck, innit??? m8 ")
	}

	return NOOP
}

func (a *Agent) performAction(action int) {
	switch action {
	case NOOP:

	case SUCK:
		a.vacuumCurrent()

	case int(env.UP):
		fallthrough
	case int(env.RIGHT):
		fallthrough
	case int(env.DOWN):
		fallthrough
	case int(env.LEFT):
		a.moveInDirection(env.Direction(action))
	}
}

func (a *Agent) getSearchDirection() env.Direction {
	dir := tile.TileFind(a.currentTile,
		tile.TILE_UNKOWN, &a.tileState)

	if dir == env.NONE {
		a.fullyExplored = true
	}

	return dir
}

func (a *Agent) vacuumCurrent() {
	// Clean the tile I'm currently standing on
	a.currentTile.GetITile().OnVacuum()
}

func (a *Agent) moveInDirection(dir env.Direction) {
	var itile env.ITile

	itile = a.currentTile.GetITile().GetNeighbour(dir)

	if itile != nil {
		a.tileState.AddDiscovery(itile)
		a.currentTile = a.tileState.GetWrapper(itile)
	} else {
		x, y := a.currentTile.GetITile().GetIndices()
		dx, dy := env.GetIndices(dir)
		a.tileState.AddDiscoveryNil(x+dx, y+dy)
	}
}
