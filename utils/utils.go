package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func GracefulShutdown() {
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	<-signalChannel

	os.Exit(0)
}