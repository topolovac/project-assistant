package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

func getDirectoryInfo(path string, config *Config) (Directory, error) {
	content, err := os.ReadDir(path)
	if err != nil {
		return Directory{}, err
	}

	files := []File{}
	directories := []Directory{}

	for _, entry := range content {

		if shouldEntryBeIgnored(entry.Name(), config) {
			continue
		}

		if entry.IsDir() {
			directoryInfo, err := getDirectoryInfo(path+"/"+entry.Name(), config)
			if err != nil {
				return Directory{}, err
			}
			directories = append(directories, directoryInfo)
		} else {
			content, err := os.ReadFile(path + "/" + entry.Name())
			if err != nil {
				return Directory{}, err
			}

			file := File{
				Name:    entry.Name(),
				Type:    entry.Type().String(),
				Content: string(content),
			}
			files = append(files, file)
		}
	}

	directory := Directory{
		Name:        path,
		Files:       files,
		Directories: directories,
	}
	return directory, err
}

func shouldEntryBeIgnored(entry_name string, config *Config) bool {
	for _, ignoreFile := range config.IgnoreSettings.IgnoreFiles {
		if entry_name == ignoreFile {
			return true
		}
	}

	for _, ignoreDir := range config.IgnoreSettings.IgnoreDirs {
		if entry_name == ignoreDir {
			return true
		}
	}

	for _, ignoreFileWithPattern := range config.IgnoreSettings.IgnoreFilesWithPattern {
		matched, _ := regexp.MatchString(ignoreFileWithPattern, entry_name)
		if matched {
			return true
		}
	}

	for _, ignoreDir := range config.IgnoreSettings.IgnoreDirs {
		if entry_name == ignoreDir {
			return true
		}
	}

	return false
}

func createOutputDir(path string, config *Config) error {
	dir, _ := os.Stat(path + "/" + config.OutputDir)
	if dir != nil {
		return nil
	}

	err := os.Mkdir(path+"/"+config.OutputDir, 0755)
	if err != nil {
		return err
	}
	return nil
}

func createMDFile(content string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

var config_file_name = "project_assistant.config.json"

func getConfig(path string) *Config {
	default_config := &Config{
		OutputDir: CONFIG_DEFAULT_OUTPUT_DIR,
		IgnoreSettings: IgnoreSettings{
			IgnoreFiles:            []string{},
			IgnoreDirs:             []string{},
			IgnoreFilesWithPattern: []string{},
		},
	}

	content, err := os.ReadFile(path + "/" + config_file_name)
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

func logStructAsJSON(s interface{}) {
	jsonVal, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonVal))
}
