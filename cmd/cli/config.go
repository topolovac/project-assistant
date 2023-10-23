package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var config_file_name = "project_assistant.config.json"

func getConfig() *Config {
	default_config := &Config{
		RootPath:  ".",
		OutputDir: CONFIG_DEFAULT_OUTPUT_DIR,
		IgnoreSettings: IgnoreSettings{
			IgnoreFiles: []string{
				".DS_Store",
				".gitingore",
			},
			IgnoreDirs: []string{
				"node_modules",
			},
			IgnoreFilesWithPattern: []string{},
		},
	}

	if _, err := os.Stat(config_file_name); errors.Is(err, os.ErrNotExist) {
		return default_config
	}

	content, err := os.ReadFile("./" + config_file_name)
	if err != nil {
		fmt.Println("Error reading config file: ", err)
		return default_config
	}

	err = json.Unmarshal(content, default_config)
	if err != nil {
		fmt.Println("Error unmarshalling config file: ", err)
		return default_config
	}

	// add config file to ignore files
	default_config.IgnoreSettings.IgnoreFiles = append(default_config.IgnoreSettings.IgnoreFiles, config_file_name)
	return default_config
}
