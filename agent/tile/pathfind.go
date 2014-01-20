package tile

import (
	"github.com/pimms/suckbot/env"
	"github.com/pimms/suckbot/util"
)

type t_pathnode struct {
	parent *t_pathnode
	tile   *t_tilewrapper
	g      int
	h      int
}

func (t t_pathnode) f() int {
	return t.g + t.h
}

func manhattanDistance(a, b *t_pathnode) int {
	ax, ay := a.tile.tile.GetIndices()
	bx, by := b.tile.tile.GetIndices()
	return util.Absi(ax-bx) + util.Absi(ay-by)
}

/* Struct managing the variables required to
 * properly find a path
 */
type t_pf_state struct {
	open      []*t_pathnode
	closed    []*t_pathnode
	nodemap   [env.MAX_SIZE][env.MAX_SIZE]t_pathnode
	startNode *t_pathnode
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

func iselement(slice []*t_pathnode, element *t_pathnode) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			return true
		}
	}

	return false
}

func (t *t_pf_state) init(maxNodes int) {
	t.open = make([]*t_pathnode, 0, maxNodes)
	t.closed = make([]*t_pathnode, 0, maxNodes)
}

func (t *t_pf_state) setStart(start *t_tilewrapper) {
	startx, starty := start.tile.GetIndices()
	t.startNode = &t.nodemap[startx][starty]
	t.addToOpen(t.startNode)

	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			// Map the manhattan distance
			node := &t.nodemap[x][y]
			node.h = util.Absi(startx-x) +
				util.Absi(starty-y)
		}
	}
}

func (t *t_pf_state) setTilestate(state *t_tilestate) {
	for x := 0; x < env.MAX_SIZE; x++ {
		for y := 0; y < env.MAX_SIZE; y++ {
			t.nodemap[x][y].tile = &state.tiles[x][y]
		}
	}
}

func (t *t_pf_state) addToOpen(node *t_pathnode) {
	t.open = append(t.open, node)
}

func (t *t_pf_state) addToClosed(node *t_pathnode) {
	t.open = remove(t.open, node)
	t.closed = append(t.closed, node)
}

func (t *t_pf_state) isOpen(node *t_pathnode) bool {
	return iselement(t.open, node)
}

func (t *t_pf_state) isClosed(node *t_pathnode) bool {
	return iselement(t.closed, node)
}

func (t *t_pf_state) getPathnode(tilewrap *t_tilewrapper) *t_pathnode {
	x, y := tilewrap.tile.GetIndices()

	return &t.nodemap[x][y]
}

/* Returns the direction the agent should take in order to
 * successfully arrive at the end-tile.
 */
func PathFind(start, end *t_tilewrapper, tilestate *t_tilestate) env.Direction {
	var pfstate t_pf_state
	var node *t_pathnode
	var success bool

	pfstate.init(env.MAX_SIZE * env.MAX_SIZE)
	pfstate.setTilestate(tilestate)
	pfstate.setStart(start)

	node = pfstate.startNode

	for len(pfstate.open) != 0 {
		// Find the node with the lowest value for f()
		// from the open list
		node = pfstate.open[0]
		for i := 1; i < len(pfstate.open); i++ {
			if pfstate.open[i].f() < node.f() {
				node = pfstate.open[i]
			}
		}

		// Stop searching when we've added the destination
		// to the closed list
		pfstate.addToClosed(node)
		if node.tile == end {
			success = true
			break
		}

		for i := 0; i < 4; i++ {
			var dir env.Direction = env.Direction(i)
			var status Status = tilestate.GetTileStatus(node.tile.tile, dir)

			if status == TILE_DISCOVERED {
				var wrapper *t_tilewrapper
				var neighbour *t_pathnode

				wrapper = tilestate.GetTile(node.tile, dir)
				neighbour = pfstate.getPathnode(wrapper)

				if pfstate.isOpen(neighbour) {
					neighbour.parent = node
					neighbour.g = node.g
				} else if !pfstate.isClosed(neighbour) {
					pfstate.addToOpen(neighbour)
					neighbour.parent = node
					neighbour.g = node.g
				}
			}
		}
	}

	if success && node.parent != nil {
		for node.parent.parent != nil {
			node = node.parent
		}

		// Get the relative position
		x0, y0 := node.tile.tile.GetIndices()
		x1, y1 := node.parent.tile.tile.GetIndices()
		x := x0 - x1
		y := y0 - y1

		return env.GetDirection(x, y)
	}

	return env.NONE
}

func TileFind(base env.ITile, goal Status, context t_tilestate) env.Direction {
	return env.NONE
}
