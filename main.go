package main

import (
	"fmt"
	"github.com/pimms/suckbot/agent"
	"github.com/pimms/suckbot/env"
)

func main() {
	fmt.Println("Hello, world!")

	controller := createController()
	a := new(agent.Agent)
	a.Initialize(controller.GetStartingTile())
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
	return controller
}
