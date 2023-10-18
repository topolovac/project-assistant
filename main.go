package main

import (
	"encoding/json"
	"log"
)

type Application struct {
	openAI *OpenAI
	config *Config
}

func main() {
	config := getConfig()

	oai := OpenAI{
		jwt: config.OpenAIKey,
	}

	oai.Init(oai.jwt)

	app := Application{
		openAI: &oai,
		config: config,
	}

	logStructAsJSON(config)

	dir_info, err := getDirectoryInfo(config.RootPath, config)
	if err != nil {
		log.Println("Error getting directory info: ", err)
		panic(err)
	}

	logStructAsJSON(dir_info)

	// save_path := "./" + config.RootPath + "/" + config.OutputDir + "/documentation.md"

	// app.CreateDocumentation(dir_info, save_path)

	err = app.UpdateCodebase(dir_info)

	if err != nil {
		log.Println("Error updating codebase: ", err)
		panic(err)
	}
}

func (app *Application) CreateDocumentation(dir_info Directory, save_path string) {
	log.Println("Creating documentation for " + app.config.RootPath + " in " + save_path)

	content, err := app.openAI.CreateDocumentationRequest(dir_info)
	if err != nil {
		log.Println("Error creating documentation: ", err)
		panic(err)
	}

	err = createOutputDir(app.config)

	if err != nil {
		log.Println("Error creating output directory: ", err)
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
