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

	scanner := bufio.NewScanner(os.Stdin)
	r.running = make(chan struct{})
	defer close(r.running)
	for {
		scanner.Scan()
		if scanner.Err() != nil {
			// nevermind any errors
			continue
		}
		text := scanner.Text()
		if text == "" {
			continue
		}

		// handle cmd
		fmt.Println("CMD: " + text)

		select {
		case <-ctx.Done():
			fmt.Println("Interrupted CLI router")
			return
		default:
		}
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
