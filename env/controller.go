package env

const (
	MAX_SIZE = 8
)

type Controller struct {
	tiles     [MAX_SIZE][MAX_SIZE]ITile
	tileSlice []ITile

	permPos  uint64
	permDirt uint64
}

func (c *Controller) InitController(tileMap [MAX_SIZE][MAX_SIZE]bool) {

	for i := 0; i < MAX_SIZE; i++ {
		for j := 0; j < MAX_SIZE; j++ {
			if tileMap[i][j] {
				c.tiles[i][j] = new(t_tile)
			}
		}
	}

	c.joinTiles()
	c.initSlice()
}

func (c Controller) CanPermute(posIdx, dirtIdx uint64) bool {
	var count uint = uint(len(c.tileSlice))
	var upos uint = uint(posIdx)
	var uidx uint = uint(dirtIdx)

	return (count > (1<<upos) && count > uidx)
}

func (c *Controller) Permute(posIdx, dirtIdx int64) {
	var dirty bool
	var flag int64

	c.permPos = uint64(posIdx)
	c.permDirt = uint64(dirtIdx)

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
}

func (c Controller) GetStartingTile() ITile {
	return c.tileSlice[c.permPos]
}

func (c *Controller) joinTiles() {
	var thisTile ITile
	var neighbour ITile
	var dx, dy int

	for x := 0; x < MAX_SIZE; x++ {
		for y := 0; y < MAX_SIZE; y++ {
			// Check all neighbours
			for d := 0; d < 4; d++ {
				dx, dy = getIndices(d)
				if validIndex(x+dx, y+dy) {
					// If a neighbour exists at the direction,
					// join them together
					thisTile = c.tiles[x][y]
					neighbour = c.tiles[x+dx][y+dy]
					if thisTile != nil && neighbour != nil {
						thisTile.setNeighbour(d, neighbour)
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

func getIndices(dir int) (int, int) {
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

func validIndex(x, y int) bool {
	return x >= 0 && x < MAX_SIZE && y >= 0 && y < MAX_SIZE
}
