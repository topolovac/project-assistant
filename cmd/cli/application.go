package main

import (
	"encoding/json"
	"log"
)

type Application struct {
	openAI *OpenAI
	config *Config
}

func (app *Application) CreateDocumentation(dir_info Directory, save_path string) {
	log.Println("Creating documentation for " + app.config.RootPath + " in " + save_path)

	content, err := app.openAI.CreateDocumentationRequest(dir_info)
	if err != nil {
		log.Println("Error creating documentation: ", err)
		panic(err)
	}

	err = createMDFile(content, save_path)

	if err != nil {
		log.Println("Error creating markdown file: ", err)
		panic(err)
	}

	log.Println("Completed creating documentation for " + app.config.RootPath + " in " + save_path)
}

func (app *Application) UpdateCodebase(dir_info Directory) error {
	log.Println("Updating codebase...")

	command := CodeUpdateCommand{
		TaskType:           "chore",
		Name:               "Add missing error logs ",
		Description:        "Add missing error logs to functions that are missing them",
		AcceptanceCriteria: "All functions that are missing error logs should have them added",
		Codebase:           dir_info,
	}

	content, err := app.openAI.CreateCodeUpdateRequest(command)

	updated_dir := &Directory{}

	if err != nil {
		log.Println("Error creating code update request: ", err)
		return err
	}

	err = json.Unmarshal([]byte(content), updated_dir)

	if err != nil {
		log.Println("Content: ", content)
		log.Println("Error unmarshalling updated directory: ", err)
		return err
	}

	log.Println("Updated directory: ")
	logStructAsJSON(updated_dir)

	err = updateDirectory(*updated_dir)

	if err != nil {
		log.Println("Error updating directory: ", err)
		return err
	}

	return nil
}

func (app *Application) createProjectSummary(dir_info Directory) error {
	log.Println("Creating file summary...")

	fd := flattenDirectory(dir_info)

	grouped_file_summary := ""

	for _, file := range fd {
		log.Println("Creating summary for file: ", file.content)
		content, err := app.openAI.CreateFileSummaryRequest(file.content)

		if err != nil {
			log.Println("Error creating file summary: ", err)
			return err
		}
		log.Println("Content: ", content)
		grouped_file_summary += "FILE_PATH:" + file.path + "\nFILE_NAME:" + file.filename + "\nCONTENT:" + content + "\n\n"
	}

	log.Println("Grouped file summary: ", grouped_file_summary)

	err := createMDFile(grouped_file_summary, app.config.RootPath+"/"+app.config.OutputDir+"/file_summary.md")

	if err != nil {
		log.Println("Error creating file summary: ", err)
		return err
	}

	content, err := app.openAI.CreateProjectSummaryRequest(grouped_file_summary)

	if err != nil {
		log.Println("Error creating project summary: ", err)
		return err
	}

	err = createMDFile(content, app.config.RootPath+"/"+app.config.OutputDir+"/project_summary.md")

	if err != nil {
		log.Println("Error creating project summary: ", err)
		return err
	}

	return nil
}
