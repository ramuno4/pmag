package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "",                  // TODO
	Long:    "",                  // TODO,
	Args:    cobra.ArbitraryArgs, // TODO
	RunE:    func(cmd *cobra.Command, args []string) error { return nil },
}

func deleteCmd() *cobra.Command {
	
	return deleteCommand
}
