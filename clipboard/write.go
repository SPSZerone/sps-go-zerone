package clipboard

import (
	"golang.design/x/clipboard"
)

func WriteTextString(content string) {
	WriteText([]byte(content))
}

func WriteText(bytes []byte) <-chan struct{} {
	err := Init()
	if err != nil {
		return nil
	}
	return clipboard.Write(clipboard.FmtText, bytes)
}

func WriteImage(bytes []byte) <-chan struct{} {
	err := Init()
	if err != nil {
		return nil
	}
	return clipboard.Write(clipboard.FmtImage, bytes)
}
