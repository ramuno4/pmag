package create

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jon1105/pmag/conf"
	"github.com/Jon1105/pmag/help"
	"github.com/Jon1105/pmag/utilities"
	"github.com/Jon1105/pmag/vcs/git"
	"github.com/Jon1105/pmag/vcs/github"
)

func Create(osArgs []string, flags []string, config conf.Config) error {
	var nArgs = len(osArgs)
	if nArgs == 2 { // pmag open -> help.Help([pmag help open])
		return help.Help([]string{osArgs[0], "help", osArgs[1]}, config)
	}
	if nArgs == 3 { // pmag create flutter
		var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Enter a project name: ")
			scanner.Scan()
			var input = scanner.Text()
			if input != "" {
				return Create(append(osArgs, input), flags, config)
			}
		}
	} else if nArgs >= 4 { // pmag create flutter new_project
		var lang, err1 = utilities.GetLanguage(osArgs[2], config.Languages)
		if err1 != nil {
			return err1
		}
		var editorPath, err2 = utilities.GetEditorPath(lang, config)
		if err2 != nil {
			return err2
		}

		// Establish Flags
		var readme = checkFlag(utilities.ContainsAny(flags, []string{"-r", "-readme"}), config.DefaultCreateREADME)
		var vcs = checkFlag(utilities.ContainsAny(flags, []string{"-v", "-vcs"}), config.DefaultVcsState)
		var public = checkFlag(utilities.Contains(flags, "-p"), config.DefaultGithubVisibility)

		var projectName = osArgs[3]
		var projectPath = filepath.Join(lang.Path, projectName)

		// I could not find a way to implement utilities.Copy without it creating the base folder, and creating it twice resulted in an error, so I had to use conditional logic
		if lang.TemplatePath != "" { // Copy Template folder
			var templatePath, err3 = filepath.Abs(lang.TemplatePath)
			if err3 != nil {
				return err3
			}
			var err4 = utilities.Copy(templatePath, projectPath)
			if err4 != nil {
				return err4
			}
		} else { // Create folder
			os.Mkdir(projectPath, 0755)
		}

		if readme {
			var err5 = createREADME(projectName, projectPath)
			if err5 != nil {
				return err5
			}
		}
		if vcs {
			var err6 error
			switch config.Vcs {
			case "git":
				err6 = git.Git(projectPath)
			case "github":
				err6 = github.Github(projectPath, public)
			}
			if err6 != nil {
				return err6
			}
		}
		if lang.InitialCommand != "" {
			var mappings = map[string]string{
				"projectName":  projectName,
				"projectPath":  projectPath,
				"languageName": lang.Name,
				"languageAcro": osArgs[2], // acronym used to indentify the language
			}

			var command = parseCommand(lang.InitialCommand, mappings)

			var err7 = utilities.RunCommand(projectPath, command)
			if err7 != nil {
				return err7
			}
		}
		return utilities.Open(projectPath, editorPath)
	}
	return nil
}

func checkFlag(flagInside, defaultValue bool) bool {
	if flagInside {
		return !defaultValue
	} else {
		return defaultValue
	}
}

func createREADME(projectName, projectPath string) error {
	var contents string = fmt.Sprintf("# %v\n---\n", strings.Title(projectName))
	var readmePath = filepath.Join(projectPath, "README.md")
	return ioutil.WriteFile(readmePath, []byte(contents), 0777)
}

// TODO find a more efficient
func parseCommand(command string, mappings map[string]string) string {
	for k, v := range mappings {
		if strings.Contains(command, k) {
			command = strings.ReplaceAll(command, "{{"+k+"}}", v)
		}
	}
	return command
}
