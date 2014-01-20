package tile

import "testing"

func TestRemove(t *testing.T) {
	item0 := new(t_pathnode)
	item1 := new(t_pathnode)

	slice := make([]*t_pathnode, 2, 2)
	slice[0] = item0
	slice[1] = item1

	slice = remove(slice, item0)

	if len(slice) != 1 {
		t.Errorf("Expected length 1, %d received\n", len(slice))
	}

	if slice[0] != item1 {
		t.Errorf("Expected 'item1' as sole element\n")
	}
}

func TestAddToOpen(t *testing.T) {
	itemOpen := new(t_pathnode)

	var state t_pf_state
	state.init(8)

	state.addToOpen(itemOpen)

	if len(state.open) != 1 {
		t.Errorf("Unexpected length: %d (expected 1)\n",
			len(state.open))
	}

	if state.open[0] != itemOpen {
		t.Errorf("I don't even...\n")
	}
}

func TestAddToClosed(t *testing.T) {
	node := new(t_pathnode)

	var state t_pf_state
	state.init(1)

	state.addToOpen(node)
	state.addToClosed(node)

	if len(state.open) != 0 {
		t.Errorf("Node not removed from open list\n")
	}

	if len(state.closed) != 1 {
		t.Error("Node not added to closed list")
	}

	if state.closed[0] != node {
		t.Errorf("The closed node it's not my node. " +
			"It's just some node who thinks 'closed' is it's list.")
	}
}
