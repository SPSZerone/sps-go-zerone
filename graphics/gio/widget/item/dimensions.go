package item

func NewDimensions() Dimensions {
	return Dimensions{
		ContentWidth:       100,
		ContentHeight:      200,
		Padding:            4,
		HighlightThickness: 4,
		HighlightRoundness: 10,
	}
}

type Dimensions struct {
	ContentWidth       int
	ContentHeight      int
	Padding            int
	HighlightThickness int
	HighlightRoundness int
}
