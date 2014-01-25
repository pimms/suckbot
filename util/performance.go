package util

// Interface to avoid import cycle
// with /env.
type PerfTile interface {
	TimeSinceClean() int
}

type SimPerf struct {
	totalScore int

	// The total number of times the
	// agent moved
	agentMoves int

	// The total number of times
	// the agent cleaned a tile
	agentCleans int

	// The percentage of tiles the agent
	// walked into that were dirty
	dirtyEntry float64

	// The number of ticks a tile were dirty
	avgDirtyTicks float64
	minDirtyTicks int
	maxDirtyTicks int
}

func (s SimPerf) GetTotalScore() int {
	return s.totalScore
}

func (s SimPerf) GetAgentMoves() int {
	return s.agentMoves
}

func (s SimPerf) GetAgentCleans() int {
	return s.agentCleans
}

func (s SimPerf) GetDirtyEntries() float64 {
	return s.dirtyEntry
}

func (s SimPerf) GetAvgDirtyTicks() float64 {
	return s.avgDirtyTicks
}

func (s SimPerf) GetMinDirtyTicks() int {
	return s.minDirtyTicks
}

func (s SimPerf) GetMaxDirtyTicks() int {
	return s.maxDirtyTicks
}

func (s *SimPerf) tileCleaned(tile PerfTile) {
	time := tile.TimeSinceClean()

	if s.minDirtyTicks == 0 || time < s.minDirtyTicks {
		s.minDirtyTicks = time
	}

	if s.maxDirtyTicks == 0 || time > s.maxDirtyTicks {
		s.maxDirtyTicks = time
	}

	var ftime = float64(time)
	s.avgDirtyTicks += ftime / float64(NumRounds())
}

/*
==================
Agent methods

Only an agent.Agent instance should call
the methods in this section.
==================
*/
func (s *SimPerf) AgentMoved() {
	s.agentMoves++
}

func (s *SimPerf) AgentCleaned(tile PerfTile) {
	s.tileCleaned(tile)
	s.agentCleans++
}

func (s *SimPerf) AgentEnteredTile(dirty bool) {
	if dirty {
		s.dirtyEntry += 1.0 / float64(NumRounds())
	}
}
