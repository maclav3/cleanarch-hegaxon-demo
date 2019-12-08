package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	zmq "github.com/zeromq/gomq"
	"github.com/zeromq/gomq/zmtp"
)

func main() {
	logger := log.NewLogger("cleanarch-hexagon-demo-client")
	client := zmq.NewClient(zmtp.NewSecurityNull())
	defer client.Close()

	err := client.Connect("tcp://localhost:5555")
	if err != nil {
		logger.WithError(err).Error("could not connect to service")
		return
	}

	err = client.Send([]byte(strings.Join(os.Args[1:], " ")))
	if err != nil {
		logger.WithError(err).Error("error sending command to server")
		return
	}

	b, err := client.Recv()
	if err != nil {
		logger.WithError(err).Error("error receiving response from server")
		return
	}
	logger.Info("Received response from server")
	fmt.Println(string(b))
}
