package item

import (
	"image"
)

type LayoutContext struct {
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
