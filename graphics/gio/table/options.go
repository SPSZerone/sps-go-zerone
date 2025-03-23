package table

import (
	"gioui.org/layout"
	"gioui.org/widget"
)

type Option func(*Table)

func OptHeaders(value ...Header) Option {
	return func(t *Table) {
		t.Headers = value
	}
}

func OptMinSize(value float32) Option {
	return func(t *Table) {
		t.MinSize = value
	}
}

func OptBorder(value widget.Border) Option {
	return func(t *Table) {
		t.Border = value
	}
}

func OptInset(value layout.Inset) Option {
	return func(t *Table) {
		t.Inset = value
	}
}

func OptHeaderLabelStyle(value LabelStyle) Option {
	return func(t *Table) {
		t.HeaderLabelStyle = value
	}
}

func OptDataLabelStyle(value LabelStyle) Option {
	return func(t *Table) {
		t.DataLabelStyle = value
	}
}
