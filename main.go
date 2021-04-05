package main

import (
	_ "embed"
	"log"

	"github.com/Jon1105/pmag/cmd"
	"github.com/Jon1105/pmag/conf"
)

//go:embed config.yaml
var configBytes []byte

var config conf.Config

func main() {
	var err error
	config, err = conf.GetConfig(configBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	cmd.Config = &config

	cmd.Execute()
}
