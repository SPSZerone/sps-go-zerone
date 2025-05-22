package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/rs/zerolog"
)

type Window interface {
	GetWindow() *glfw.Window

	GetLogger() *zerolog.Logger
}
