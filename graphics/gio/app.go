package gio

import (
	"context"
)

type App interface {
	GetContext() context.Context
	Run(win Window)
	RunWindow(win Window)
}
