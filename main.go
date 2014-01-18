package main

import (
	"fmt"
	"github.com/pimms/suckbot/env"
)

func main() {
	fmt.Println("Hello, world!")

	var c env.Controller
	var b [25][25]bool
	c.InitController(b)
}

func createController() *env.Controller {
	var controller *env.Controller
	controller = new(env.Controller)

	var tileMap [25][25]bool
	tileMap[0][0] = true
	tileMap[1][0] = true
	tileMap[2][0] = true
	tileMap[1][1] = true

	controller.InitController(tileMap)
}
