package nanomsg

import (
	"context"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/pkg/errors"

	"nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/rep"

	_ "nanomsg.org/go/mangos/v2/transport/tcp"
)

type Server struct {
	cancelFn context.CancelFunc
	logger   log.Logger
	socket   mangos.Socket
}

func NewServer(ctx context.Context, logger log.Logger, address string) (*Server, error) {
	ctx, cancel := context.WithCancel(ctx)

	var socket mangos.Socket
	socket, err := rep.NewSocket()
	if err != nil {
		return nil, errors.Wrap(err, "error creating socket")
	}

	go func() {
		<-ctx.Done()
		err = socket.Close()
		if err != nil {
			logger.WithError(err).Error("Error opening socket")
		}
	}()

	err = socket.Listen("tcp://" + address)
	if err != nil {
		return nil, errors.Wrap(err, "Error dialing socket")
	}

	return &Server{
		cancelFn: cancel,
		logger:   logger,
		socket:   socket,
	}, nil
}

func (s *Server) Listen() (<-chan []byte, error) {
	respCh := make(chan []byte)
	go s.receive(respCh)
	return respCh, nil
}

func (s *Server) receive(respCh chan []byte) {
	for {
		msg, err := s.socket.Recv()
		if err != nil {
			s.logger.WithError(err).Error("Error receiving message")
		}

		respCh <- msg
	}
}

func (s *Server) Send(b []byte) error {
	return s.socket.Send(b)
}

func (s *Server) Close() error {
	s.cancelFn()
	return nil
}
