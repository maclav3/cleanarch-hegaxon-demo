package cli

import (
	"context"
	"fmt"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	zmq "github.com/zeromq/gomq"
	"github.com/zeromq/gomq/zmtp"
)

type Router struct {
	cancelFn func()
	running  chan struct{}
	rootCmd  *cobra.Command
	logger   log.Logger
	address  string
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
// It blocks while the router is running.
func (r *Router) Run(ctx context.Context) {
	fmt.Println("Enter command or type 'help' to show available commands.")
	ctx, r.cancelFn = context.WithCancel(ctx)

	r.running = make(chan struct{})
	defer close(r.running)
	commandsCh := make(chan string)
	responseCh := make(chan string)
	go func() {
		err := r.listen(ctx, commandsCh, responseCh)
		if err != nil {

		}
	}()
	go func() {
		<-ctx.Done()
		close(commandsCh)
		close(responseCh)
	}()

	for cmd := range commandsCh {
		r.logger.WithField("cmd", cmd).Debug("command received")

		// handle cmd
		resp := "OK"
		responseCh <- resp
	}
}

func (r Router) listen(ctx context.Context, commandsCh, responseCh chan string) error {
	server := zmq.NewServer(zmtp.NewSecurityNull())
	_, err := server.Bind(r.address)
	if err != nil {
		return errors.Wrap(err, "could not bind zmq server")
	}
	go func() {
		<-ctx.Done()
		server.Close()
	}()

	if err != nil {
		return errors.Wrap(err, "could not bind socket to address")
	}

	for cmd := range server.RecvChannel() {
		if cmd.MessageType != zmtp.UserMessage {
			continue
		}
		r.logger.WithField("msg", fmt.Sprintf("%+v", cmd)).Debug("command received")
		commandsCh <- string(cmd.Body[0])
		resp := <-responseCh
		err = server.Send([]byte(resp))
		if err != nil {
			r.logger.WithError(err).Error("error responding over zmq socket")
		}
	}
	return nil
}

func (r *Router) Shutdown() error {
	if r.running == nil {
		return errors.New("router not running")
	}
	r.cancelFn()
	<-r.running
	return nil
}
