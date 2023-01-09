package vcs

import "github.com/Jon1105/pmag/utilities"

func Git(path string) error {
	return utilities.RunCommand("", "git", "init", path)
}
