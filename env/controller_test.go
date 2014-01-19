package env

import "testing"

func createDummyController() *Controller {
	var tileTestMap [MAX_SIZE][MAX_SIZE]bool
	tileTestMap[0][0] = true
	tileTestMap[1][0] = true
	tileTestMap[2][0] = true
	tileTestMap[0][1] = true
	tileTestMap[1][1] = true
	tileTestMap[2][1] = true

	controller := new(Controller)
	controller.InitController(tileTestMap)

	return controller
}

func TestSetMap(t *testing.T) {
	// We need to validate with the bool map, so
	// createDummyController() cannot be used.
	var tileTestMap [MAX_SIZE][MAX_SIZE]bool
	tileTestMap[2][5] = true
	tileTestMap[3][5] = true
	tileTestMap[3][6] = true

	controller := new(Controller)
	controller.InitController(tileTestMap)

	for i := 0; i < MAX_SIZE; i++ {
		for j := 0; j < MAX_SIZE; j++ {
			var expNeighbours [4]bool = expectedNeighbours(tileTestMap, i, j)
			if !doesMeetExpectation(expNeighbours, controller.tiles[i][j]) {
				t.Error("Does not meet expectation")
			}
		}
	}
}

func doesMeetExpectation(check [4]bool, tiile ITile) bool {
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

func expectedNeighbours(tileMap [MAX_SIZE][MAX_SIZE]bool, x, y int) [4]bool {
	var res [4]bool
	for i := 0; i < 4; i++ {
		tx, ty := getIndices(i)
		if validIndex(x+tx, y+ty) {
			res[i] = tileMap[x+tx][y+ty]
		}
	}
	return res
}

func TestTileIndices(t *testing.T) {
	var cont *Controller = createDummyController()

	for x := 0; x < MAX_SIZE; x++ {
		for y := 0; y < MAX_SIZE; y++ {
			if cont.tiles[x][y] != nil {
				var ax, ay int
				ax, ay = cont.tiles[x][y].GetIndices()

				if ax != x || ay != y {
					t.Errorf("Invalid indices returned: "+
						"Expected [%i,%i], got [%i,%i]\n",
						x, y, ax, ay)
				}
			}
		}
	}
}

func TestTileCanPermute(t *testing.T) {
	var pos, dirt uint64
	var cont *Controller

	cont = createDummyController()

	// There are 6 tiles in the dummy controller. We should
	// thus be able to permute with position between 0 and 5,
	// and with the dirt between 0 and 63.
	for pos = 0; pos < 6; pos++ {
		if !cont.CanPermute(pos, dirt) {
			t.Errorf("Unable to permute with pos=%d and dirt=%d",
				pos, dirt)
		}
	}

	pos = 0
	for dirt = 0; dirt < 64; dirt++ {
		if !cont.CanPermute(pos, dirt) {
			t.Errorf("Unable to permute with pos=%d and dirt=%d",
				pos, dirt)
		}
	}

	// Ensure that pos >= 6 or dirt >= 64 fails
	if cont.CanPermute(6, 0) {
		t.Errorf("Able to permute with pos=6 and dirt=0")
	}

	if cont.CanPermute(0, 64) {
		t.Errorf("Able to permute with pos=0 and dirt=64")
	}
}
