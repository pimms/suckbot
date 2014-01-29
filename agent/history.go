package agent

import (
	"github.com/pimms/suckbot/env"
)

type t_state struct {
	action int
	tile   env.ITile
	next   *t_state
	prev   *t_state
}

type t_history struct {
	elems  []t_state
	length int
	head   *t_state
	tail   *t_state

	roundVisited map[env.ITile]bool
}

func maxHistory() int {
	return env.MAX_SIZE * env.MAX_SIZE * 2
}

func (h *t_history) hasCompletedRound(totalTiles int) bool {
	return len(h.roundVisited) == totalTiles
}

func (h *t_history) onNewRound() {
	h.roundVisited = make(map[env.ITile]bool)
}

func (h *t_history) addHistory(action int, tile env.ITile) {
	var node = new(t_state)
	node.action = action
	node.tile = tile

	h.roundVisited[tile] = true

	if h.length == 0 {
		h.head = node
		h.tail = node
		h.length++
	} else {
		if h.length >= maxHistory() {
			h.tail = h.tail.prev
		} else {
			h.length++
		}

		h.head.prev = node
		node.next = h.head
		h.head = node
	}
}

func (h *t_history) getStatistics(totalTiles int) (map[int]int, int) {
	var stat map[int]int = make(map[int]int)
	var unique map[*t_state]bool = make(map[*t_state]bool)
	var numActions int
	var node *t_state

	node = h.head
	for len(unique) < totalTiles && node != nil {
		unique[node] = true
		stat[node.action]++
		numActions++

		node = node.next
	}

	return stat, numActions
}

func (h *t_history) getActionPercent(action, totalTiles int) float64 {
	var stat map[int]int
	var numAct int

	stat, numAct = h.getStatistics(totalTiles)

	return float64(stat[action]) / float64(numAct)
}
