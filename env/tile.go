package env

const (
	/* DIRECTIONS */
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3

	/* STATES */
	CLEAN = 0
	DIRTY = 1
)

type Direction int32
type TileState int32

/* Public interface for the Tile structures.
 * Provides only read-access to neighbours and states.
 */
type Tile interface {
	GetNeighbour(direction Direction) *Tile
	GetState() TileState
}
