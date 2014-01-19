package tile

import (
	"github.com/pimms/suckbot/env"
	"testing"
)

func createDummyEnvController() *env.Controller {
	var c *env.Controller = new(env.Controller)

	var tileMap [env.MAX_SIZE][env.MAX_SIZE]bool
	tileMap[0][0] = true
	tileMap[1][0] = true
	tileMap[2][0] = true
	tileMap[1][1] = true

	c.InitController(tileMap)
	return c
}

func TestUndiscoveredByDefault(t *testing.T) {
	state := new(t_tilestate)

	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			if state.tiles[x][y].explored {
				t.Errorf("Tile [%i,%i] is falsely discovered", x, y)
			}
		}
	}
}

func TestAddTile(t *testing.T) {
	var cont *env.Controller
	var state t_tilestate
	var tile env.ITile
	var neigh env.ITile
	var x, y int

	cont = createDummyEnvController()
	tile = cont.GetStartingTile()
	state.AddDiscovery(tile)
	x, y = tile.GetIndices()

	if state.tiles[x][y].tile != tile {
		t.Error("Non-matching tile-state")
	}

	neigh = tile.GetNeighbour(env.RIGHT)
	state.AddDiscovery(neigh)

	if state.tiles[x+1][y].tile != neigh {
		t.Error("Non-matching tile-state")
	}
}
