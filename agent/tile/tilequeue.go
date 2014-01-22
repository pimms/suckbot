package tile

type TileQueue struct {
	elements []*TileWrapper
}

func (t *TileQueue) AddUnique(element *TileWrapper) {
	if t.elements == nil || len(t.elements) == 0 {
		t.elements = make([]*TileWrapper, 0)
	}

	for i := 0; i < len(t.elements); i++ {
		if t.elements[i] == element {
			return
		}
	}

	t.elements = append(t.elements, element)
}

func (t *TileQueue) MoveToBack(element *TileWrapper) {
	var toMove int = -1

	for i := 0; i < len(t.elements) && toMove < 0; i++ {
		if t.elements[i] == element {
			toMove = i
		}
	}

	if toMove >= 0 {
		e := t.elements[toMove]
		t.elements = append(t.elements[:toMove], t.elements[toMove+1:]...)
		t.elements = append(t.elements, e)
	}
}

func (t *TileQueue) GetHead() *TileWrapper {
	return t.elements[0]
}
