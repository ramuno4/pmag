package vcs

import (
	"fmt"

	"github.com/Jon1105/pmag/conf"
	"github.com/Jon1105/pmag/help"
)

// TODO implement Vcs command
func Vcs(osArgs []string, flags []string, config conf.Config) error {
	var nArgs int = len(osArgs)
	if nArgs == 2 { // pmag open -> help.Help([pmag help open])
		return help.Help([]string{osArgs[0], "help", osArgs[1]})
	}
	switch config.Vcs {
	case "github":
		
	case "git":
	
	default:
		return fmt.Errorf("no version control system set. please update config.yaml")
	}
	return nil
}

