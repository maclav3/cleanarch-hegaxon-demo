package cli

import (
	"bytes"
	"context"
	"strings"

	"github.com/spf13/pflag"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/app"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/nanomsg"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Router struct {
	cancelFn func()
	running  chan struct{}

	rootCmd *cobra.Command

	app    *app.Application
	logger log.Logger

	address string
	server  *nanomsg.Server
}

func NewRouter(logger log.Logger, app *app.Application, address string) *Router {
	rootCmd := &cobra.Command{
		Use:   "cleanarch-demo",
		Short: "Execute application commands/queries from CLI",
		Long: "The cleanarch-demo CLI port is called from the command line." +
			"It parses arguments from the command line, " +
			"executes application commands/queries and marshals the output to stdout.",
	}
	r := &Router{
		rootCmd: rootCmd,
		app:     app,
		logger:  logger,
		address: address,
	}
	r.registerBookCommands()

	return r
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
		args := strings.Split(string(msg), " ")

		cmd, args, err := r.rootCmd.Find(args)
		if err != nil {
			resp := errors.Wrap(err, "error parsing command").Error()
			err = r.server.Send([]byte(resp))
			if err != nil {
				r.logger.WithError(err).Error("Error sending response via nanomsg")
			}
		}

		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)

		if cmd.RunE == nil {
			err = cmd.Help()
		} else {
			err = cmd.RunE(cmd, args)
		}

		if errors.Cause(err) == pflag.ErrHelp {
			_ = cmd.Help()
		} else if err != nil {
			r.logger.WithError(err).Error("Error executing command")
			buf.Reset()
			resp := errors.Wrap(err, "error executing command").Error()
			_, _ = buf.Write([]byte(resp))
		}

		err = r.server.Send(buf.Bytes())
		if err != nil {
			r.logger.WithError(err).Error("Error sending response via nanomsg")
		}
	}

	r.logger.Info("Message handling stopped")
}

func (r *Router) Close() error {
	if r.running == nil {
		return errors.New("router not running")
	}
	r.cancelFn()
	<-r.running
	return nil
}
