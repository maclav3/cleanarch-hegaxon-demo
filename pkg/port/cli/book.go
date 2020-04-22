package cli

import (
	"strconv"
	"strings"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"

	"github.com/gofrs/uuid"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/book"
	app "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/query/book"
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

func (r *Router) listBooksCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "List existing books in the directory",
		RunE: func(c *cobra.Command, args []string) error {
			flags := pflag.NewFlagSet("add-book", pflag.ContinueOnError)
			loaned := flags.String("loaned", "y/n or empty", "'y' for loaned only, 'n' for not loaned only; all otherwise")
			*loaned = strings.ToLower(strings.TrimSpace(*loaned))

			err := flags.Parse(args)
			if err != nil {
				return errors.Wrap(err, "error parsing flags")
			}

			var loanedParam *bool
			if *loaned == "y" {
				loanedParam = new(bool)
				*loanedParam = true
			} else if *loaned == "n" {
				loanedParam = new(bool)
				*loanedParam = false
			}

			books, err := r.app.Queries.ListBooks.Query(app.ListQuery{
				Loaned: loanedParam,
			})
			if err != nil {
				return errors.Wrap(err, "error querying for books")
			}

			table := tablewriter.NewWriter(c.OutOrStdout())
			table.SetHeader([]string{"#", "ID", "Title", "Author", "Loaned"})
			for i, book := range books {
				table.Append([]string{
					strconv.Itoa(i),
					book.ID().String(),
					book.Title(),
					book.Author(),
					boolToString(book.Loaned()),
				})
			}

			table.Render()
			return nil
		},
	}

	c.Flags().String("loaned", "y/n or empty", "'y' for loaned only, 'n' for not loaned only; all otherwise")
	return c
}

func (r *Router) loanBookCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "loan",
		Short: "Loan a new book for a reader",
		RunE: func(c *cobra.Command, args []string) error {
			flags := pflag.NewFlagSet("loan-book", pflag.ContinueOnError)
			bookID := flags.String("book_id", "", "the ID of the book")
			readerID := flags.String("reader_id", "", "the ID of the reader loaning the book")

			err := flags.Parse(args)
			if err != nil {
				return errors.Wrap(err, "error parsing flags")
			}

			cmd := book.Loan{
				BookID:   domain.ID(*bookID),
				ReaderID: reader.ID(*readerID),
			}
			err = r.app.Commands.LoanBook.Handle(cmd)
			if err != nil {
				return errors.Wrap(err, "error calling add book command handler")
			}

			_, err = c.OutOrStdout().Write([]byte("Loaned OK"))
			if err != nil {
				return errors.Wrap(err, "error writing OK response")
			}
			return nil
		},
	}
	// this is duplicated so that the flags may be known without running the command
	// for generalized treatment of the `-h` flag.
	// todo: this could be done in a prettier way; or go-generated
	c.Flags().String("book_id", "", "the ID of the book")
	c.Flags().String("reader_id", "", "the ID of the reader loaning the book")

	return c
}

func boolToString(b bool) string {
	if b {
		return "YES"
	}
	return "NO"
}

func (r *Router) registerBookCommands() {
	// IDEA: it would be a good idea to `go generate` the commands and their args from the app layer
	// It might be tricky to parse the args into correct types, but should be doable
	bookCmd := r.bookCmd()
	bookCmd.AddCommand(r.listBooksCmd())
	bookCmd.AddCommand(r.addBookCmd())
	bookCmd.AddCommand(r.loanBookCmd())

	r.rootCmd.AddCommand(bookCmd)
}
