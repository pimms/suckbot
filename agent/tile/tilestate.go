package tile

import (
	"github.com/pimms/suckbot/env"
)

type Status int

const (
	TILE_UNKOWN     Status = 0
	TILE_DISCOVERED Status = 1
	TILE_WALL       Status = 2
	TILE_INVALID    Status = 3
)

/* The t_tilestate holds the data about the discovered
 * tiles and the status of the neighbours as well.
 *
 * Tiles are stored in an N by N array, and each tile has
 * it's indices associated with it.
 */
type t_tilestate struct {
	tiles [env.MAX_SIZE][env.MAX_SIZE]t_tilewrapper
}

/* Wrapper around an ITile and it's discovered status. If a
 * tile has been explored but it's tile member is nil, it's
 */
type t_tilewrapper struct {
	tile     env.ITile
	explored bool
}

/*
=======================
	Implementation
=======================
*/
func (t *t_tilestate) AddDiscovery(tile env.ITile) {
	var x, y int
	var twrap t_tilewrapper

	x, y = tile.GetIndices()

	// Create a new t_tilewrapper
	twrap.tile = tile
	t.tiles[x][y] = twrap
	t.tiles[x][y].explored = true
}

func (t *t_tilestate) GetTileStatus(tile env.ITile, dir env.Direction) Status {
	var x, y int

	x, y = tile.GetIndices()
	dx, dy := env.GetIndices(dir)
	x += dx
	y += dy

	if !env.ValidIndex(x, y) {
		return TILE_INVALID
	}

	if t.tiles[x][y].explored {
		if t.tiles[x][y].tile == nil {
			return TILE_WALL
		} else {
			return TILE_DISCOVERED
		}
	} else {
		return TILE_UNKOWN
	}
}
