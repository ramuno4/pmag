package main

import (
	_ "embed"
	"log"

	"os"

	"github.com/Jon1105/pmag/conf"
	"github.com/Jon1105/pmag/create"
	"github.com/Jon1105/pmag/help"
	"github.com/Jon1105/pmag/open"
	"github.com/Jon1105/pmag/utilities"
	// "github.com/Jon1105/pmag/vcs"
)

//go:embed config.yaml
var configYaml []byte

func main() {
	var conf, err = conf.GetConfig(configYaml)
	if err != nil {
		log.Fatal(err)
	}

	var osArgs []string = os.Args

	if len(osArgs) == 1 {
		osArgs = append(osArgs, "help")
	}

	// filter out all flags
	var flags []string
	for i, v := range osArgs {
		if v[:1] == "-" {
			osArgs = utilities.Remove(osArgs, i)
			flags = append(flags, v)
		}
	}

	var command string = osArgs[1]

	var err2 error

	switch command {
	case "create":
		err2 = create.Create(osArgs, flags, conf)
	case "open":
		err2 = open.Open(osArgs, flags, conf)
	// case "vcs":
	// 	err2 = vcs.Vcs(osArgs, flags, conf)
	case "help":
		err2 = help.Help(osArgs, conf)
	default:
		log.Fatalf("invalid command %s\nUse `%s help` to see a list of commands", command, osArgs[0])
	}

	if err2 != nil {
		log.Fatal(err2)
	}
}
