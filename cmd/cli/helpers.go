package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
)

func getDirectoryInfo(path string, ignore_settings *IgnoreSettings) (Directory, error) {
	content, err := os.ReadDir(path)
	if err != nil {
		log.Println("Error reading directory: ", err)
		return Directory{}, err
	}

	files := []File{}
	directories := []Directory{}

	for _, entry := range content {

		if shouldEntryBeIgnored(entry.Name(), ignore_settings) {
			continue
		}

		if entry.IsDir() {
			directoryInfo, err := getDirectoryInfo(path+"/"+entry.Name(), ignore_settings)
			if err != nil {
				log.Println("Error getting directory info: ", err)
				return Directory{}, err
			}
			directories = append(directories, directoryInfo)
		} else {
			content, err := os.ReadFile(path + "/" + entry.Name())
			if err != nil {
				log.Println("Error reading file: ", err)
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
	return directory, nil
}

func shouldEntryBeIgnored(entry_name string, ignore_settings *IgnoreSettings) bool {
	for _, ignoreFile := range ignore_settings.IgnoreFiles {
		if entry_name == ignoreFile {
			return true
		}
	}

	for _, ignoreDir := range ignore_settings.IgnoreDirs {
		if entry_name == ignoreDir {
			return true
		}
	}

	for _, ignoreFileWithPattern := range ignore_settings.IgnoreFilesWithPattern {
		matched, _ := regexp.MatchString(ignoreFileWithPattern, entry_name)
		if matched {
			return true
		}
	}

	return false
}

func createOutputDir(config *Config) error {
	dir, _ := os.Stat(config.RootPath + "/" + config.OutputDir)
	if dir != nil {
		log.Println("Output directory already exists")
		return nil
	}

	err := os.Mkdir(config.RootPath+"/"+config.OutputDir, 0755)
	if err != nil {
		log.Println("Error creating output directory: ", err)
		return err
	}
	log.Println("Output directory created")
	return nil
}

func createMDFile(content string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Println("Error creating markdown file: ", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Println("Error writing to markdown file: ", err)
		return err
	}
	log.Println("Markdown file created")
	return nil
}

func logStructAsJSON(s interface{}) {
	jsonVal, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Println("Error logging JSON: ", err)
		panic(err)
	}

	log.Println(string(jsonVal))
}

func updateDirectory(dir Directory) error {
	log.Println("Updating directory: ", dir.Name)
	// update files
	for _, file := range dir.Files {
		if file.IsUpdated {
			log.Println("Updating file " + file.Name + "...")
			err := updateOrCreateFile(dir.Name+"/"+file.Name, file.Content)
			if err != nil {
				log.Println("Error updating file: ", err)
				return err
			}
			log.Println("File " + file.Name + " updated")
		}
	}

	// create new files
	for _, file := range dir.Files {
		if file.IsNew {
			log.Println("Creating file " + file.Name + "...")
			err := updateOrCreateFile(dir.Name+"/"+file.Name, file.Content)
			if err != nil {
				log.Println("Error creating file: ", err)
				return err
			}
			log.Println("File " + file.Name + " created")
		}
	}

	// update directories
	for _, directory := range dir.Directories {
		err := updateDirectory(directory)
		if err != nil {
			log.Println("Error updating directory: ", err)
			return err
		}
	}

	log.Println("Directory " + dir.Name + " updated")
	return nil
}

func updateOrCreateFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Println("Error writing file: ", err)
		return err
	}
	return nil
}

type FlattenedDir struct {
	filename string
	path     string
	content  string
}

func flattenDirectory(dir Directory) []FlattenedDir {
	files := []FlattenedDir{}

	for _, file := range dir.Files {
		files = append(files, FlattenedDir{
			filename: file.Name,
			path:     dir.Name + "/" + file.Name,
			content:  file.Content,
		})
	}

	for _, directory := range dir.Directories {
		files = append(files, flattenDirectory(directory)...)
	}

	return files
}
