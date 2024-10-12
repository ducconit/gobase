package utils

import (
	"context"
	"os"
	"os/signal"
)

func WaitOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()
	<-ctx.Done()
	f()
}

func RegisterOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	go WaitOSSignalHandler(f, signals...)
}

func RegisterSignalDefaultHandler(f func()) {
	RegisterOSSignalHandler(f, os.Interrupt, os.Kill)
}
