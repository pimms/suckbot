package main

import (
	"fmt"
	"github.com/pimms/suckbot/agent"
	"github.com/pimms/suckbot/arg"
	"github.com/pimms/suckbot/env"
	"github.com/pimms/suckbot/util"
	"time"
)

func main() {
	arg.BindArgs()
	var visual bool = arg.Visual()
	var rounds int = arg.NumRounds()
	var delay int = arg.DelayMS()

	var renderer t_renderer
	if visual {
		renderer.createWindow()
	}

	// Initialize the controller and the agent
	controller := createController()
	a := new(agent.Agent)
	a.Initialize(controller.GetStartingTile())

	var perfs []util.SimPerf
	perfs = make([]util.SimPerf, controller.GetMaxPermCount())

	var posPerm, dirtPerm uint64 = 0, 0
	for controller.CanPermute(posPerm, dirtPerm) {
		controller.Permute(posPerm, dirtPerm)
		permNo := controller.GetPermNumber(posPerm, dirtPerm)

		// Run the simulation
		for i := 0; i < rounds; i++ {
			if visual {
				renderer.pollEvents()
				if renderer.shouldExit {
					break
				}

				renderer.renderFrame(controller, a)
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}

			controller.Tick()
			a.Tick(&perfs[permNo])
		}

		if arg.Verbose() {
			printSimPerf(perfs[permNo])
		}

		// Increment the permutations - if no more permutations
		// are allowed with the incremented dirtPerm, increment
		// posPerm. If posPerm now holds an invalid value, the
		// embracing for-loop will terminate.
		dirtPerm++
		if !controller.CanPermute(posPerm, dirtPerm) {
			dirtPerm = 0
			posPerm++
		}
	}

	printSimPerfs(&perfs)
}

func printSimPerf(perf util.SimPerf) {
	s := make([]util.SimPerf, 1, 1)
	s[0] = perf
	printSimPerfs(&s)
}

func printSimPerfs(simPerf *[]util.SimPerf) {
	fmt.Printf("PERFORMANCE\n")
	fmt.Printf("Num simulations               -> %d\n", len(*simPerf))
	fmt.Printf("Ticks per simulation          -> %d\n", arg.NumRounds())

	fmt.Printf("ATTRIBUTE                     AVG       MIN       MAX\n")

	printPerfStat("Total score", simPerf, util.GetTotalScore)
	printPerfStat("Agent moves", simPerf, util.GetAgentMoves)
	printPerfStat("Agent cleans", simPerf, util.GetAgentCleans)
	printPerfStat("Dirty entry %", simPerf, util.GetDirtyEntries)
	printPerfStat("Dirty duration", simPerf, util.GetAvgDirtyTicks)

	print("\n\n")
}

func printPerfStat(context string,
	s *[]util.SimPerf, fp func(util.SimPerf) float64) {
	var min float64 = 1000000000.0
	var max float64 = -1000000000.0
	var avg float64 = 0.0

	var count = float64(len(*s))

	for i := 0; i < len(*s); i++ {
		var f = fp((*s)[i])

		if f < min {
			min = f
		}

		if f > max {
			max = f
		}

		avg += f / count
	}

	fmt.Printf("%s%s%s%s\n",
		util.StrW(context, 30),
		util.StrW(fmt.Sprintf("%0.3f", avg), 10),
		util.StrW(fmt.Sprintf("%0.3f", min), 10),
		util.StrW(fmt.Sprintf("%0.3f", max), 10))
}

func createController() *env.Controller {
	var controller *env.Controller
	controller = new(env.Controller)

	var tileMap [env.MAX_SIZE][env.MAX_SIZE]bool
	tileMap[0][0] = true
	tileMap[1][0] = true
	tileMap[2][0] = true
	tileMap[1][1] = true
	tileMap[1][2] = true
	tileMap[1][3] = true
	tileMap[0][3] = true
	tileMap[2][3] = true
	tileMap[2][2] = true
	tileMap[2][4] = true
	tileMap[2][5] = true
	tileMap[3][5] = true

	controller.InitController(tileMap)
	return controller
}
