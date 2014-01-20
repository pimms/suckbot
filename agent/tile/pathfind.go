package tile

import (
	"github.com/pimms/suckbot/env"
	"github.com/pimms/suckbot/util"
)

type t_pathnode struct {
	parent *t_pathnode
	tile   env.ITile
	g      int
	h      int
}

func (t t_pathnode) f() int {
	return t.g + t.h
}

func (t t_pathnode) manhattanDistance(other t_pathnode) int {
	tx, ty := t.tile.GetIndices()
	ox, oy := other.tile.GetIndices()
	return util.Absi(tx-ox) + util.Absi(ty-oy)
}

/* Struct managing the variables required to
 * properly find a path
 */
type t_pf_state struct {
	open   []*t_pathnode
	closed []*t_pathnode
}

func remove(slice []*t_pathnode, element *t_pathnode) []*t_pathnode {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			// Delete item number "i"
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}

	return slice
}

func (t *t_pf_state) init(maxNodes int) {
	t.open = make([]*t_pathnode, 0, maxNodes)
	t.closed = make([]*t_pathnode, 0, maxNodes)
}

func (t *t_pf_state) addToOpen(node *t_pathnode) {
	t.open = append(t.open, node)
}

func (t *t_pf_state) addToClosed(node *t_pathnode) {
	t.open = remove(t.open, node)
	t.closed = append(t.closed, node)
}

/* Returns the direction the agent should take in order to
 * successfully arrive at the end-tile.
 */
func PathFind(start, end env.ITile, heuristics bool) env.Direction {
	return env.NONE
}

func TileFind(base env.ITile, goal Status, context t_tilestate) env.Direction {
	return env.NONE
}
