package item

import (
	"image"
)

type LayoutContext struct {
	ContentWidth  int
	ContentHeight int

	Padding       int
	PaddingDouble int

	HighlightThickness       int
	HighlightThicknessDouble int
	HighlightThicknessHalf   int

	HighlightRoundness int

	Offset int
	Size   image.Point
}
