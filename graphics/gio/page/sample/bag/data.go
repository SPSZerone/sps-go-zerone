package bag

import (
	"time"

	spsgio "github.com/SPSZerone/sps-go-zerone/graphics/gio"
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bag"
)

func NewData(app *spsgio.Application) Data {
	return Data{
		Items: [][]bag.Item{
			{
				NewItem("Item A 1", app),
				NewItem("Item A 2", app),
				NewItem("Item A 3", app),
			},
			{
				NewItem("Item B 1", app),
				NewItem("Item B 2", app),
				NewItem("Item B 3", app),
			},
		},
		UpdateTime: time.Now(),
	}
}

type Data struct {
	Items      [][]bag.Item
	UpdateTime time.Time
}

func (d *Data) GetItems(index int) ([]bag.Item, time.Time) {
	if index < 0 || index >= len(d.Items) {
		return nil, d.UpdateTime
	}
	return d.Items[index], d.UpdateTime
}
