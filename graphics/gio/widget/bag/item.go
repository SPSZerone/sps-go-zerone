package bag

import (
	"time"

	"gioui.org/layout"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

type ItemData func() (items []Item, itemUpdateTime time.Time)

type Item interface {
	GetItem() *spsitem.Item
	UpdateItem(item spsitem.Item)

	Layout(app *spsgio.Application, gtx layout.Context, highlight bool) (dimensions layout.Dimensions, clicked bool)
	LayoutContent(app *spsgio.Application, gtx layout.Context, item *spsitem.Item, layoutCtx spsitem.LayoutContext) layout.Dimensions
	LayoutDetail(app *spsgio.Application, gtx layout.Context, title string) layout.Dimensions
}
