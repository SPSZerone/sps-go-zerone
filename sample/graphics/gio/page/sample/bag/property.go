package bag

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	spscopy "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/copy"
	spsproperty "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/property"
)

func NewProperty(key, value any) Property {
	return Property{
		Property:  spsproperty.New(),
		CopyKey:   spscopy.New(),
		CopyValue: spscopy.New(),
		Key:       key,
		Value:     value,
	}
}

type Property struct {
	spsproperty.Property

	CopyKey   spscopy.Copy
	CopyValue spscopy.Copy

	Key   any
	Value any
}

func (p *Property) FormatKey() string {
	return fmt.Sprintf("%v", p.Key)
}

func (p *Property) FormatValue() string {
	return fmt.Sprintf("%s", p.Value)
}

func (p *Property) Layout(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return p.CopyValue.LayoutCopyFlexChild(
		theme,
		gtx,
		func() string {
			return p.FormatValue()
		},
		layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left:  unit.Dp(8),
				Right: unit.Dp(8),
			}.Layout(gtx, material.H6(theme, p.FormatKey()).Layout)
		}),
		layout.Flexed(0.8, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left:  unit.Dp(8),
				Right: unit.Dp(8),
			}.Layout(gtx, material.Body1(theme, p.FormatValue()).Layout)
		}),
	)
}

func (p *Property) LayoutRigidKey(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	content := fmt.Sprintf("%v", p.Key)
	return LayoutRigid(theme, gtx, &p.CopyKey, content)
}

func (p *Property) LayoutRigidValue(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	content := fmt.Sprintf("%v", p.Value)
	return LayoutRigid(theme, gtx, &p.CopyValue, content)
}

func (p *Property) LayoutFlexedKey(theme *material.Theme, gtx layout.Context, ratio float32) layout.Dimensions {
	content := fmt.Sprintf("%v", p.Key)
	return LayoutFlexed(theme, gtx, &p.CopyKey, ratio, content)
}

func (p *Property) LayoutFlexedValue(theme *material.Theme, gtx layout.Context, ratio float32) layout.Dimensions {
	content := fmt.Sprintf("%v", p.Value)
	return LayoutFlexed(theme, gtx, &p.CopyValue, ratio, content)
}

func LayoutFlexed(theme *material.Theme, gtx layout.Context, copy *spscopy.Copy, ratio float32, content string) layout.Dimensions {
	return copy.LayoutFlexedContent(
		gtx,
		1-ratio,
		func(gtx layout.Context) layout.Dimensions {
			return material.Button(theme, copy.GetClickable(), "Copy").Layout(gtx)
		},
		func() string {
			return content
		},
		func() (weight float32, widget layout.Widget) {
			weight = ratio
			widget = func(gtx layout.Context) layout.Dimensions {
				return material.H6(theme, content).Layout(gtx)
			}
			return
		},
	)
}

func LayoutRigid(theme *material.Theme, gtx layout.Context, copy *spscopy.Copy, content string) layout.Dimensions {
	return copy.LayoutRigidContent(
		gtx,
		func(gtx layout.Context) layout.Dimensions {
			return copy.LayoutCopyIcon(theme, gtx)
		},
		func() string {
			return content
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.H6(theme, content).Layout(gtx)
		},
	)
}
