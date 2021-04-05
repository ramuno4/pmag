package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jon1105/pmag/utilities"
	"github.com/Jon1105/pmag/vcs"
	"github.com/spf13/cobra"

	"github.com/otiai10/copy"
)

var (
	readme       bool
	vcsState     bool
	ghVisibility bool
)

var create = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "create a new project",
	Long:    "", // TODO
	Example: "pmag create go <project name>",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var lang, err1 = utilities.GetLanguage(args[0], Config.Languages)
		if err1 != nil {
			return err1
		}
		var editorPath, err2 = utilities.GetEditorPath(lang.EditorPath, Config.DefaultEditorPath)
		if err2 != nil {
			return err2
		}

		var projectName = args[1]
		var projectPath = filepath.Join(lang.Path, projectName)

		if lang.TemplatePath != "" { // Copy Template folder
			var templatePath, err3 = filepath.Abs(lang.TemplatePath)
			if err3 != nil {
				return err3
			}
			var err4 = copy.Copy(templatePath, projectPath)
			if err4 != nil {
				return err4
			}
		} else { // Create folder
			os.Mkdir(projectPath, 0755)
		}

		if readme { // Create Readme
			var contents string = fmt.Sprintf("# %v\n---\n", strings.Title(projectName))
			var readmePath = filepath.Join(projectPath, "README.md")
			var err5 = ioutil.WriteFile(readmePath, []byte(contents), 0777)
			if err5 != nil {
				return err5
			}
		}

		if vcsState {
			var err6 error
			switch Config.Vcs {
			case "git":
				err6 = vcs.Git(projectPath)
			case "github":
				err6 = vcs.Github(Config.GhKey, projectPath, ghVisibility)
			}
			if err6 != nil {
				return err6
			}
		}
		if len(lang.InitialCommand) != 0 {
			var mappings = map[string]string{
				"projectName":  projectName,
				"projectPath":  projectPath,
				"languageName": lang.Name,
			}
			var command, err7 = parseCommand(lang.InitialCommand, mappings)
			if err7 != nil {
				return err7
			}

			var err8 = utilities.RunCommand(projectPath, command[0], command[1:]...)
			if err8 != nil {
				return err8
			}
		}
		return utilities.Open(projectPath, editorPath)
	},
}

func createCmd() *cobra.Command {
	// TODO fill in usage
	create.Flags().BoolVarP(&readme, "readme", "r", Config.DefaultCreateREADME, "")
	create.Flags().BoolVarP(&vcsState, "vcs", "v", Config.DefaultVcsState, "")
	create.Flags().BoolVarP(&ghVisibility, "public", "p", Config.DefaultGithubVisibility, "")

	return create
}

func parseCommand(command []string, mappings map[string]string) ([]string, error) {
	for i := range command {
		var idxStart int
		var idxEnd int
		for {
			idxStart = strings.Index(command[i], "{{")
			idxEnd = strings.Index(command[i], "}}")
			if idxStart == -1 || idxEnd == -1 {
				break
			}
			var replacement = mappings[command[i][idxStart+2:idxEnd]]
			if replacement == "" {
				return nil, fmt.Errorf("invalid variable name :\"{{%s}}\". Please update config.yaml", command[i][idxStart+2:idxEnd])
			}

			command[i] = command[i][:idxStart] + replacement + command[i][idxEnd+2:]
		}
	}
	return command, nil
}
