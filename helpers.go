package main

import (
	"fmt"
	"os"
)

func getDirectoryInfo(path string) (Directory, error) {
	content, err := os.ReadDir(path)
	if err != nil {
		return Directory{}, err
	}

	files := []File{}
	directories := []Directory{}

	for _, entry := range content {
		if entry.IsDir() {
			directoryInfo, err := getDirectoryInfo(path + "/" + entry.Name())
			if err != nil {
				return Directory{}, err
			}
			directories = append(directories, directoryInfo)
		} else {
			content, err := os.ReadFile(path + "/" + entry.Name())
			if err != nil {
				return Directory{}, err
			}

			fmt.Println(entry.Type())
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
