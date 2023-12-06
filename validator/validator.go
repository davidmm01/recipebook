package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type Recipe struct {
	Name      string   `validate:"required"`
	DateAdded string   `yaml:"date_added" validate:"required"`
	Source    struct { // revisit Source, consider turning into Author blob instead?
		Name          string
		Url           string
		Modifications string
		Submitter     string
	}
	Type         string   `validate:"required,oneof=component meal"` // TODO: make enum
	Descriptors  []string `validate:"required"`
	Cuisine      string   `validate:"required"`
	Ingredients  []string `validate:"required"`
	Instructions []string `validate:"required"`
	Usage        []string
	Notes        []string
	Next         []string
}

const RECIPES_DIR = "recipes/"

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	files, err := os.ReadDir(RECIPES_DIR)
	if err != nil {
		log.Fatalf("error trying to read directory: %v", err)
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	errCount := 0
	passCount := 0

	for i, f := range files {
		filename := fmt.Sprintf("%s%s", RECIPES_DIR, f.Name())
		fmt.Printf("%d. %s\n", i+1, filename)

		contents, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("💣error trying to read file: %v", err)
		}

		recipe := Recipe{}
		err = yaml.Unmarshal([]byte(contents), &recipe)
		if err != nil {
			errCount += 1
			fmt.Printf("💣error unmarhsalling %s: %v\n", filename, err)
			err = nil
			continue
		}

		err = validate.Struct(recipe)
		if err != nil {
			errCount += 1
			fmt.Printf("💣error validating struct for %s\n%v\n", filename, err)
			err = nil
			continue
		}

		passCount += 1
		fmt.Println("💚 PASS")
	}

	var emoji string
	if errCount == 0 {
		emoji = "😎😎😎😎"
	} else {
		emoji = "💣💣💣💣"
	}
	fmt.Printf("\n%s FAIL = %d, PASS = %d %s\n", emoji, errCount, passCount, emoji)
	os.Exit(errCount)
}
