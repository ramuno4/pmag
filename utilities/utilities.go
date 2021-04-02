package utilities

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Jon1105/pmag/conf"
)

func Remove(slice []string, i int) []string {
	fmt.Println(slice)
	return append(slice[:i], slice[i+1:]...)
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

func Contains(array []string, obj string) bool {
	for _, val := range array {
		if val == obj {
			return true
		}
	}
	return false
}

func ContainsAny(array []string, array2 []string) bool {
	for _, val := range array {
		if Contains(array2, val) {
			return true
		}
	}
	return false
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

func Copy(source, destination string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath string = strings.Replace(path, source, "", 1)
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		}
		var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
		if err1 != nil {
			return err1
		}
		return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)

	})
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

func Filter(array []string, filter func(string) bool) ([]string, []string) {
	var valid = []string{}
	var invalid = []string{}

	for _, v := range array {
		if filter(v) {
			valid = append(valid, v)
		} else {
			invalid = append(invalid, v)
		}
	}

	return valid, invalid
}
