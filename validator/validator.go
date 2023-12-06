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

		// fmt.Printf("Unmarshalled %s successfully\n", filename)

		// fmt.Println("recipe.Name:", recipe.Name)
		// fmt.Println("recipe.DateAdded:", recipe.DateAdded)
		// fmt.Println("recipe.Source:", recipe.Source)
		// fmt.Println("recipe.Source.Name:", recipe.Source.Name)
		// fmt.Println("recipe.Source.Modifications:", recipe.Source.Modifications)
		// fmt.Println("recipe.Source.Submitter:", recipe.Source.Submitter)
		// fmt.Println("recipe.Source.Url:", recipe.Source.Url)
		// fmt.Println("recipe.Source:", recipe.Source)
		// fmt.Println("recipe.Type:", recipe.Type)
		// fmt.Println("recipe.Descriptors:", recipe.Descriptors)
		// fmt.Println("recipe.Cuisine:", recipe.Cuisine)
		// fmt.Println("recipe.Ingredients:", recipe.Ingredients)
		// fmt.Println("recipe.Instructions:", recipe.Instructions)
		// fmt.Println("recipe.Usage:", recipe.Usage)
		// fmt.Println("recipe.Notes:", recipe.Notes)
		// fmt.Println("recipe.Next:", recipe.Next)

		err = validate.Struct(recipe)
		if err != nil {
			errCount += 1
			fmt.Printf("  error validating struct for %s: %v\n", filename, err)
			fmt.Printf("  %s\n", err.Error())
			err = nil
			continue
		}

		passCount += 1
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
