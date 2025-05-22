package main

import (
	spsglfw "github.com/SPSZerone/sps-go-zerone/graphics/glfw"
	spsglfwwin "github.com/SPSZerone/sps-go-zerone/graphics/glfw/window"
)

const (
	Name = "SPS GLFW Sample"
)

func main() {
	spsglfw.Run(newWin, loop)
}

func newWin() spsglfw.Window {
	return spsglfwwin.NewWindow(
		spsglfwwin.OptWidth(1920),
		spsglfwwin.OptHeight(1080),
		spsglfwwin.OptTitle(Name),
	)
}

func loop(win spsglfw.Window) {

}
