package cli

import "github.com/spf13/cobra"

var bookCmd = cobra.Command{
	Use:   "book",
	Short: "Manage the book directory",
}

var addBookCmd = cobra.Command{
	Use:   "add",
	Short: "Add a new book to the directory",
}

func (r *Router) registerBookCommands() {

}
