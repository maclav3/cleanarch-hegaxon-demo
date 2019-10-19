package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/service"
	"github.com/pkg/errors"
)

func main() {
	ctx := context.Background()
	svc := service.NewService(ctx)

	svc.Logger.Info("Starting service...")
	err := svc.Run(ctx)
	if err != nil {
		panic(errors.Wrap(err, "error during service startup"))
	}
	svc.Logger.Info("Service started")

	svc.AddFixtures()

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-s
		svc.Logger.Info("SIGTERM/SIGINT caught, stopping")
		err = svc.Shutdown()
		if err != nil {
			panic(errors.Wrap(err, "error during service shutdown caused by SIGINT"))
		}
	}()

	<-svc.Running()
	svc.Logger.Info("Service stopped")
}
