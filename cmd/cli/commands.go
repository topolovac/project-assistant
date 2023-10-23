package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var Info = &cli.Command{
	Name:   "info",
	Usage:  "Pring directory info",
	Action: info,
}

func info(ctx *cli.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	config := getConfig()

	dir_info, err := getDirectoryInfo(dir+"/"+config.RootPath, config)
	if err != nil {
		log.Println("Error getting directory info: ", err)
		panic(err)
	}

	logStructAsJSON(dir_info)

	return nil
}
