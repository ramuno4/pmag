package help

import (
	"embed"
	"fmt"
)

//go:embed helpFiles/*
var helpmsg embed.FS

// TODO Find a way to reflect default values in config inside help messages for create
func Help(osArgs []string) error {

	var nArgs = len(osArgs)
	var command string
	if nArgs == 2 { // pmag help
		command = "pmag"
	} else if nArgs == 3 { // pmag help vcs
		command = osArgs[2]
	} else {
		return fmt.Errorf("invalid number or arugments provided. expected %d or %d, got %d", 2, 3, nArgs)
	}
	var helpText, err1 = readHelpFile(command, osArgs)
	if err1 != nil {
		return err1
	}
	var _, err3 = fmt.Println(helpText)
	return err3
}

func readHelpFile(command string, osArgs []string) (string, error) {
	var data, err = helpmsg.ReadFile("helpFiles/" + command + ".txt")
	if err != nil {
		return "", fmt.Errorf("invalid command %q\nUse `%s help` to see a list of commands", command, osArgs[0])
	}
	return string(data), nil
}
