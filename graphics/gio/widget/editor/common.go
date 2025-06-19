package editor

import (
	"gioui.org/widget/material"
)

type (
	OnSubmit func(editor *Editor)
	Style    func(editor *Editor, style *material.EditorStyle)
)
