package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	spsgl "github.com/SPSZerone/sps-go-zerone/graphics/gl"
	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

var (
	Logger = spslog.NewLogger()
)

func Run(opts ...Option) {
	o := NewOptions(opts...)

	spsgl.Require()

	// init pre
	if o.OnInitPre != nil {
		o.OnInitPre()
	}

	// glfw init
	err := glfw.Init()
	if err != nil {
		Logger.Fatal().Msgf("failed to initialize glfw: %v", err)
	}
	defer glfw.Terminate()

	InitWindowHint()
	if o.OnGLFWInit != nil {
		o.OnGLFWInit()
	}

	// new window
	win := o.NewWindow()
	window := win.GetWindow()

	// gl init
	spsgl.Init()
	if o.OnGLInit != nil {
		o.OnGLInit()
	}

	// init post
	if o.OnInitPost != nil {
		o.OnInitPost()
	}

	// stop
	defer func() {
		if o.OnStop != nil {
			o.OnStop()
		}
	}()

	// loop
	for !window.ShouldClose() {
		glfw.PollEvents()

		if o.OnLoop != nil {
			o.OnLoop(win)
		}

		window.SwapBuffers()
	}
}
