package slider

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
)

func NewSlider(opts ...Option) Slider {
	s := Slider{
		Min: 0,
		Max: 1,
	}
	s.Update(opts...)
	return s
}

type OnValueChanged func(slider *Slider)
type Style func(slider *Slider, style *material.SliderStyle)

type Slider struct {
	widget.Float
	Min, Max float32

	Name string
	Desc string
}

func (s *Slider) Update(opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *Slider) Layout(
	theme *material.Theme, gtx layout.Context,
	valueInFront bool, ratioInFront float32,
	style Style,
	onValueChanged OnValueChanged,
) layout.Dimensions {
	key := material.Body1(theme, fmt.Sprintf("%v %v [%v,%v]", s.Name, s.Float.Value, s.Min, s.Max)).Layout
	value := func(gtx layout.Context) layout.Dimensions {
		return s.LayoutSlider(theme, gtx, style, onValueChanged)
	}
	var aWidget, bWidget layout.Widget
	if valueInFront {
		aWidget, bWidget = value, key
	} else {
		aWidget, bWidget = key, value
	}
	return spslayout.FlexInset{Ratio: ratioInFront}.LayoutABWidget(gtx, aWidget, bWidget)
}

func (s *Slider) LayoutSlider(
	theme *material.Theme, gtx layout.Context,
	style Style,
	onValueChanged OnValueChanged,
) layout.Dimensions {
	if s.Float.Update(gtx) {
		if s.Float.Value < s.Min {
			s.Float.Value = s.Min
		}
		if s.Float.Value > s.Max {
			s.Float.Value = s.Max
		}
		if onValueChanged != nil {
			onValueChanged(s)
		}
	}
	sliderStyle := material.Slider(theme, &s.Float)
	if style != nil {
		style(s, &sliderStyle)
	}
	return sliderStyle.Layout(gtx)
}

func (s *Slider) UpdateMinMax(min, max float32) {
	s.Min = min
	s.Max = max
}
