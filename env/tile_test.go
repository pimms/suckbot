package env

import "testing"

func TestValidDirections(t *testing.T) {
	var myTile *t_tile = new(t_tile)
	for i := 0; i < 4; i++ {
		myTile.neighbours[i] = new(t_tile)
	}

	// Ensure that 0-4 are non-nil
	for i := 0; i < 4; i++ {
		var dir = Direction(i)
		n := myTile.GetNeighbour(dir)

		if n == nil {
			t.Errorf("Unexpected nil neighbour: %i\n", i)
		}
	}

	// Ensure that anything else is nil
	neg := myTile.GetNeighbour(-1)
	if neg != nil {
		t.Errorf("Non nil neighbour for -1")
	}

	pos := myTile.GetNeighbour(5)
	if pos != nil {
		t.Errorf("Non nil neighbour for 5")
	}
}

func TestSetAndGet(t *testing.T) {
	var myTile *t_tile = new(t_tile)

	var newTile = new(t_tile)
	result := myTile.setNeighbour(0, newTile)
	if !result {
		t.Error("Unexpected failure of setting neighbour")
	}

	// Ensure that they are each other`s neighbours
	if myTile.GetNeighbour(0) != newTile {
		t.Error("Bad neighbour")
	}

	if newTile.GetNeighbour(2) != myTile {
		t.Error("Bad neighbour")
	}
}
