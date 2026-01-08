package signal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func NotifyContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(ctx, ShutdownSignals()...)
}

func ShutdownSignals() []os.Signal {
	return []os.Signal{
		// SIGINT (Ctrl+C) | SIGTERM (Kill)
		syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGHUP, syscall.SIGQUIT,
	}
}
