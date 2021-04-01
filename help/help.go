package help

import (
	"embed"
	"fmt"

	"github.com/Jon1105/pmag/utilities"
	"github.com/Jon1105/pmag/conf"
)

//go:embed helpFiles/*
var helpmsg embed.FS

// TODO Find a way to reflect default values in config inside help messages for create
func Help(osArgs []string, config conf.Config) error {

	var nArgs = len(osArgs)
	if nArgs == 2 { // pmag help
		var helpText, err1 = helpFile("helpFiles/pmag.txt")
		if err1 != nil {
			return fmt.Errorf("error reading help message for pmag command")
		}
		var _, err2 = fmt.Println(helpText)
		return err2
	} else if nArgs == 3 { // pmag help vcs
		var command = osArgs[2]
		if !utilities.Contains([]string{"create", "open", "vcs"}, command) {
			return fmt.Errorf("invalid command %s\nUse `%s help` to see a list of commands", command, osArgs[0])
		}
		var helpText, err1 = helpFile("helpFiles/" + command + ".txt")
		if err1 != nil {
			return fmt.Errorf("error reading help message for %s command", command)
		}
		var _, err3 = fmt.Println(helpText)
		return err3
	}
	return fmt.Errorf("invlalid number or arugment provided. expected %d or %d, got %d", 2, 3, nArgs)
}

func helpFile(path string) (string, error) {
	var data, err2 = helpmsg.ReadFile(path)
	if err2 != nil {
		return "", err2
	}
	return string(data), nil

}
