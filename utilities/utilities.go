package utilities

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Jon1105/pmag/conf"
)

// String slice functions

func Remove(slice []string, i int) []string {
	fmt.Println(slice)
	return append(slice[:i], slice[i+1:]...)
}

func Contains(slice []string, obj string) bool {
	for _, val := range slice {
		if val == obj {
			return true
		}
	}
	return false
}

func ContainsAny(slice []string, slice2 []string) bool {
	for _, val := range slice {
		if Contains(slice2, val) {
			return true
		}
	}
	return false
}

func GetLanguage(str string, languages []conf.Language) (conf.Language, error) {
	var langString string = strings.ToLower(str)

	for _, v := range languages {
		if Contains(v.Acros, langString) {
			return v, nil
		}
	}

	return conf.Language{}, fmt.Errorf("invalid language name %q", str)
}

func InferLanguage(osArgs []string, config conf.Config) (string, conf.Language, error) {
	// osArgs = [pmag.exe open doit]
	if len(osArgs) < 3 {
		return "", conf.Language{}, fmt.Errorf("not enough arguments") // normally should not occur
	}
	var name = osArgs[2]

	for _, lang := range config.Languages {
		var projects, err = GetProjects(lang.Path)
		if err != nil {
			continue
		}
		for _, v := range projects {
			if v.Name() == name {
				return filepath.Join(lang.Path, v.Name()), lang, nil
			}
		}
	}

	return "", conf.Language{}, fmt.Errorf("could not infer project")
}

func Open(projectPath, editorPath string) error {

	return RunCommand("", editorPath, projectPath)
}

func GetEditorPath(language conf.Language, config conf.Config) (string, error) {
	if language.EditorPath != "" {
		return language.EditorPath, nil
	} else if config.DefaultEditorPath != "" {
		return config.DefaultEditorPath, nil
	} else {
		return "", fmt.Errorf("DefaultEditorPath must not be empty")
	}
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func RunCommand(dir, name string, arg ...string) error {
	var cmd = exec.Command(name, arg...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func Filter(slice []string, filter func(string) bool) ([]string, []string) {
	var valid = []string{}
	var invalid = []string{}

	for _, v := range slice {
		if filter(v) {
			valid = append(valid, v)
		} else {
			invalid = append(invalid, v)
		}
	}

	return valid, invalid
}
