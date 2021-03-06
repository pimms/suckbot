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

/* The TileState holds the data about the discovered
 * tiles and the status of the neighbours as well.
 *
 * Tiles are stored in an N by N array, and each tile has
 * it's indices associated with it.
 */
type TileState struct {
	tiles [env.MAX_SIZE][env.MAX_SIZE]TileWrapper
}

/* Wrapper around an ITile and it's discovered status. If a
 * tile has been explored but it's tile member is nil, it's
 */
type TileWrapper struct {
	tile     env.ITile
	explored bool
}

func (t *TileWrapper) GetITile() env.ITile {
	return t.tile
}

/*
=======================
Implementation
=======================
*/
func (t *TileState) AddDiscovery(tile env.ITile) {
	var x, y int
	var twrap TileWrapper

	x, y = tile.GetIndices()

	// Create a new TileWrapper
	twrap.tile = tile
	t.tiles[x][y] = twrap
	t.tiles[x][y].explored = true
}

func (t *TileState) AddDiscoveryNil(x, y int) {
	if env.ValidIndex(x, y) {
		t.tiles[x][y].tile = nil
		t.tiles[x][y].explored = true
	}
}

func (t *TileState) GetTileStatus(tile env.ITile, dir env.Direction) Status {
	var x, y int

	x, y = tile.GetIndices()
	dx, dy := env.GetIndices(dir)
	x += dx
	y += dy

	return t.GetTileStatusAtCoord(x, y)
}

func (t *TileState) GetTileStatusAtCoord(x, y int) Status {
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

func (t *TileState) GetTile(tile *TileWrapper,
	dir env.Direction) *TileWrapper {

	dx, dy := env.GetIndices(dir)
	x, y := tile.tile.GetIndices()

	if env.ValidIndex(x+dx, y+dy) {
		return &t.tiles[x+dx][y+dy]
	}

	return nil
}

func (t *TileState) GetWrapper(tile env.ITile) *TileWrapper {
	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			if t.tiles[x][y].tile == tile {
				return &t.tiles[x][y]
			}
		}
	}

	return nil
}
