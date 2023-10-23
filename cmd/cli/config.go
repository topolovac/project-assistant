package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var config_file_name = "project_assistant.config.json"

func getConfig() *Config {

	flags := getFlags()

	default_config := &Config{
		RootPath:  flags.directory_name,
		OpenAIKey: flags.jwt,
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

	content, err := os.ReadFile(flags.directory_name + "/" + config_file_name)
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

func getFlags() *Flags {
	jwt := flag.String("jwt", "", "OpenAI API key")
	directory_name := flag.String("directory", "", "Directory to create documentation for")

	flag.Parse()

	if *jwt == "" || *directory_name == "" {
		panic("Please provide a valid jwt and directory name")
	}

	return &Flags{
		jwt:            *jwt,
		directory_name: *directory_name,
	}
}
