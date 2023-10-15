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

func createOutputDir(config *Config) error {
	dir, _ := os.Stat(config.RootPath + "/" + config.OutputDir)
	if dir != nil {
		return nil
	}

	err := os.Mkdir(config.RootPath+"/"+config.OutputDir, 0755)
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

func logStructAsJSON(s interface{}) {
	jsonVal, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonVal))
}
