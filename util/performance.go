package util

import (
	"github.com/pimms/suckbot/arg"
	"github.com/pimms/suckbot/env"
)

/* The score of the agent is calculated based on
 * the following rules:
 * - Minus one point for each move
 * - Plus three for cleaning a tile
 */

type SimPerf struct {
	// The total number of times the
	// agent moved
	agentMoves int

	// The total number of times
	// the agent cleaned a tile
	agentCleans int

	// The percentage of tiles the agent
	// walked into that were dirty
	dirtyEntry float64

	// One point is awarded for each clean tile
	// each tick
	cleanTicks int

	// The number of ticks a tile were dirty
	avgDirtyTicks float64
	minDirtyTicks int
	maxDirtyTicks int
}

func GetTotalScore(s SimPerf) float64 {
	var score float64 = 0.0

	score -= float64(s.agentMoves)
	score += float64(s.agentCleans) * 3.0

	return score
}

func GetAgentMoves(s SimPerf) float64 {
	return float64(s.agentMoves)
}

func GetAgentCleans(s SimPerf) float64 {
	return float64(s.agentCleans)
}

func GetDirtyEntries(s SimPerf) float64 {
	return float64(s.dirtyEntry)
}

func GetAvgDirtyTicks(s SimPerf) float64 {
	return float64(s.avgDirtyTicks)
}

func GetMinDirtyTicks(s SimPerf) float64 {
	return float64(s.minDirtyTicks)
}

func GetMaxDirtyTicks(s SimPerf) float64 {
	return float64(s.maxDirtyTicks)
}

func GetCleanTicks(s SimPerf) float64 {
	return float64(s.cleanTicks)
}

func (s *SimPerf) tileCleaned(tile env.ITile) {
	var time int = tile.TimeSinceClean()

	if s.minDirtyTicks == 0 || time < s.minDirtyTicks {
		s.minDirtyTicks = time
	}

	if s.maxDirtyTicks == 0 || time > s.maxDirtyTicks {
		s.maxDirtyTicks = time
	}

	var ftime = float64(time)
	s.avgDirtyTicks += ftime / float64(arg.NumRounds())
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

func (s *SimPerf) AgentCleaned(tile env.ITile) {
	s.tileCleaned(tile)
	s.agentCleans++
}

func (s *SimPerf) AgentEnteredTile(dirty bool) {
	if dirty {
		s.dirtyEntry += 1.0 / float64(arg.NumRounds())
	}
}

/*
====================
Controller methods

The following methods should only be called
by the Controller-instance.
====================
*/
func (s *SimPerf) SetCleanTilesThisTick(count int) {
	s.cleanTicks += count
}
