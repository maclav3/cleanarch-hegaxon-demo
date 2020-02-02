package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/port/cli"

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
	// todo: define a proper client in the cli package
	client, err := nanomsg.NewClient(ctx, logger, "localhost:5555")
	if err != nil {
		logger.WithError(err).Error("Error creating nanomsg client")
		return
	}

	var msg bytes.Buffer
	for _, arg := range os.Args[1:] {
		_, err = msg.WriteString(arg)
		if err != nil {
			panic(err)
		}

		err = msg.WriteByte(cli.MessageBreak)
		if err != nil {
			panic(err)
		}
	}

	resp, err := client.Send(msg.Bytes())
	if err != nil {
		logger.WithError(err).Error("Error sending message via nanomsg")
	}

	fmt.Println(string(resp))
}
