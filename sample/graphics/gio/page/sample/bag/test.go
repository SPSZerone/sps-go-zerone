package bag

import (
	"fmt"

	"gioui.org/widget/material"
)

type PropertyKey int

const (
	PropertyKeyId PropertyKey = iota
	PropertyKeyName
	PropertyKeyIcon
	PropertyKeyMAX
)

func NewTestData(theme *material.Theme, groupCount, count int) TestData {
	testData := TestData{}
	testData.Items = make([][]Item, groupCount)
	for i := 0; i < groupCount; i++ {
		testData.Items[i] = make([]Item, count)
		for j := 0; j < count; j++ {
			data := NewData()
			data.AddProperties(NewTestProperties(i, j)...)
			testData.Items[i][j] = NewItem(data, theme)
		}
	}
	return testData
}

func NewTestProperties(i, j int) []Property {
	properties := make([]Property, PropertyKeyMAX)
	for key := PropertyKey(0); key < PropertyKeyMAX; key++ {
		properties[key] = NewProperty(GetTestKey(key), GetTestValue(key, i, j))
	}
	return properties
}

func GetTestKey(key PropertyKey) any {
	switch key {
	case PropertyKeyId:
		return "Id"
	case PropertyKeyName:
		return "Name"
	case PropertyKeyIcon:
		return "Icon"
	default:
	}
	return ""
}

func GetTestValue(key PropertyKey, i, j int) any {
	switch key {
	case PropertyKeyId:
		return fmt.Sprintf("id-%v-%v", i, j)
	case PropertyKeyName:
		return fmt.Sprintf("name-%v-%v", i, j)
	case PropertyKeyIcon:
		return fmt.Sprintf("icon-%v-%v", i, j)
	default:
	}
	return ""
}

type TestData struct {
	Items [][]Item
}

func (d *TestData) GetItems(index int) []Item {
	if index < 0 || index >= len(d.Items) {
		return nil
	}
	return d.Items[index]
}
