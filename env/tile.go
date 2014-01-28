package env

import (
	"math/rand"
)

type Direction int

const (
	/* DIRECTIONS */
	NONE  Direction = -1
	UP    Direction = 0
	RIGHT Direction = 1
	DOWN  Direction = 2
	LEFT  Direction = 3

	/* STATES */
	CLEAN = false
	DIRTY = true

	// After "CLEAN_MIN" iterations, there is a
	// "DIRTY_PERC"*100 chance of becoming
	// dirty.
	l_CLEAN_MIN  = 15
	l_DIRTY_PERC = 0.001
)

type TileState bool

/* Public interface for the Tile structures.
 * Provides only read-access to neighbours and states.
 */
type ITile interface {
	GetNeighbour(direction Direction) ITile
	GetState() TileState
	GetIndices() (int, int)
	OnVacuum()
	TimeSinceClean() int

	setState(state TileState)
	setNeighbour(direction Direction, neigh ITile) bool
	tick()
}

type t_tile struct {
	neighbours [4]ITile
	state      TileState

	xpos  int
	ypos  int
	timer int
}

/* Interface ITile implementation */
func (this *t_tile) GetState() TileState {
	return this.state
}

func (this *t_tile) GetNeighbour(direction Direction) ITile {
	if direction >= 0 && direction <= 3 {
		return this.neighbours[direction]
	}

	return nil
}

func (this *t_tile) GetIndices() (int, int) {
	return this.xpos, this.ypos
}

func (this *t_tile) OnVacuum() {
	this.setState(CLEAN)
}

func (this *t_tile) TimeSinceClean() int {
	if this.state != DIRTY {
		return 0
	}

	return this.timer
}

/* Private methods */
func (this *t_tile) setNeighbour(direction Direction, neigh ITile) bool {
	if direction >= 0 && direction <= 3 {
		this.neighbours[direction] = neigh

		var opposite Direction = Direction((int(direction) + 2) % 4)
		if neigh.GetNeighbour(opposite) != this {
			return neigh.setNeighbour(opposite, this)
		}

		return true
	}

	return false
}

func (this *t_tile) setState(state TileState) {
	this.state = state
	this.timer = 0
}

func (this *t_tile) tick() {
	this.timer++

	if this.timer > l_CLEAN_MIN && this.state == CLEAN {
		if rand.Float32() <= l_DIRTY_PERC {
			this.setState(DIRTY)
		}
	}
}
