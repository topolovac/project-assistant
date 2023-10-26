package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var config_file_name = "project_assistant.config.json"

func getConfig() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(config_file_name); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("Config file does not exist. Create project_assistant.config.json in the root of your project.")
	}

	content, err := os.ReadFile("./" + config_file_name)
	if err != nil {
		return nil, err
	}

	conf := &Config{}

	err = json.Unmarshal(content, conf)
	if err != nil {
		fmt.Println("Error unmarshalling config file: ", err)
		return nil, errors.New("Could not parse config file. Please make sure it is valid JSON.")
	}

	// append default ignore files
	files := []string{
		".DS_Store",
		".gitingore",
	}

	conf.IgnoreSettings.IgnoreFiles = append(conf.IgnoreSettings.IgnoreFiles, files...)
	conf.RootPath = wd + "/" + conf.RootPath
	return conf, err
}
