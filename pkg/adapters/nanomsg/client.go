package nanomsg

import (
	"context"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/pkg/errors"

	"nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/req"

	_ "nanomsg.org/go/mangos/v2/transport/tcp"
)

type Client struct {
	cancelFn context.CancelFunc
	logger   log.Logger

	socket mangos.Socket
}

func NewClient(ctx context.Context, logger log.Logger, address string) (*Client, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	var socket mangos.Socket
	socket, err := req.NewSocket()
	if err != nil {
		return nil, errors.Wrap(err, "error creating socket")
	}

	go func() {
		<-ctx.Done()
		err = socket.Close()
		if err != nil {
			logger.WithError(err).Error("Error closing socket")
		}
	}()

	err = socket.Dial("tcp://" + address)
	if err != nil {
		return nil, errors.Wrap(err, "error dialing socket")
	}

	return &Client{
		cancelFn: cancel,
		logger:   logger,
		socket:   socket,
	}, nil
}

func (c *Client) Send(b []byte) ([]byte, error) {
	err := c.socket.Send(b)
	if err != nil {
		return nil, errors.Wrap(err, "error sending message")
	}

	msg, err := c.socket.RecvMsg()
	if err != nil {
		return nil, errors.Wrap(err, "error receiving response")
	}

	return msg.Body, nil
}

func (c *Client) Close() error {
	c.cancelFn()
	return nil
}
