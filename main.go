package main

import (
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

	config := getConfig(*directory_name)

	err := createOutputDir(*directory_name, config)

	if err != nil {
		panic(err)
	}

	logStructAsJSON(config)

	dir_info, err := getDirectoryInfo(*directory_name, config)
	if err != nil {
		panic(err)
	}

	logStructAsJSON(dir_info)

	oai := OpenAI{
		jwt: *jwt,
	}

	oai.Init(oai.jwt)

	save_path := "./" + *directory_name + "/" + config.OutputDir + "/documentation.md"

	fmt.Println("Creating documentation for " + *directory_name + " in " + save_path)

	content, err := oai.CreateDocumentationRequest(dir_info)
	if err != nil {
		panic(err)
	}

	err = createMDFile(content, save_path)

	if err != nil {
		panic(err)
	}

	fmt.Println("Completed creating documentation for " + *directory_name + " in " + save_path)
}
