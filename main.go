package main

import (
	"github.com/pimms/suckbot/agent"
	"github.com/pimms/suckbot/env"
	"github.com/pimms/suckbot/util"
	"time"
)

func main() {
	util.BindArgs()
	var visual bool = util.Visual()
	var rounds int = util.NumRounds()
	var delay int = util.DelayMS()

	var renderer t_renderer
	if visual {
		renderer.createWindow()
	}

	// Initialize the controller and the agent
	controller := createController()
	a := new(agent.Agent)
	a.Initialize(controller.GetStartingTile())

	var posPerm, dirtPerm uint64 = 0, 0
	for controller.CanPermute(posPerm, dirtPerm) {
		controller.Permute(posPerm, dirtPerm)

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
			a.Tick(nil)
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
