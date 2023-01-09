package utilities

import (
	"fmt"
	"os"
	"os/exec"
)


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

