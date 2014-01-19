package env

const (
	/* DIRECTIONS */
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3

	/* STATES */
	CLEAN = false
	DIRTY = true
)

type TileState bool

/* Public interface for the Tile structures.
 * Provides only read-access to neighbours and states.
 */
type Tile interface {
	GetNeighbour(direction int) Tile
	GetState() TileState

	setNeighbour(direction int, neigh Tile) bool
	setState(state TileState)
}

/* The Tile implementation */
type tile struct {
	neighbours [4]Tile
	state      TileState
}

/* Interface Tile implementation */
func (this *tile) GetState() TileState {
	return this.state
}

func (this *tile) GetNeighbour(direction int) Tile {
	if direction >= 0 && direction <= 3 {
		return this.neighbours[direction]
	}

	return nil
}

/* Private methods */
func (this *tile) setNeighbour(direction int, neigh Tile) bool {
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

func (this *tile) setState(state TileState) {
	this.state = state
}
