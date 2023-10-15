package main

import (
	"fmt"
)

func main() {
	config := getConfig()

	logStructAsJSON(config)

	dir_info, err := getDirectoryInfo(config.RootPath, config)
	if err != nil {
		panic(err)
	}

	logStructAsJSON(dir_info)

	oai := OpenAI{
		jwt: config.OpenAIKey,
	}

	oai.Init(oai.jwt)

	save_path := "./" + config.RootPath + "/" + config.OutputDir + "/documentation.md"

	fmt.Println("Creating documentation for " + config.RootPath + " in " + save_path)

	content, err := oai.CreateDocumentationRequest(dir_info)
	if err != nil {
		panic(err)
	}

	err = createOutputDir(config)

	if err != nil {
		panic(err)
	}

	err = createMDFile(content, save_path)

	if err != nil {
		panic(err)
	}

	fmt.Println("Completed creating documentation for " + config.RootPath + " in " + save_path)
}
