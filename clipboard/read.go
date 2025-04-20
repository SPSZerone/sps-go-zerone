package clipboard

import (
	"golang.design/x/clipboard"
)

func ReadTextString() string {
	return string(ReadText())
}

func ReadText() []byte {
	err := Init()
	if err != nil {
		return nil
	}
	return clipboard.Read(clipboard.FmtText)
}

func ReadImage() []byte {
	err := Init()
	if err != nil {
		return nil
	}
	return clipboard.Read(clipboard.FmtImage)
}
