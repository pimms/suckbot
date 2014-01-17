package env

import "testing"

func TestSetMap(t *testing.T) {
	var tileTestMap [25][25]bool

	tileTestMap[2][5] = true
	tileTestMap[3][5] = true
	tileTestMap[3][6] = true

	controller := new(Controller)
	controller.InitController(tileTestMap)

	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			var expNeighbours [4]bool = expectedNeighbours(tileTestMap, i, j)
			if !doesMeetExpectation(expNeighbours, controller.tiles[i][j]) {
				t.Error("Does not meet expectation")
			}
		}
	}
}

func doesMeetExpectation(check [4]bool, tiile Tile) bool {
	if tiile == nil {
		return true
	}

	for i := 0; i < 4; i++ {
		var hasNeighbour bool = tiile.GetNeighbour(i) != nil

		if hasNeighbour != check[i] {
			return false
		}
	}
	return true
}

func expectedNeighbours(tileMap [25][25]bool, x, y int) [4]bool {
	var res [4]bool
	for i := 0; i < 4; i++ {
		tx, ty := getIndices(i)
		if validIndex(x+tx, y+ty) {
			res[i] = tileMap[x+tx][y+ty]
		}
	}
	return res
}
