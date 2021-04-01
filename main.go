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
	var args, flags = utilities.Filter(osArgs, func(str string) bool {
		return str[:1] != "-"
	})

	var command string = args[1]

	var err2 error

	switch command {
	case "create":
		err2 = create.Create(args, flags, conf)
	case "open":
		err2 = open.Open(args, flags, conf)
	// case "vcs":
	// 	err2 = vcs.Vcs(args, flags, conf)
	case "help":
		err2 = help.Help(args, conf)
	default:
		log.Fatalf("invalid command %s\nUse `%s help` to see a list of commands", command, args[0])
	}

	if err2 != nil {
		log.Fatal(err2)
	}
}
