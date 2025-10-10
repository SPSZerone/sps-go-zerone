package image

type Op int

const (
	OpOriginally Op = iota
	OpFlipVertically
	OpFlipHorizontally
	OpFlipVerticallyHorizontally
)
