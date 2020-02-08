package cli

import (
	"bytes"
	"context"
	"os"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/nanomsg"

	"github.com/pkg/errors"
)

type Client struct {
	logger log.Logger
	addr   string
}

func NewClient(logger log.Logger, addr string) (*Client, error) {
	if logger == nil {
		return nil, errors.New("logger is nil")
	}

	return &Client{
		logger: logger,
		addr:   addr,
	}, nil
}

// SendCmdFromArgs sends a command parsed from the os.Args, allowing to call the client like a cobra Command.
func (c *Client) SendCmdFromArgs(ctx context.Context) ([]byte, error) {
	client, err := nanomsg.NewClient(ctx, c.logger, c.addr)
	if err != nil {
		return nil, errors.Wrap(err, "error creating nanomsg client")
	}
	defer func() {
		err = client.Close()
		if err != nil {
			c.logger.WithError(err).Error("Error closing nanomsg client")
		}
	}()

	var msg bytes.Buffer
	for _, arg := range os.Args[1:] {
		_, err = msg.WriteString(arg)
		if err != nil {
			panic(err)
		}

		err = msg.WriteByte(argsSeparator)
		if err != nil {
			panic(err)
		}
	}

	resp, err := client.Send(msg.Bytes())
	if err != nil {
		c.logger.WithError(err).Error("Error sending message via nanomsg")
	}
	return resp, nil
}
