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

type TileState int

/* Public interface for the Tile structures.
 * Provides only read-access to neighbours and states.
 */
type ITile interface {
	GetNeighbour(direction int) ITile
	GetState() TileState

	setNeighbour(direction int, neigh ITile) bool
}

type t_tile struct {
	neighbours [4]ITile
	state      TileState
}

/* Interface ITile implementation */
func (this *t_tile) GetState() TileState {
	return this.state
}

func (this *t_tile) GetNeighbour(direction int) ITile {
	if direction >= 0 && direction <= 3 {
		return this.neighbours[direction]
	}

	return nil
}

/* Private methods */
func (this *t_tile) setNeighbour(direction int, neigh ITile) bool {
	if direction >= 0 && direction <= 3 {
		this.neighbours[direction] = neigh

		var opposite int = (direction + 2) % 4
		if neigh.GetNeighbour(opposite) != this {
			return neigh.setNeighbour(opposite, this)
		}

		return true
	}

	return false
}
