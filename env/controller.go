package env

type Controller struct {
	tiles [25][25]Tile
}

func (c *Controller) InitController() {
	c.tiles[0][0] = new(tile)
	c.tiles[1][0] = new(tile)
	c.tiles[2][0] = new(tile)
	c.tiles[0][1] = new(tile)
	c.tiles[1][1] = new(tile)
	c.tiles[2][1] = new(tile)
	c.joinTiles()
}

func (c *Controller) joinTiles() {
	for x := 0; x < 25; x++ {
		for y := 0; y < 25; y++ {
			// Check all neighbours
			for d := 0; d < 4; d++ {
				dx, dy = getIndices(d)
				if validIndex(x+dx, y+dy) {
					// If a neighbour exists at the direction,
					// join them together
					thisTile := c.tiles[x][y]
					neighbour := c.tiles[x+dx][y+dy]
					if thisTile != nil && neighbour != nil {
						thisTile.setNeighbour(d, neighbour)
					}
				}
			}
		}
	}
}

func getIndices(dir int) (x, y int) {
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
	return x >= 0 && x < 25 && y >= 0 && y < 25
}
