package bag

import (
	"fmt"
	"time"

	"gioui.org/widget/material"

	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func NewTestData(theme *material.Theme, groupCount, count int) Data {
	d := Data{
		UpdateTime: time.Now(),
	}
	d.Items = make([][]spsitem.Item, groupCount)
	for i := 0; i < groupCount; i++ {
		d.Items[i] = make([]spsitem.Item, count)
		for j := 0; j < count; j++ {
			d.Items[i][j] = NewItem(fmt.Sprintf("Item %v %v", i, j), theme)
		}
	}
	return d
}

type Data struct {
	Items      [][]spsitem.Item
	UpdateTime time.Time
}

func (d *Data) GetItems(index int) ([]spsitem.Item, time.Time) {
	if index < 0 || index >= len(d.Items) {
		return nil, d.UpdateTime
	}
	return d.Items[index], d.UpdateTime
}
