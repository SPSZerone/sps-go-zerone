package bag

import (
	"fmt"
	"time"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func NewTestData(app *spsgio.Application, groupCount, count int) Data {
	d := Data{
		UpdateTime: time.Now(),
	}
	d.Items = make([][]spsitem.Item, groupCount)
	for i := 0; i < groupCount; i++ {
		d.Items[i] = make([]spsitem.Item, count)
		for j := 0; j < count; j++ {
			d.Items[i][j] = NewItem(fmt.Sprintf("Item %v %v", i, j), app)
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
