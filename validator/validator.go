package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Recipe struct {
	Name      string
	DateAdded string // can this be some kind of date field maybe?
	Source    struct {
		Name          string
		Url           string
		Modifications string
		Submitter     string
	}
	Type         string
	Descriptors  []string
	Cuisine      string
	Ingredients  []string
	Instructions []string
	Usage        []string
	Notes        []string
	Next         []string
}

const RECIPES_DIR = "recipes/"

func main() {
	files, err := os.ReadDir(RECIPES_DIR)
	if err != nil {
		log.Fatalf("error trying to read directory: %v", err)
	}

	for _, f := range files {
		filename := fmt.Sprintf("%s%s", RECIPES_DIR, f.Name())
		contents, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("error trying to read file: %v", err)
		}

		recipe := Recipe{}
		err = yaml.Unmarshal([]byte(contents), &recipe)
		if err != nil {
			fmt.Printf("error unmarhsalling file: %v\n", err)
		}

		// TODO: now write and check rules against `recipe`, collect some helpful stats/insights
		fmt.Printf("--- t:\n%v\n\n", recipe)
	}
}
