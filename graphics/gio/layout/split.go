package layout

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	spsdrawing "github.com/SPSZerone/sps-go-zerone/graphics/gio/architecture/drawing"
	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
)

type ValueType byte

const (
	ValueTypeRatio ValueType = iota
	ValueTypeAbsolute
)

type Split struct {
	Flex layout.Flex

	Ratio   float32
	Bar     unit.Dp
	BarDraw func(gtx layout.Context, split *Split)

	bar     int
	halfBar int
	drag    bool
	dragID  pointer.ID
	dragPos float32
}

const defaultBarWidth = unit.Dp(10)

func (s *Split) GetBar() (int, int) {
	return s.bar, s.halfBar
}

func (s *Split) Layout(gtx layout.Context, aWidget, bWidget layout.Widget) layout.Dimensions {
	bar := gtx.Dp(s.Bar)
	if bar <= 1 {
		bar = gtx.Dp(defaultBarWidth)
	}
	halfBar := bar >> 1
	s.bar, s.halfBar = bar, halfBar

	// = aWidgetSize
	proportion := (s.Ratio + 1) / 2
	aWidgetSize := 0
	bWidgetSize := 0
	aWidgetMaxSize := 0
	switch s.Flex.Axis {
	case layout.Vertical:
		aWidgetSize = int(proportion*float32(gtx.Constraints.Max.Y) - float32(halfBar))
		aWidgetMaxSize = gtx.Constraints.Max.Y - bar
	default:
		aWidgetSize = int(proportion*float32(gtx.Constraints.Max.X) - float32(halfBar))
		aWidgetMaxSize = gtx.Constraints.Max.X - bar
	}
	if aWidgetSize < 0 {
		aWidgetSize = 0
	}
	if aWidgetSize > aWidgetMaxSize {
		aWidgetSize = aWidgetMaxSize
	}

	// = bWidgetOffset
	bWidgetOffset := aWidgetSize + bar

	// = bWidgetSize
	switch s.Flex.Axis {
	case layout.Vertical:
		bWidgetSize = gtx.Constraints.Max.Y - bWidgetOffset
	default:
		bWidgetSize = gtx.Constraints.Max.X - bWidgetOffset
	}

	{ // = handle input
		// = barRect
		var barRect image.Rectangle
		switch s.Flex.Axis {
		case layout.Vertical:
			barRect = image.Rect(0, aWidgetSize, gtx.Constraints.Max.X, bWidgetOffset)
		default:
			barRect = image.Rect(aWidgetSize, 0, bWidgetOffset, gtx.Constraints.Max.X)
		}
		barPos := image.Pt(barRect.Min.X, barRect.Min.Y)

		area := clip.Rect(barRect).Push(gtx.Ops)
		if s.BarDraw == nil {
			switch s.Flex.Axis {
			case layout.Vertical:
				BarPretty(gtx, s, barRect, barPos, 20, 10)
			default:
				BarPretty(gtx, s, barRect, barPos, 10, 20)
			}
		} else {
			s.BarDraw(gtx, s)
		}

		// = register for input
		event.Op(gtx.Ops, s)
		switch s.Flex.Axis {
		case layout.Vertical:
			pointer.CursorRowResize.Add(gtx.Ops)
		default:
			pointer.CursorColResize.Add(gtx.Ops)
		}

		for {
			ev, ok := gtx.Event(pointer.Filter{
				Target: s,
				Kinds:  pointer.Press | pointer.Drag | pointer.Release | pointer.Cancel,
			})
			if !ok {
				break
			}

			e, ok := ev.(pointer.Event)
			if !ok {
				continue
			}

			switch e.Kind {
			case pointer.Press:
				if s.drag {
					break
				}

				s.dragID = e.PointerID
				switch s.Flex.Axis {
				case layout.Vertical:
					s.dragPos = e.Position.Y
				default:
					s.dragPos = e.Position.X
				}
				s.drag = true

			case pointer.Drag:
				if s.dragID != e.PointerID {
					break
				}

				var posCur, posMax float32
				switch s.Flex.Axis {
				case layout.Vertical:
					posCur = e.Position.Y
					posMax = float32(gtx.Constraints.Max.Y)
				default:
					posCur = e.Position.X
					posMax = float32(gtx.Constraints.Max.X)
				}
				deltaPos := posCur - s.dragPos
				s.dragPos = posCur

				deltaRatio := deltaPos * 2 / posMax
				s.Ratio += deltaRatio

				if e.Priority < pointer.Grabbed {
					gtx.Execute(pointer.GrabCmd{
						Tag: s,
						ID:  s.dragID,
					})
				}

			case pointer.Release:
				fallthrough
			case pointer.Cancel:
				s.drag = false
			default:

			}
		}

		area.Pop()
	}

	{
		gtx := gtx
		switch s.Flex.Axis {
		case layout.Vertical:
			gtx.Constraints = layout.Exact(image.Pt(gtx.Constraints.Max.X, aWidgetSize))
		default:
			gtx.Constraints = layout.Exact(image.Pt(aWidgetSize, gtx.Constraints.Max.Y))
		}
		aWidget(gtx)
	}

	{
		var opOffset op.TransformStack
		bGtx := gtx
		switch s.Flex.Axis {
		case layout.Vertical:
			opOffset = op.Offset(image.Pt(0, bWidgetOffset)).Push(gtx.Ops)
			bGtx.Constraints = layout.Exact(image.Pt(bGtx.Constraints.Max.X, bWidgetSize))
		default:
			opOffset = op.Offset(image.Pt(bWidgetOffset, 0)).Push(gtx.Ops)
			bGtx.Constraints = layout.Exact(image.Pt(bWidgetSize, bGtx.Constraints.Max.Y))
		}
		bWidget(bGtx)
		opOffset.Pop()
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func BarPretty(gtx layout.Context, split *Split, bounds image.Rectangle, position image.Point, xColor, yColor int) {
	img := spsdrawing.NewImageNRGBADynamicColor(bounds, xColor, yColor)
	spsdrawing.DrawImage(gtx.Ops, img, paint.FilterNearest, f32.Pt(1, 1), position)

	bar, _ := split.GetBar()
	borderWidth := bar / 5
	borderHalfWidth := borderWidth >> 1
	borderSize := image.Pt(bounds.Max.X-bounds.Min.X-borderWidth, bounds.Max.Y-bounds.Min.Y-borderWidth)
	borderPos := image.Pt(position.X+borderHalfWidth, position.Y+borderHalfWidth)
	spsdrawing.DrawStrokeRectR(
		gtx.Ops,
		borderSize,
		color.NRGBA{A: 255, R: 161, G: 66, B: 244},
		borderPos,
		float32(borderWidth),
		0,
	)
}

func BarDynamicColor(gtx layout.Context, color1, color2 int) {
	spscolor.Fill(gtx, spscolor.DynamicColor(color1), spscolor.DynamicColor(color2))
}

func BarSimple(gtx layout.Context, color color.NRGBA) {
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
}
