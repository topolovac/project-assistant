package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dir, err := getDirectoryInfo("./example_project")
	if err != nil {
		panic(err)
	}

	jsonVal, err := json.MarshalIndent(dir, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonVal))
}
