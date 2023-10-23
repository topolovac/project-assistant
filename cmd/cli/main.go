package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// config := getConfig()

	// get OPEN_API_KEY from env
	jwt := os.Getenv("OPEN_API_KEY")
	log.Println("OPEN_API_KEY: ", jwt)

	// oai := OpenAI{
	// 	jwt: ,
	// }

	// oai.Init(oai.jwt)

	// app := Application{
	// 	openAI: &oai,
	// 	config: config,
	// }

	// logStructAsJSON(config)

	// err := createOutputDir(app.config)

	// if err != nil {
	// 	log.Println("Error creating output directory: ", err)
	// 	panic(err)
	// }

	// app.createProjectSummary(dir_info)

	app := &cli.App{
		Name:  "pa",
		Usage: "Project Assistant for your codebase",
		Commands: []*cli.Command{
			Info,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
