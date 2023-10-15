package main

import (
	"fmt"
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
		panic(err)
	}

	logStructAsJSON(dir_info)

	// save_path := "./" + config.RootPath + "/" + config.OutputDir + "/documentation.md"

	// app.CreateDocumentation(dir_info, save_path)

	err = app.UpdateCodebase(dir_info)

	if err != nil {
		panic(err)
	}
}

func (app *Application) CreateDocumentation(dir_info Directory, save_path string) {
	fmt.Println("Creating documentation for " + app.config.RootPath + " in " + save_path)

	content, err := app.openAI.CreateDocumentationRequest(dir_info)
	if err != nil {
		panic(err)
	}

	err = createOutputDir(app.config)

	if err != nil {
		panic(err)
	}

	err = createMDFile(content, save_path)

	if err != nil {
		panic(err)
	}

	fmt.Println("Completed creating documentation for " + app.config.RootPath + " in " + save_path)
}

func (app *Application) UpdateCodebase(dir_info Directory) error {
	fmt.Println("Updating codebase...")

	// command := CodeUpdateCommand{
	// 	TaskType:           "FEATURE",
	// 	Name:               "Update files function",
	// 	Description:        "Implement function that will update existing codebase",
	// 	AcceptanceCriteria: "Function accepts Directory struct and updates existing files or creates new ones",
	// 	Codebase:           dir_info,
	// }

	// content, err := app.openAI.CreateCodeUpdateRequest(command)

	// updated_dir := &Directory{}

	// if err != nil {
	// 	return err
	// }

	// err = json.Unmarshal([]byte(content), updated_dir)

	// if err != nil {
	// 	return err
	// }

	err := updateDirectory(dir_info)

	if err != nil {
		return err
	}

	return nil
}
