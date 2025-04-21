package bag

import (
	"time"

	spsbag "github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/bag"
)

func (d *TestData) GetItems(index int) ([]spsbag.Item, time.Time) {
	if index < 0 || index >= len(d.Items) {
		return nil, d.UpdateTime
	}
	return d.Items[index], d.UpdateTime
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
