package agent

import (
	"github.com/pimms/suckbot/agent/tile"
	"github.com/pimms/suckbot/env"
)

type Agent struct {
	tileState tile.TileState
}

func (a *Agent) CHEAT_GetTileStatus(x, y int) tile.Status {
	return a.tileState.GetTileStatusAtCoord(x, y)
}

func (a *Agent) Initialize(startTile env.ITile) {

}
