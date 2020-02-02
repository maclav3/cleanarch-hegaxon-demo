package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/nanomsg"

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
	client, err := nanomsg.NewClient(ctx, logger, "localhost:5555")
	if err != nil {
		logger.WithError(err).Error("Error creating nanomsg client")
		return
	}

	msg := strings.Join(os.Args[1:], " ")

	resp, err := client.Send([]byte(msg))
	if err != nil {
		logger.WithError(err).Error("Error sending message via nanomsg")
	}

	fmt.Println(string(resp))
}
