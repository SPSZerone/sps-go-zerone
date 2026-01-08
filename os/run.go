package os

import (
	"fmt"
	"runtime/debug"

	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

type (
	OnRecover func(r any)
)

func RunSafe(fn func(), onRecover OnRecover) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}

			stack := string(debug.Stack())

			spslog.Logger.Error().Err(err).Msgf("[Recovery] panic -> %v\nStack trace:\n%s\n", err, stack)

			if onRecover != nil {
				onRecover(r)
			}
		}
	}()

	fn()
}
