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
	tileMap[0][1] = true
	tileMap[1][1] = true
	tileMap[2][1] = true

	c.InitController(tileMap)
	return c
}

func _TestUndiscoveredByDefault(t *testing.T) {
	state := new(t_tilestate)

	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			if state.tiles[x][y].explored {
				t.Errorf("Tile [%i,%i] is falsely discovered", x, y)
			}
		}
	}
}

func _TestAddTile(t *testing.T) {
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

	if !state.tiles[x][y].explored {
		t.Error("Discovered tile is not flagged as explored")
	}

	neigh = tile.GetNeighbour(env.RIGHT)
	state.AddDiscovery(neigh)

	if state.tiles[x+1][y].tile != neigh {
		t.Error("Non-matching tile-state")
	}
}

func TestTileStatus(t *testing.T) {
	var cont *env.Controller
	var state t_tilestate
	var base env.ITile
	var result Status

	// base is positioned at [0,0]
	cont = createDummyEnvController()
	base = cont.GetStartingTile()
	state.AddDiscovery(base)

	// Left is an invalid index [-1, 0]
	result = state.GetTileStatus(base, env.LEFT)
	if result != TILE_INVALID {
		t.Errorf("Expected %d, received %d\n", TILE_INVALID, result)
	}

	// Down is an invalid index [0, -1]
	result = state.GetTileStatus(base, env.DOWN)
	if result != TILE_INVALID {
		t.Errorf("Expected %d, received %d\n", TILE_INVALID, result)
	}

	// Right is an undiscovered tile
	result = state.GetTileStatus(base, env.RIGHT)
	if result != TILE_UNKOWN {
		t.Errorf("Expected %d, received %d\n", TILE_UNKOWN, result)
	}

	// Up is an undiscovered tile
	result = state.GetTileStatus(base, env.UP)
	if result != TILE_UNKOWN {
		t.Errorf("Expected %d, received %d\n", TILE_UNKOWN, result)
	}

	// Discover tile to the right
	base = base.GetNeighbour(env.RIGHT)
	state.AddDiscovery(base)

	result = state.GetTileStatus(base, env.LEFT)
	if result != TILE_DISCOVERED {
		t.Errorf("Expected %d, received %d\n", TILE_DISCOVERED, result)
	}

	// Move to the tile above the original base
	base = base.GetNeighbour(env.UP)
	state.AddDiscovery(base)
	base = base.GetNeighbour(env.LEFT)
	state.AddDiscovery(base)

	result = state.GetTileStatus(base, env.DOWN)
	if result != TILE_DISCOVERED {
		t.Errorf("Expected %d, received %d\n", TILE_DISCOVERED, result)
	}
}
