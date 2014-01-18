package env

type Controller struct {
	tiles [25][25]ITile
}

func (c *Controller) InitController(tileMap [25][25]bool) {

	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if tileMap[i][j] {
				c.tiles[i][j] = new(t_tile)
			}
		}
	}

	c.joinTiles()
}

func (c *Controller) joinTiles() {
	for x := 0; x < 25; x++ {
		for y := 0; y < 25; y++ {
			// Check all neighbours
			for d := 0; d < 4; d++ {
				dx, dy := getIndices(d)
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
	return x >= 0 && x < 25 && y >= 0 && y < 25
}
