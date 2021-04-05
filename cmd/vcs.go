package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var vcsCommand = &cobra.Command{
	Use:   "vcs",
	Short: "Initialize a version control system in your project",
	Long:  "", // TODO,
	Args: cobra.ArbitraryArgs, // TODO
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vcs", args)
	},
}

func vcsCmd() *cobra.Command {
	return vcsCommand
}
