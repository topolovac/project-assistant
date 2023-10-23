package main

type Directory struct {
	Name        string      `json:"name"`
	Files       []File      `json:"files"`
	Directories []Directory `json:"directories"`
}

type File struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	IsUpdated bool   `json:"is_updated"`
	IsNew     bool   `json:"is_new"`
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

type CodeUpdateCommand struct {
	TaskType           string    `json:"task_type"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	AcceptanceCriteria string    `json:"acceptance_criteria"`
	Codebase           Directory `json:"codebase"`
}
