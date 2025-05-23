package glfw

import (
	"image"
	"time"

	"gioui.org/gpu"
	"gioui.org/io/input"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/rs/zerolog"

	"github.com/go-gl/glfw/v3.3/glfw"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

func NewDefaultContext() Context {
	logger := spslog.NewLogger()

	gpuCtx, err := NewGPU()
	if err != nil {
		logger.Fatal().Msgf("Error creating GPU context: %v", err)
	}

	return Context{
		GPU:    gpuCtx,
		Theme:  spsgio.NewTheme(),
		Logger: logger,
	}
}

func NewContext() Context {
	return Context{
		Logger: spslog.NewLogger(),
	}
}

type Context struct {
	GPU gpu.GPU

	Theme  *material.Theme
	Ops    op.Ops
	Router input.Router

	Logger zerolog.Logger
}

func (c *Context) LayoutContext(window *glfw.Window) (gtx layout.Context, size image.Point) {
	scale, _ := window.GetContentScale()
	width, height := window.GetFramebufferSize()
	size = image.Point{X: width, Y: height}
	c.Ops.Reset()
	gtx = layout.Context{
		Ops:    &c.Ops,
		Now:    time.Now(),
		Source: c.Router.Source(),
		Metric: unit.Metric{
			PxPerDp: scale,
			PxPerSp: scale,
		},
		Constraints: layout.Exact(size),
	}
	return
}

func (c *Context) Frame(gtx layout.Context, size image.Point) {
	c.GPU.Frame(gtx.Ops, gpu.OpenGLRenderTarget{}, size)
	c.Router.Frame(gtx.Ops)
}

func (c *Context) Release() {
	c.GPU.Release()
}
