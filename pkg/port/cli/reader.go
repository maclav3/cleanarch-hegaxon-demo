package cli

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/reader"
	domain "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (r *Router) readerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reader",
		Short: "Manage the directory of readers",
	}
}

func (r *Router) addReaderCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "add",
		Short: "Add a new reader to the directory",
		RunE: func(c *cobra.Command, args []string) error {
			flags := pflag.NewFlagSet("add-reader", pflag.ContinueOnError)
			name := flags.String("name", "", "the name of the reader")

			err := flags.Parse(args)
			if err != nil {
				return errors.Wrap(err, "error parsing flags")
			}

			cmd := reader.Add{
				ID:   domain.ID(uuid.Must(uuid.NewV4()).String()),
				Name: *name,
			}
			err = r.app.Commands.AddReader.Handle(cmd)
			if err != nil {
				return errors.Wrap(err, "error calling add reader command handler")
			}

			_, err = c.OutOrStdout().Write([]byte("ID: " + cmd.ID.String()))
			if err != nil {
				return errors.Wrap(err, "error writing OK response")
			}
			return nil
		},
	}
	c.Flags().String("name", "", "the name of the reader")
	return c
}

func (r *Router) activateReaderCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "activate",
		Short: "Activate an existing reader",
		RunE: func(c *cobra.Command, args []string) error {
			flags := pflag.NewFlagSet("activate-reader", pflag.ContinueOnError)
			id := flags.String("id", "", "the id of the reader")

			err := flags.Parse(args)
			if err != nil {
				return errors.Wrap(err, "error parsing flags")
			}

			cmd := reader.Activate{
				ID: domain.ID(*id),
			}
			err = r.app.Commands.ActivateReader.Handle(cmd)
			if err != nil {
				return errors.Wrap(err, "error calling activate reader command handler")
			}

			_, err = c.OutOrStdout().Write([]byte("ID: " + cmd.ID.String()))
			if err != nil {
				return errors.Wrap(err, "error writing OK response")
			}
			return nil
		},
	}
	c.Flags().String("id", "", "the id of the reader")
	return c
}

func (r *Router) deactivateReaderCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "deactivate",
		Short: "Deactivate an existing reader",
		RunE: func(c *cobra.Command, args []string) error {
			flags := pflag.NewFlagSet("deactivate-reader", pflag.ContinueOnError)
			id := flags.String("id", "", "the id of the reader")

			err := flags.Parse(args)
			if err != nil {
				return errors.Wrap(err, "error parsing flags")
			}

			cmd := reader.Deactivate{
				ID: domain.ID(*id),
			}
			err = r.app.Commands.DeactivateReader.Handle(cmd)
			if err != nil {
				return errors.Wrap(err, "error calling deactivate reader command handler")
			}

			_, err = c.OutOrStdout().Write([]byte("ID: " + cmd.ID.String()))
			if err != nil {
				return errors.Wrap(err, "error writing OK response")
			}
			return nil
		},
	}
	c.Flags().String("id", "", "the id of the reader")
	return c
}

func (r *Router) registerReaderCommands() {
	readerCmd := r.readerCmd()
	readerCmd.AddCommand(r.addReaderCmd())
	readerCmd.AddCommand(r.activateReaderCmd())
	readerCmd.AddCommand(r.deactivateReaderCmd())

	r.rootCmd.AddCommand(readerCmd)
}
