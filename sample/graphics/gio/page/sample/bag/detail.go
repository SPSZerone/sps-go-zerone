package bag

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/widget/material"

	spscolor "github.com/SPSZerone/sps-go-zerone/graphics/gio/color"
	spslayout "github.com/SPSZerone/sps-go-zerone/graphics/gio/layout"
	spsdivider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/divider"
	spseditor "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/editor"
	spsslider "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/slider"
	spsspacer "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/spacer"
)

func (i *Item) LayoutDetail(
	theme *material.Theme, gtx layout.Context,
) layout.Dimensions {
	menuWidgets := []func(gtx layout.Context) layout.Dimensions{
		func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyId))
		},
		func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyName))
		},
		func(gtx layout.Context) layout.Dimensions {
			return i.LayoutDetailProperty(theme, gtx, i.Data.GetProperty(PropertyKeyIcon))
		},
	}
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
			return i.UI.GetMenuDetail(0).SetWidgets(menuWidgets...).Layout(
				theme, gtx,
				layout.W,
				image.Pt(gtx.Constraints.Max.X>>1, 0),
				func(gtx layout.Context) layout.Dimensions {
					return material.H6(theme, "Menu ...").Layout(gtx)
				},
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return i.UI.GetMenuDetail(1).SetWidgets(menuWidgets...).Layout(
				theme, gtx,
				layout.W,
				image.Pt(gtx.Constraints.Max.X>>2, 0),
				func(gtx layout.Context) layout.Dimensions {
					return material.H6(theme, "Menu ...Menu ...Menu ...Menu ...Menu ...Menu ...").Layout(gtx)
				},
			)
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
			if len(i.Suggests) == 0 {
				var colorIdx int
				nextColor := func() color.NRGBA {
					c := spscolor.DynamicColor(colorIdx)
					colorIdx++
					return c
				}
				for idx := 0; idx < 10; idx++ {
					content := fmt.Sprintf("Suggest %d", idx)
					suggest := spseditor.NewSuggest(
						content,
						func(gtx layout.Context) layout.Dimensions {
							return material.Body1(theme, fmt.Sprintf("%v display", content)).Layout(gtx)
						},
						spseditor.SgtOptBGColor(nextColor(), nextColor()),
					)
					i.Suggests = append(i.Suggests, &suggest)
				}
			}

			const suggestHeight = 32
			const suggestListHeight = suggestHeight * 3.5

			return i.UI.SuggestEditor.LayoutDefault(
				theme, gtx,
				suggestListHeight, suggestHeight,
				len(i.Suggests),
				func(gtx layout.Context, index int, size image.Point, onClick func(content string)) layout.Dimensions {
					suggest := i.Suggests[index]
					return suggest.Layout(theme, gtx, size, func() {
						onClick(fmt.Sprintf("%v Custom", suggest.Content))
					})
				},
			)
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
		return spsspacer.NewWithWidth(8).Layout(gtx)
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
		return spsspacer.NewWithWidth(8).Layout(gtx)
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
	return i.LayoutDetailProperty1(theme, gtx, property)
}

func (i *Item) LayoutDetailProperty1(theme *material.Theme, gtx layout.Context, property *Property) layout.Dimensions {
	return property.Layout(theme, gtx)
}

func (i *Item) LayoutDetailProperty2(theme *material.Theme, gtx layout.Context, property *Property) layout.Dimensions {
	ratio := float32(0.5)
	remainRatio := 1 - ratio
	return property.LayoutFlexWidgets(
		gtx,
		func() (weight float32, widget layout.Widget) {
			weight = ratio
			widget = func(gtx layout.Context) layout.Dimensions {
				return property.LayoutRigidKey(theme, gtx)
				//return property.LayoutFlexedKey(theme, gtx, 0.7)
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
