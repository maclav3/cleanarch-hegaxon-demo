package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/port/cli"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-s
		cancel()
	}()

	logger := log.NewLogger("cleanarch-hexagon-demo-client")
	client, err := cli.NewClient(logger, "localhost:5555")
	if err != nil {
		logger.WithError(err).Error("error creating cli client")
		os.Exit(1)
	}

	resp, err := client.SendCmdFromArgs(ctx)
	if err != nil {
		logger.WithError(err).Error("error sending command")
		os.Exit(1)
	}

	fmt.Println(string(resp))
}
