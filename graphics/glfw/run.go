package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

var (
	Logger = spslog.NewLogger()
)

func Run(
	newWin func() Window,
	loop func(win Window),
) {
	require()

	err := glfw.Init()
	if err != nil {
		Logger.Fatal().Msgf("failed to initialize glfw: %v", err)
	}
	defer glfw.Terminate()

	glfwInitWindowHint()

	win := newWin()
	window := win.GetWindow()

	glInit()

	for !window.ShouldClose() {
		glfw.PollEvents()

		if loop != nil {
			loop(win)
		}

		window.SwapBuffers()
	}
}
