package main

import (
	"flag"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	config, err := getConfig()
	if err != nil {
		log.Println("Error getting config: ", err)
		panic(err)
	}

	// get open-api-key flag
	var jwt string
	flag.StringVar(&jwt, "oaikey", "", "OpenAI API Key")

	log.Printf("OPEN_API_KEY value: ", jwt)

	oai := OpenAI{
		jwt: jwt,
	}

	oai.Init(oai.jwt)

	app := Application{
		openAI: &oai,
		config: config,
	}

	cli := &cli.App{
		Name:  "pa",
		Usage: "Project Assistant for your codebase",
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Initialize project assistant. Creates basic documentation for your project and file summary.",
				Action: app.initialise_documentation,
			},
		},
	}

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
