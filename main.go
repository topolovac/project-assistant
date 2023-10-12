package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

func main() {

	// get jwt and directory name from flags
	jwt := flag.String("jwt", "", "OpenAI API key")
	directory_name := flag.String("directory", "", "Directory to create documentation for")

	flag.Parse()

	if *jwt == "" || *directory_name == "" {
		panic("Please provide a valid jwt and directory name")
	}

	dir, err := getDirectoryInfo(*directory_name)
	if err != nil {
		panic(err)
	}

	jsonVal, err := json.MarshalIndent(dir, "", "\t")
	if err != nil {
		panic(err)
	}

	oai := OpenAI{
		jwt: *jwt,
	}

	oai.Init(oai.jwt)

	content, err := oai.CreateDocumentationRequest(string(jsonVal))
	if err != nil {
		panic(err)
	}

	createMDFile(content, "./documentation_example_results/"+*directory_name+".md")

	fmt.Println("Completed creating documentation for " + *directory_name + " in ./documentation_example_results/" + *directory_name + ".md")
}
