package clipboard

import (
	"sync"

	"golang.design/x/clipboard"
)

var (
	once sync.Once
)

func Init() (err error) {
	once.Do(func() {
		err = clipboard.Init()
	})
	return
}
