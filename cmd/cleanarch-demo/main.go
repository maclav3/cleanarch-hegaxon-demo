package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/service"
	"github.com/pkg/errors"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	svc := service.NewService(ctx)
	svc.AddFixtures()

	svc.Logger.Info("Starting service...")
	err := svc.Run()
	if err != nil {
		panic(errors.Wrap(err, "error during service startup"))
	}
	svc.Logger.Info("Service started")

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		svc.Logger.Info("SIGINT caught, stopping")
		err = svc.Shutdown()
		if err != nil {
			panic(errors.Wrap(err, "error during service shutdown caused by SIGINT"))
		}
	}()

	<-svc.Running()
	svc.Logger.Info("Service stopped")
}
