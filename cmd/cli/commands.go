package main

import (
	"errors"

	"github.com/urfave/cli/v2"
)

var Init = &cli.Command{
	Name:   "init",
	Usage:  "Initialize project assistant. Creates basic documentation for your project and file summary.",
	Action: initialise_documentation,
}

func initialise_documentation(ctx *cli.Context) error {
	config, err := getConfig()
	if err != nil {
		return errors.New("Error getting config: " + err.Error())
	}

	logStructAsJSON(config)

	err = createOutputDir(config)
	if err != nil {
		return errors.New("Error creating output directory: " + err.Error())
	}

	dir_info, err := getDirectoryInfo(config.RootPath, &config.IgnoreSettings)
	if err != nil {
		return errors.New("Error getting directory info: " + err.Error())
	}

	logStructAsJSON(dir_info)

	return nil
}
