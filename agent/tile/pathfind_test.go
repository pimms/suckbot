package tile

import (
	"github.com/pimms/suckbot/env"
	"testing"
)

func TestRemove(t *testing.T) {
	item0 := new(t_pathnode)
	item1 := new(t_pathnode)

	slice := make([]*t_pathnode, 2, 2)
	slice[0] = item0
	slice[1] = item1

	slice = remove(slice, item0)

	if len(slice) != 1 {
		t.Errorf("Expected length 1, %d received\n", len(slice))
	}

	if slice[0] != item1 {
		t.Errorf("Expected 'item1' as sole element\n")
	}
}

func TestAddToOpen(t *testing.T) {
	itemOpen := new(t_pathnode)

	var state t_pf_state
	state.init(8)

	state.addToOpen(itemOpen)

	if len(state.open) != 1 {
		t.Errorf("Unexpected length: %d (expected 1)\n",
			len(state.open))
	}

	if state.open[0] != itemOpen {
		t.Errorf("I don't even...\n")
	}
}

func TestAddToClosed(t *testing.T) {
	node := new(t_pathnode)

	var state t_pf_state
	state.init(1)

	state.addToOpen(node)
	state.addToClosed(node)

	if len(state.open) != 0 {
		t.Errorf("Node not removed from open list\n")
	}

	if len(state.closed) != 1 {
		t.Error("Node not added to closed list")
	}

	if state.closed[0] != node {
		t.Errorf("The closed node it's not my node. " +
			"It's just some node who thinks 'closed' is it's list.")
	}
}

func TestPathfinding(t *testing.T) {
	//var cont *env.Controller
	var tilestate *t_tilestate
	_, tilestate = pathfindEnvironment(true)

	var start *t_tilewrapper
	var end *t_tilewrapper
	var dir env.Direction

	start = &tilestate.tiles[0][0]
	end = &tilestate.tiles[2][0]

	dir = PathFind(start, end, tilestate)
	if dir != env.UP {
		t.Errorf("Expected %d, got %d\n", env.UP, dir)
	}

	start = &tilestate.tiles[0][1]
	dir = PathFind(start, end, tilestate)
	if dir != env.UP {
		t.Errorf("Expected %d, got %d\n", env.UP, dir)
	}

	start = &tilestate.tiles[0][2]
	dir = PathFind(start, end, tilestate)
	if dir != env.RIGHT {
		t.Errorf("Expected %d, got %d\n", env.RIGHT, dir)
	}

	start = &tilestate.tiles[2][2]
	dir = PathFind(start, end, tilestate)
	if dir != env.DOWN {
		t.Errorf("Expected %d, got %d\n", env.DOWN, dir)
	}

	start = &tilestate.tiles[2][0]
	dir = PathFind(start, end, tilestate)
	if dir != env.NONE {
		t.Errorf("Expected %d, got %d\n", env.NONE, dir)
	}
}

func TestTileFinding(t *testing.T) {
	var tilestate *t_tilestate
	var dir env.Direction
	var tile *t_tilewrapper

	_, tilestate = pathfindEnvironment(false)
	tile = &tilestate.tiles[0][0]

	expect := []env.Direction{
		env.UP, env.UP,
		env.RIGHT, env.RIGHT,
		env.DOWN, env.DOWN,
		env.NONE}

	for i := 0; i < len(expect); i++ {
		var x, y, dx, dy int

		dir = TileFind(tile, TILE_UNKOWN, tilestate)
		x, y = tile.tile.GetIndices()
		dx, dy = env.GetIndices(dir)

		if dir != expect[i] {
			t.Errorf("Expected %d, got %d! ", expect[i], dir)
			return
		}

		if dir != env.NONE {
			// Discover and move to the new tile
			tilestate.AddDiscovery(tile.tile.GetNeighbour(dir))
			tile = &tilestate.tiles[x+dx][y+dy]
		}
	}
}

func pathfindEnvironment(discoverAll bool) (*env.Controller,
	*t_tilestate) {
	var cont *env.Controller
	var tile *t_tilestate
	var tileMap [env.MAX_SIZE][env.MAX_SIZE]bool

	tileMap[0][0] = true
	tileMap[0][1] = true
	tileMap[0][2] = true
	tileMap[1][2] = true
	tileMap[2][2] = true
	tileMap[2][1] = true
	tileMap[2][0] = true

	cont = new(env.Controller)
	cont.InitController(tileMap)

	tile = new(t_tilestate)
	tile.AddDiscovery(cont.GetStartingTile())

	// Mark all walls where there are no tiles
	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			if !tileMap[x][y] {
				tile.tiles[x][y].explored = true
			}
		}
	}

	if discoverAll {
		t := cont.GetStartingTile()

		t = t.GetNeighbour(env.UP) // 0 1
		tile.AddDiscovery(t)

		t = t.GetNeighbour(env.UP) // 0 2
		tile.AddDiscovery(t)

		t = t.GetNeighbour(env.RIGHT) // 1 2
		tile.AddDiscovery(t)

		t = t.GetNeighbour(env.RIGHT) // 2 2
		tile.AddDiscovery(t)

		t = t.GetNeighbour(env.DOWN) // 2 1
		tile.AddDiscovery(t)

		t = t.GetNeighbour(env.DOWN)
		tile.AddDiscovery(t) // 2 0
	}

	return cont, tile
}
