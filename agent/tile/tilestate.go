package tile

import "github.com/pimms/suckbot/env"

/* The t_tilestate holds the data about the discovered
 * tiles and the status of the neighbours as well.
 *
 * Tiles are stored in an N by N array, and each tile has
 * it's indices associated with it.
 */
type t_tilestate struct {
	tiles [8][8]t_tilewrapper
}

/* Wrapper around an ITile and it's discovered status. If a
 * tile has been explored but it's tile member is nil, it's
 * a wall.
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
	return true
}
