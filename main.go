package main

import (
	"flag"
	"github.com/pimms/suckbot/agent"
	"github.com/pimms/suckbot/env"
	"time"
)

func main() {
	var visual = flag.Bool("visual", false,	"Visualize the agent")
	var rounds = flag.Int("rounds", 1000, "The number of rounds to simulate")
	flag.Parse()

	var renderer t_renderer
	if *visual {
		renderer.createWindow()
	}

	// Initialize the controller and the agent
	controller := createController()
	a := new(agent.Agent)
	a.Initialize(controller.GetStartingTile())

	for i := 0; i < *rounds; i++ {
		if *visual {
			renderer.pollEvents()
			if renderer.shouldExit {
				break
			}

			renderer.renderFrame(controller, a)
			time.Sleep(500 * time.Millisecond)
		}

		a.Tick()
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

	controller.InitController(tileMap)
	controller.Permute(1, 14)
	return controller
}
