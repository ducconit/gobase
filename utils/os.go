package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// RegisterOSSignalHandler registers a signal handler that runs in a goroutine.
// It waits for the specified signals and executes f when received.
func RegisterOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, signals...)
		<-ch
		close(ch)
		f()
	}()
}

// WaitOSSignalHandler waits for the specified OS signals and executes f when received.
func WaitOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	<-ch
	close(ch)
	f()
}

// WaitOSSignalGracefulShutdown waits for OS signals (Interrupt, SIGTERM, SIGQUIT) and executes graceful shutdown.
// It calls f with a context that has the specified timeout.
func WaitOSSignalGracefulShutdown(ctx context.Context, f func(ctx context.Context), timeout time.Duration) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch
	close(ch)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	f(ctx)
}
