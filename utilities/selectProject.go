package utilities

import (
	"io/fs"
	"io/ioutil"
	"bufio"
	"os"
	"fmt"
	"strconv"
	"path/filepath"

	"github.com/Jon1105/pmag/conf"
)

func GetProjects(path string) ([]fs.FileInfo, error) {
	var files, err = ioutil.ReadDir(path)
	if err != nil {
		return []fs.FileInfo{}, err
	}
	var projects []fs.FileInfo
	for _, file := range files {
		if file.IsDir() {
			projects = append(projects, file)
		}
	}
	return projects, nil
}

// returns shouldCancel, projectPath, error
func PickProject(language conf.Language, projects []fs.FileInfo) (string, error) {
	// Print all Projects
	for index, file := range projects {
		fmt.Printf("%d: %s\n", index+1, file.Name())
	}

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pick a project: ")
		scanner.Scan()
		var input string = scanner.Text()
		if num, err := strconv.Atoi(input); err == nil && num <= len(projects) {
			return filepath.Join(language.Path, projects[num-1].Name()), nil
		} else if input == "" {
			continue
		} else { // input == "." || input == "doit"
			var path string = filepath.Join(language.Path, input)
			var exists, err4 = Exists(path)
			if err4 != nil {
				return "", err4
			}
			if exists {
				return path, nil
			} else if !exists {
				fmt.Println("Invalid Entry")
				continue
			}
		}
	}
}
