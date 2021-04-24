package cmd

import (
	"fmt"
	"os"

	"github.com/Jon1105/pmag/vcs"
	"github.com/spf13/cobra"
)

var vcsCommand = &cobra.Command{
	Use:   "vcs",
	Short: "Initialize a version control system in your current directory",
	Long:  "", // TODO,
	Example: "pmag vcs --github -p",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		var dir, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("could not retrieve the current project path")
		}
		if githubFlag {
			return vcs.Github(Config.GhKey, dir, !ghVisibilityFlag)
		} else if gitFlag {
			return vcs.Git(dir)
		} else {
			return fmt.Errorf("vcs used but no version control system set in config.yaml or through flags. please update config.yaml or use the --git or --github flags")
		}
	},
}

func vcsCmd() *cobra.Command {
	vcsCommand.PersistentFlags().BoolVarP(&gitFlag, "git", "g", Config.Vcs == "git", "Use git as the version control system")
	vcsCommand.PersistentFlags().BoolVar(&githubFlag, "github", Config.Vcs == "github", "Use github as the version control system")
	vcsCommand.PersistentFlags().BoolVarP(&ghVisibilityFlag, "public", "p", Config.DefaultGithubVisibility, "If github is used, make the created repository public")
	return vcsCommand
}
