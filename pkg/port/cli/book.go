package cli

import (
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/book"
	domain "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
)

func (r *Router) bookCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "book",
		Short: "Manage the book directory",
	}
}

func (r *Router) addBookCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "add",
		Short: "Add a new book to the directory",
		RunE: func(c *cobra.Command, args []string) error {
			flags := pflag.NewFlagSet("add-book", pflag.ContinueOnError)
			author := flags.String("author", "", "the author of the book")
			title := flags.String("title", "", "the title of the book")

			err := flags.Parse(args)
			if err != nil {
				return errors.Wrap(err, "error parsing flags")
			}

			cmd := book.Add{
				ID:     domain.ID(uuid.Must(uuid.NewV4()).String()),
				Author: *author,
				Title:  *title,
			}
			err = r.app.Commands.AddBook.Handle(cmd)
			if err != nil {
				return errors.Wrap(err, "error calling add book command handler")
			}

			_, err = c.OutOrStdout().Write([]byte("ID: " + cmd.ID.String()))
			if err != nil {
				return errors.Wrap(err, "error writing OK response")
			}
			return nil
		},
	}
	// this is duplicated so that the flags may be known without running the command
	// for generalized treatment of the `-h` flag.
	// todo: this could be done in a prettier way; or go-generated
	c.Flags().String("author", "", "the author of the book")
	c.Flags().String("title", "", "the title of the book")

	return c
}

func (r *Router) registerBookCommands() {
	// IDEA: it would be a good idea to `go generate` the commands and their args from the app layer
	// It might be tricky to parse the args into correct types, but should be doable
	bookCmd := r.bookCmd()
	bookCmd.AddCommand(r.addBookCmd())

	r.rootCmd.AddCommand(bookCmd)
}
