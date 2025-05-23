package glfw

import (
	"gioui.org/gpu"
	"gioui.org/io/input"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/rs/zerolog"

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

	Theme *material.Theme
	Ops   op.Ops
	Queue input.Router

	Logger zerolog.Logger
}

func (c *Context) Release() {
	c.GPU.Release()
}
