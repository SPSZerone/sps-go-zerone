package item

import (
	"image"

	"gioui.org/layout"
)

type LayoutContext struct {
	InsetOuter  layout.Inset
	InsetInner  layout.Inset
	ContentSize image.Point

	Padding       int
	PaddingDouble int

	HighlightThickness       int
	HighlightThicknessDouble int
	HighlightThicknessHalf   int

	HighlightRoundness int

	Offset int
	Size   image.Point
}
