package bag

import (
	"fmt"
	"time"

	"gioui.org/widget/material"

	spsbag "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bag"
)

func NewTestData(theme *material.Theme, groupCount, count int) TestData {
	d := TestData{
		UpdateTime: time.Now(),
	}
	d.Items = make([][]spsbag.Item, groupCount)
	for i := 0; i < groupCount; i++ {
		d.Items[i] = make([]spsbag.Item, count)
		for j := 0; j < count; j++ {
			id := fmt.Sprintf("%v-%v", i, j)
			name := fmt.Sprintf("n-%v-%v", i, j)
			d.Items[i][j] = NewItem(NewData(id, name), theme)
		}
	}
	return d
}

type TestData struct {
	Items      [][]spsbag.Item
	UpdateTime time.Time
}
