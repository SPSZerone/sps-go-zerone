package bag

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/widget/material"

	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spsslider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/slider"
	spsspacer "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/spacer"
	spssurface "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/surface"
)

func (i *Item) LayoutDetail(
	theme *material.Theme, gtx layout.Context,
) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			baseInfo := material.H5(theme, "Item Detail")
			baseInfo.Font.Style = font.Italic
			baseInfo.Font.Weight = font.Bold
			return spslayout.DefaultInset.Layout(gtx, baseInfo.Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spssurface.NewSurface().Layout(theme, gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H6(theme, "Use ...").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spsdivider.Divider{}.Layout(theme, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutUseEditor(theme, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutUseSlider(theme, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H5(theme, "Info").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return spsdivider.Divider{}.Layout(theme, gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyId))
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyName))
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyIcon))
		}),
	)
}

func (i *Item) LayoutUseEditor(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	spacer := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return spsspacer.NewWidthSpacer(8).Layout(gtx)
	})
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		spacer,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return i.UI.LayoutUseEditor(theme, gtx)
		}),
		spacer,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.UI.LayoutUseEditorBtn(theme, gtx)
		}),
		spacer,
	)
}

func (i *Item) LayoutUseSlider(theme *material.Theme, gtx layout.Context) layout.Dimensions {
	spacer := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return spsspacer.NewWidthSpacer(8).Layout(gtx)
	})
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		spacer,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return i.UI.LayoutUseSlider(
				theme, gtx,
				func(slider *spsslider.Slider, style *material.SliderStyle) {

				},
				func(slider *spsslider.Slider) {

				},
			)
		}),
		spacer,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.UI.LayoutUseSliderBtn(theme, gtx)
		}),
		spacer,
	)
}

func (i *Item) LayoutDetailProperty(theme *material.Theme, gtx layout.Context, property *Property) layout.Dimensions {
	ratio := float32(0.2)
	remainRatio := 1 - ratio
	return property.LayoutFlexWidgets(
		gtx,
		func() (weight float32, widget layout.Widget) {
			weight = remainRatio
			widget = func(gtx layout.Context) layout.Dimensions {
				//return property.LayoutRigidKey(theme, gtx)
				return property.LayoutFlexedKey(theme, gtx, 0.7)
			}
			return
		},
		func() (weight float32, widget layout.Widget) {
			weight = remainRatio
			widget = func(gtx layout.Context) layout.Dimensions {
				//return property.LayoutRigidValue(theme, gtx)
				return property.LayoutFlexedValue(theme, gtx, 0.7)
			}
			return
		},
	)
}
