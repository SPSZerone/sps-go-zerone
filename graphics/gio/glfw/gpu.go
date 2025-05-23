package glfw

import (
	"gioui.org/gpu"

	spsos "github.com/SPSZerone/sps-go-zerone/os"
)

func NewGPU() (gpu.GPU, error) {
	desktopGL := spsos.IsDarwin()
	return gpu.New(gpu.OpenGL{ES: !desktopGL, Shared: true})
}
