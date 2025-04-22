package bag

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget/material"

	spscopy "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/copy"
	spsproperty "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/property"
)

func NewProperty(key, value any) Property {
	return Property{
		Property: spsproperty.New(),
		Key:      key,
		Value:    value,
	}
}

type Property struct {
	spsproperty.Property

	CopyKey   spscopy.Copy
	CopyValue spscopy.Copy

	Key   any
	Value any
}

func (p *Property) LayoutKey(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	content := fmt.Sprintf("%v", p.Key)
	return Layout(theme, gtx, content, &p.CopyKey)
}

func (p *Property) LayoutValue(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	content := fmt.Sprintf("%v", p.Value)
	return Layout(theme, gtx, content, &p.CopyValue)
}

func Layout(theme *material.Theme, gtx layout.Context, content string, copy *spscopy.Copy) layout.Dimensions {
	return copy.LayoutCopyRigidContent(
		gtx,
		func(gtx layout.Context) layout.Dimensions {
			return material.Button(theme, copy.GetClickable(), "Copy").Layout(gtx)
		},
		func() string {
			return content
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.H6(theme, content).Layout(gtx)
		},
	)
}
