package gio

type NewWindow func(app App) Window

type Window interface {
	Run()
}
