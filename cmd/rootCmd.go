package cmd

import (
	"log"

	"github.com/Jon1105/pmag/conf"
	"github.com/spf13/cobra"
)

var Config *conf.Config

var rootCmd = &cobra.Command{
	Use:     "pmag",
	Short:   "pmag is a project manager for multiple programming languages and frameworks",
	Example: "pmag create golang <project name> -vcs -readme",
}

var ( // Flag variables
	gitFlag          bool
	githubFlag       bool
	ghVisibilityFlag bool
)

func Execute() {
	rootCmd.AddCommand(
		createCmd(),
		openCmd(),
		vcsCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
