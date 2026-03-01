package errgroup

import (
	"context"
	"fmt"
	"runtime/debug"

	"golang.org/x/sync/errgroup"

	spslog "github.com/SPSZerone/sps-go-zerone/log/zerolog"
)

type (
	OnRecover func(r any)
)

func SafeGo(g *errgroup.Group, ctx context.Context, fn func() error, onRecover OnRecover) {
	g.Go(func() error {
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

		if ctx.Err() != nil {
			return ctx.Err()
		}

		return fn()
	})
}
