package open

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/Jon1105/pmag/conf"
	"github.com/Jon1105/pmag/help"
	"github.com/Jon1105/pmag/utilities"
)

func Open(osArgs []string, flags []string, config conf.Config) error {
	var nArgs int = len(osArgs)
	if nArgs == 2 { // pmag open -> help.Help([pmag help open])
		return help.Help([]string{osArgs[0], "help", osArgs[1]})
	}

	if nArgs == 3 { // pmag open flutter
		var lang, err1 = utilities.GetLanguage(osArgs[2], config.Languages)
		if err1 != nil {
			if config.InferLanguage {
				var projectPath, lang, err1 = utilities.InferLanguage(osArgs, config)
				if err1 != nil {
					return fmt.Errorf("failed to infer language: project %q not found in any languages", osArgs[2])
				}

				var editorPath, err2 = utilities.GetEditorPath(lang, config)
				if err2 != nil {
					return err2
				}
				return utilities.Open(projectPath, editorPath)
			} else {
				return fmt.Errorf("invalid language name %q.\nIf you wish %q to automatically infer the language for project %q, you may turn this feature on in config.yaml", osArgs[2], osArgs[0], osArgs[2])
			}
		}

		var editorPath, err2 = utilities.GetEditorPath(lang, config)
		if err2 != nil {
			return err2
		}

		var projects, err3 = utilities.GetProjects(lang.Path)
		if err3 != nil {
			return err3
		}
		var projectPath, err4 = utilities.PickProject(lang, projects)
		if err4 != nil {
			return err4
		}
		return utilities.Open(projectPath, editorPath)

	} else if nArgs > 3 { // pmag open flutter doit [extra arguments ignored]
		var lang, err1 = utilities.GetLanguage(osArgs[2], config.Languages)
		if err1 != nil { // language not given
			return err1
		}

		var editorPath, err2 = utilities.GetEditorPath(lang, config)
		if err2 != nil {
			return err2
		}

		var projects, err3 = utilities.GetProjects(lang.Path)
		if err3 != nil {
			return err3
		}

		var input string = osArgs[3]
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
			return fmt.Errorf("%q does not exist\nUse `%s create` to create a project", path, osArgs[0])
		}
	}
	return nil // all possible value of nArgs accounted for previously

}
