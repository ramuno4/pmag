package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/Jon1105/pmag/utilities"
	"github.com/spf13/cobra"
)

var openCommand = &cobra.Command{
	Use:     "open",
	Short:   "Open an existing project",
	Long:    "", // TODO
	Example: "pmag open go <project name>\npmag create go",
	Args:    cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var lang, err1 = utilities.GetLanguage(args[0], Config.Languages)
		if len(args) == 1 { // pmag open flutter or pmag open Todo
			if err1 != nil {
				if Config.InferLanguage {
					var projectPath, lang, err1 = utilities.InferLanguage(args, Config)
					if err1 == nil {
						// return fmt.Errorf("failed to infer language: project %q not found in any languages", args[0])
						var editorPath, err2 = utilities.GetEditorPath(lang.EditorPath, Config.DefaultEditorPath)
						if err2 != nil {
							return err2
						}
						return utilities.Open(projectPath, editorPath)
					}
				}

				var emsg = fmt.Sprintf("invalid language name %q.", args[0])
				if !Config.InferLanguage {
					emsg += fmt.Sprintf("\nIf you wish \"pmag open\" to automatically infer the language for project %q, you may turn this feature on in config.yaml", args[0])
				}
				return errors.New(emsg)

			}

			var projects, err3 = utilities.GetProjects(lang.Path)
			if err3 != nil {
				return err3
			}
			var projectPath, err4 = utilities.PickProject(lang, projects)
			if err4 != nil {
				return err4
			}

			var editorPath, err2 = utilities.GetEditorPath(lang.EditorPath, Config.DefaultEditorPath)
			if err2 != nil {
				return err2
			}
			return utilities.Open(projectPath, editorPath)
		} else {
			if err1 != nil { // language not given
				return err1
			}

			var editorPath, err2 = utilities.GetEditorPath(lang.EditorPath, Config.DefaultEditorPath)
			if err2 != nil {
				return err2
			}

			var projects, err3 = utilities.GetProjects(lang.Path)
			if err3 != nil {
				return err3
			}

			var input string = args[1]
			if num, err5 := strconv.Atoi(input); err5 == nil {
				if num <= len(projects) {
					return utilities.Open(filepath.Join(lang.Path, projects[num-1].Name()), editorPath)
				}
				return fmt.Errorf("project index out of range: projects[%d]; Max: %d", num, len(projects))

			}
			var path string = filepath.Join(lang.Path, input)
			var exists, err6 = utilities.Exists(path)
			if err6 != nil {
				return err6
			}
			if exists {
				return utilities.Open(path, editorPath)
			} else if !exists {
				return fmt.Errorf("%q does not exist\nUse \"pmag create\" to create a project", path)
			}

		}
		return nil
	},
}

func openCmd() *cobra.Command {
	return openCommand
}
