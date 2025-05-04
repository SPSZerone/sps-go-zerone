// Package nerdfont provides the Nerd Font collection.
//
// Gen by `go run $(go env GOMODCACHE)/golang.org/x/image@v0.18.0/font/gofont/gen.go`
package nerdfont

import (
	"fmt"
	"sync"

	"gioui.org/font"
	"gioui.org/font/opentype"

	"github.com/SPSZerone/sps-go-zerone/font/nerdfont/meslolgsnerdfontmonobold"
	"github.com/SPSZerone/sps-go-zerone/font/nerdfont/meslolgsnerdfontmonobolditalic"
	"github.com/SPSZerone/sps-go-zerone/font/nerdfont/meslolgsnerdfontmonoitalic"
	"github.com/SPSZerone/sps-go-zerone/font/nerdfont/meslolgsnerdfontmonoregular"
)

var (
	regOnce    sync.Once
	reg        []font.FontFace
	once       sync.Once
	collection []font.FontFace
)

const (
	MesloLGSNerdFontMono = "MesloLGS Nerd Font Mono"
)

func loadRegular() {
	regOnce.Do(func() {
		faces, err := opentype.ParseCollection(meslolgsnerdfontmonoregular.TTF)
		if err != nil {
			panic(fmt.Errorf("failed to parse font: %v", err))
		}
		reg = faces
		collection = append(collection, reg[0])
	})
}

func Regular() []font.FontFace {
	loadRegular()
	return reg
}

func Collection() []font.FontFace {
	loadRegular()
	once.Do(func() {
		register(meslolgsnerdfontmonobold.TTF)
		register(meslolgsnerdfontmonoitalic.TTF)
		register(meslolgsnerdfontmonobolditalic.TTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(ttf []byte) {
	faces, err := opentype.ParseCollection(ttf)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	collection = append(collection, faces[0])
}
