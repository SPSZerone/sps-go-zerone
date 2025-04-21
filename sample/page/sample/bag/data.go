package bag

import (
	"time"

	spsitem "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/item"
)

func (d *TestData) GetItems(index int) ([]spsitem.Item, time.Time) {
	if index < 0 || index >= len(d.Items) {
		return nil, d.UpdateTime
	}
	items := make([]spsitem.Item, len(d.Items[index]))
	for i := 0; i < len(d.Items[index]); i++ {
		items[i] = d.Items[index][i].Item
	}
	return items, d.UpdateTime
}

func NewData(id any, name string) Data {
	return Data{
		Id:   id,
		Name: name,
	}
}

type Data struct {
	Id   any
	Name string
}
