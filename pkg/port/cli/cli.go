package cli

import (
	"context"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/nanomsg"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Router struct {
	cancelFn func()
	running  chan struct{}
	rootCmd  *cobra.Command
	logger   log.Logger
	address  string
	server   *nanomsg.Server
}

func NewRouter(logger log.Logger, address string) *Router {
	rootCmd := &cobra.Command{
		Use:   "cleanarch-demo",
		Short: "Execute application commands/queries from CLI",
		Long: "The cleanarch-demo CLI port is called from the command line." +
			"It parses arguments from the command line, " +
			"executes application commands/queries and marshals the output to stdout.",
	}
	return &Router{
		rootCmd: rootCmd,
		logger:  logger,
		address: address,
	}
}

// Run starts the Router router and keeps parsing commands from the standard input.
func (r *Router) Run(ctx context.Context) error {
	ctx, r.cancelFn = context.WithCancel(ctx)

	r.running = make(chan struct{})
	defer close(r.running)

	var err error
	r.server, err = nanomsg.NewServer(ctx, r.logger, r.address)
	if err != nil {
		return errors.Wrap(err, "Could not run nanomsg server")
	}

	msgCh, err := r.server.Listen()
	if err != nil {
		return errors.Wrap(err, "error listening to messages")
	}

	go r.handle(msgCh)
	return nil
}

func (r *Router) handle(msgCh <-chan []byte) {
	for msg := range msgCh {
		_ = msg
		// handle msg
		resp := "OK"

		err := r.server.Send([]byte(resp))
		if err != nil {
			r.logger.WithError(err).Error("Error sending response via nanomsg")
		}
	}
}

func (r *Router) Close() error {
	if r.running == nil {
		return errors.New("router not running")
	}
	r.cancelFn()
	<-r.running
	return nil
}
