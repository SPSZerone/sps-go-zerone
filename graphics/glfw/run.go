package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	spsgl "github.com/SPSZerone/sps-go-zerone/graphics/gl"
	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

var (
	Logger = spslog.NewLogger()
)

func Run(opts ...Option) (err error) {
	o := NewOptions(opts...)

	spsgl.Require()

	// init pre
	if o.OnInitPre != nil {
		err = o.OnInitPre()
		if err != nil {
			Logger.Error().Msgf("failed OnInitPre: %v", err)
			return
		}
	}

	// glfw init
	err = glfw.Init()
	if err != nil {
		Logger.Error().Msgf("failed glfw.Init: %v", err)
		return
	}
	defer glfw.Terminate()

	InitWindowHint()
	if o.OnGLFWInit != nil {
		err = o.OnGLFWInit()
		if err != nil {
			Logger.Error().Msgf("failed OnGLFWInit: %v", err)
			return
		}
	}

	// new window
	win := o.NewWindow()
	window := win.GetWindow()

	// gl init
	spsgl.Init()
	if o.OnGLInit != nil {
		err = o.OnGLInit()
		if err != nil {
			Logger.Error().Msgf("failed OnGLInit: %v", err)
			return
		}
	}

	// init post
	if o.OnInitPost != nil {
		err = o.OnInitPost()
		if err != nil {
			Logger.Error().Msgf("failed OnInitPost: %v", err)
			return
		}
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

	return
}
