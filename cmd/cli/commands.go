package main

import (
	"errors"

	"github.com/urfave/cli/v2"
)

func (app *Application) initialise_documentation(ctx *cli.Context) error {
	config, err := getConfig()
	if err != nil {
		return errors.New("Error getting config: " + err.Error())
	}

	err = createOutputDir(config)
	if err != nil {
		return errors.New("Error creating output directory: " + err.Error())
	}

	dir_info, err := getDirectoryInfo(config.RootPath, &config.IgnoreSettings)
	if err != nil {
		return errors.New("Error getting directory info: " + err.Error())
	}

	err = app.CreateProjectSummary(dir_info)
	if err != nil {
		return errors.New("Error creating project summary: " + err.Error())
	}

	return nil
}
