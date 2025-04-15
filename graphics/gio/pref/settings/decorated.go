package settings

import (
	"github.com/SPSZerone/sps-go-zerone/graphics/gio/widget/setting"
)

func NewDecorated() Decorated {
	return Decorated{
		Bool: setting.Bool{
			Name: "Decorated",
			Desc: "Use decorated",
		},
	}
}

type Decorated struct {
	setting.Bool
}
