package env

import (
	"fmt"
	"github.com/pimms/suckbot/util"
)

const (
	MAX_SIZE = 8
)

type Controller struct {
	tiles     [MAX_SIZE][MAX_SIZE]ITile
	tileSlice []ITile

	permPos  uint64
	permDirt uint64
}

// Peek function to allow a rendering object to draw
// the tiles.
func (c *Controller) CHEAT_GetTiles() [MAX_SIZE][MAX_SIZE]ITile {
	return c.tiles
}

func (c *Controller) InitController(tileMap [MAX_SIZE][MAX_SIZE]bool) {

	for i := 0; i < MAX_SIZE; i++ {
		for j := 0; j < MAX_SIZE; j++ {
			if tileMap[i][j] {
				var tt *t_tile = new(t_tile)
				tt.xpos = i
				tt.ypos = j

				c.tiles[i][j] = tt
			}
		}
	}

	c.joinTiles()
	c.initSlice()
}

func (c Controller) CanPermute(posIdx, dirtIdx uint64) bool {
	// The permutation number cannot exceed the absolute
	// maximum.
	if c.getPermNumber(posIdx, dirtIdx) >= c.getMaxPermCount() {
		return false
	}

	// the pos and dirt variables must still be valid, however.
	if dirtIdx >= (1 << uint64(len(c.tileSlice))) {
		return false
	}

	if posIdx >= uint64(len(c.tileSlice)) {
		return false
	}

	return true
}

func (c *Controller) Permute(posIdx, dirtIdx uint64) {
	var dirty bool
	var flag uint64

	c.permPos = posIdx
	c.permDirt = dirtIdx

	// Flag dirty / clean tiles
	for i := 0; i < len(c.tileSlice); i++ {
		flag = (dirtIdx & (1 << uint(i)))
		dirty = flag != 0

		if dirty {
			c.tileSlice[i].setState(DIRTY)
		} else {
			c.tileSlice[i].setState(CLEAN)
		}
	}

	fmt.Printf("[PERMUTATION  %d /  %d ]\n",
		c.getPermNumber(posIdx, dirtIdx)+1,
		c.getMaxPermCount())

}

func (c Controller) GetStartingTile() ITile {
	return c.tileSlice[c.permPos]
}

func (c *Controller) Tick() {
	for i := 0; i < len(c.tileSlice); i++ {
		c.tileSlice[i].tick()
	}
}

func (c *Controller) getPermNumber(posIdx, dirtIdx uint64) uint64 {
	return (posIdx)*(1<<uint64(len(c.tileSlice))) + dirtIdx
}

func (c *Controller) getMaxPermCount() uint64 {
	// The manual maximum (-1 if undefined)
	var maxParam int = util.MaxPermutations()
	if maxParam >= 0 {
		return uint64(maxParam)
	}

	// The physical maximum (N * 2^N)
	var nTiles = uint64(len(c.tileSlice))
	return nTiles * (1 << nTiles)
}

func (c *Controller) joinTiles() {
	var thisTile ITile
	var neighbour ITile
	var dx, dy int
	var direction Direction

	for x := 0; x < MAX_SIZE; x++ {
		for y := 0; y < MAX_SIZE; y++ {
			// Check all neighbours
			for d := 0; d < 4; d++ {
				direction = Direction(d)

				dx, dy = GetIndices(direction)
				if ValidIndex(x+dx, y+dy) {
					// If a neighbour exists at the direction,
					// join them together
					thisTile = c.tiles[x][y]
					neighbour = c.tiles[x+dx][y+dy]

					if thisTile != nil && neighbour != nil {
						thisTile.setNeighbour(direction, neighbour)
					}
				}
			}
		}
	}
}

func (c *Controller) initSlice() {
	var count int
	var idx int

	// Count the number of tiles
	for x := 0; x < MAX_SIZE; x++ {
		for y := 0; y < MAX_SIZE; y++ {
			if c.tiles[x][y] != nil {
				count++
			}
		}
	}

	// We know the length, create the slice
	c.tileSlice = make([]ITile, count)

	// Reference the tiles in a linear array
	for x := 0; x < MAX_SIZE; x++ {
		for y := 0; y < MAX_SIZE; y++ {
			if c.tiles[x][y] != nil {
				c.tileSlice[idx] = c.tiles[x][y]
				idx++
			}
		}
	}
}

func GetDirection(x, y int) Direction {
	if x == 0 && y == 1 {
		return UP
	}

	if x == 1 && y == 0 {
		return RIGHT
	}

	if x == 0 && y == -1 {
		return DOWN
	}

	if x == -1 && y == 0 {
		return LEFT
	}

	return NONE
}

func GetIndices(dir Direction) (int, int) {
	switch dir {
	case UP:
		return 0, 1
	case RIGHT:
		return 1, 0
	case DOWN:
		return 0, -1
	case LEFT:
		return -1, 0
	}

	return 0, 0
}

func ValidIndex(x, y int) bool {
	return x >= 0 && x < MAX_SIZE && y >= 0 && y < MAX_SIZE
}
