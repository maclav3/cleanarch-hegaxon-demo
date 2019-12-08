package cli

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Router struct {
	cancelFn func()
	running  chan struct{}
	rootCmd  *cobra.Command
}

func NewRouter() *Router {
	rootCmd := &cobra.Command{
		Use:   "cleanarch-demo",
		Short: "Execute application commands/queries from CLI",
		Long: "The cleanarch-demo CLI port is called from the command line." +
			"It parses arguments from the command line, " +
			"executes application commands/queries and marshals the output to stdout.",
	}
	return &Router{
		rootCmd: rootCmd,
	}
}

// Run starts the Router router and keeps parsing commands from the standard input.
// It blocks while the router is running.
func (r *Router) Run(ctx context.Context) {
	fmt.Println("Enter command or type 'help' to show available commands.")
	ctx, r.cancelFn = context.WithCancel(ctx)

	r.running = make(chan struct{})
	defer close(r.running)
	for cmd := range r.commands(ctx) {
		// handle cmd
		fmt.Println("CMD: " + cmd)
	}
}

func (r Router) commands(ctx context.Context) <-chan string {
	commandsCh := make(chan string)
	go r.listen(ctx, commandsCh)
	go func() {
		<-ctx.Done()
		close(commandsCh)
	}()
	return commandsCh
}

func (r Router) listen(ctx context.Context, commandsCh chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		scanner.Scan()
		if scanner.Err() != nil {
			// nevermind any errors
			continue
		}
		text := scanner.Text()
		if text == "" {
			continue
		}

		commandsCh <- text
	}
}

func (r *Router) Shutdown() error {
	if r.running == nil {
		return errors.New("router not running")
	}
	r.cancelFn()
	<-r.running
	return nil
}
