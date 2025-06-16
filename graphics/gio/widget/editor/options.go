package editor

import (
	"image/color"
)

type Option func(*Editor)

func OptHint(value string) Option {
	return func(editor *Editor) {
		editor.Hint = value
	}
}

func OptSingleLine(value bool) Option {
	return func(editor *Editor) {
		editor.SingleLine = value
	}
}

func OptSubmit(value bool) Option {
	return func(editor *Editor) {
		editor.Submit = value
	}
}

func OptReadOnly(value bool) Option {
	return func(editor *Editor) {
		editor.ReadOnly = value
	}
}

type SuggestOption func(*Suggest)

func SgtOptBGColor(bgColor1, bgColor2 color.NRGBA) SuggestOption {
	return func(s *Suggest) {
		s.BGColor1, s.BGColor2 = bgColor1, bgColor2
	}
}
