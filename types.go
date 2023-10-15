package main

type Directory struct {
	Name        string
	Files       []File
	Directories []Directory
}

type File struct {
	Name    string
	Type    string
	Content string
}

type IgnoreSettings struct {
	IgnoreFiles            []string `json:"ignore_files"`
	IgnoreDirs             []string `json:"ignore_dirs"`
	IgnoreFilesWithPattern []string `json:"ignore_files_with_pattern"`
}

var CONFIG_DEFAULT_OUTPUT_DIR = "project_assistant"

type Config struct {
	RootPath       string         `json:"root_path"`
	OpenAIKey      string         `json:"openai_key"`
	OutputDir      string         `json:"output_dir"`
	IgnoreSettings IgnoreSettings `json:"ignore_settings"`
}

type Flags struct {
	jwt            string
	directory_name string
}
