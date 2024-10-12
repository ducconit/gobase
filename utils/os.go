package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	<-ch

	f()
}

func RegisterOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	go WaitOSSignalHandler(f, signals...)
}

func RegisterSignalDefaultHandler(f func()) {
	RegisterOSSignalHandler(f, syscall.SIGINT, syscall.SIGTERM)
}
