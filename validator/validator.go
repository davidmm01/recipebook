package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

type Recipe struct {
	Name      string   `validate:"required"`
	DateAdded string   `yaml:"date_added" validate:"required"`
	Source    struct { // Source should be an array
		Name          string // maybe keep this
		Url           string // keep this
		Modifications string // see below RE note
		// lots of recipes have Type, but i think its pointless and can be replaced anyway
		Submitter string // Submitter shoudl probably be its own thing
		// replace type and modification with generic note section?
	}
	Type         string   `validate:"required,oneof=component meal cocktail"` // TODO: make enum
	Descriptors  []string `validate:"required"`
	Cuisine      string   `validate:"required"`
	Ingredients  []string `validate:"required"`
	Instructions []string `validate:"required"`
	Usage        []string
	Notes        []string
	Next         []string
}

type Cuisine struct {
	Name  string
	Count int
}

type Descriptor struct {
	Name  string
	Count int
}

// relative path from makefile to the recipe files
const RECIPES_DIR = "recipes/"

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// fileNameRegexStr only matches strings that are lower case, snake case, and end in the `.yaml`
const fileNameRegexStr = `[a-z]+(_[a-z]+)*.yaml$`

func getValidRecipes(showOutput bool) (recipes []Recipe, cuisines []Cuisine, descriptors []Descriptor) {
	// TODO: I hate this showOutput
	files, err := os.ReadDir(RECIPES_DIR)
	if err != nil {
		log.Fatalf("error trying to read directory: %v", err)
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	passCount := 0

	fileNameRegex := regexp.MustCompile(fileNameRegexStr)

	// set cusines & descriptors up as blank maps so that we can just get the unique ones
	cuisinesMap := map[string]int{}
	descriptorsMap := map[string]int{}

	for i, f := range files {
		fileName := f.Name()
		filePath := fmt.Sprintf("%s%s", RECIPES_DIR, fileName)

		if showOutput {
			fmt.Printf("%d. %s\n", i+1, filePath)
		}

		// TODO: collect up more errors before erroring out

		if !fileNameRegex.Match([]byte(fileName)) {
			if showOutput {
				fmt.Printf("💣 filename '%s' does not conform to lower case, snake case, or missing extension '.yaml'\n", fileName)
			}
			continue
		}

		contents, err := os.ReadFile(filePath)
		if err != nil {
			if showOutput {
				fmt.Printf("💣 error trying to read file: %v", err)
			}
			continue
		}

		recipe := Recipe{}
		err = yaml.Unmarshal([]byte(contents), &recipe)
		if err != nil {
			if showOutput {
				fmt.Printf("💣 error unmarhsalling %s: %v\n", filePath, err)
			}
			err = nil
			continue
		}

		err = validate.Struct(recipe)
		if err != nil {
			if showOutput {
				fmt.Printf("💣 error validating struct for %s\n%v\n", filePath, err)
			}
			err = nil
			continue
		}

		// the name of the recipe should be the title case representation of the filename
		c := cases.Title(language.English)
		fileNameNoExt := fileName[:len(fileName)-5]
		fileNameNoExtSpaces := strings.ReplaceAll(fileNameNoExt, "_", " ")
		fileNameAsTitleCase := c.String(fileNameNoExtSpaces)
		recipeNameIgnoreApostrophes := strings.ReplaceAll(recipe.Name, "'", "")
		if recipeNameIgnoreApostrophes != fileNameAsTitleCase {
			if showOutput {
				fmt.Printf("💣recipe name '%s' should be the title case variation of file name '%s'\n", recipe.Name, fileName)
				fmt.Println("        recipe.Name:", recipe.Name)
				fmt.Println("fileNameAsTitleCase:", fileNameAsTitleCase)
			}
			err = nil
			continue
		}

		passCount += 1
		recipes = append(recipes, recipe)

		_, ok := cuisinesMap[recipe.Cuisine]
		if ok {
			cuisinesMap[recipe.Cuisine]++
		} else {
			cuisinesMap[recipe.Cuisine] = 1
		}
		for _, descriptor := range recipe.Descriptors {
			_, ok := descriptorsMap[descriptor]
			if ok {
				descriptorsMap[descriptor]++
			} else {
				descriptorsMap[descriptor] = 1
			}
		}

		if showOutput {
			fmt.Println("💚 PASS")
		}
	}

	for key, value := range cuisinesMap {
		cuisines = append(cuisines, Cuisine{Name: key, Count: value})
		// fmt.Printf("name=%s, value=%d\n", key, value)
	}
	for key, value := range descriptorsMap {
		descriptors = append(descriptors, Descriptor{Name: key, Count: value})
		// fmt.Printf("name=%s, value=%d\n", key, value)
	}

	// fmt.Println("cuisines:", cuisines)
	// fmt.Println("descriptors:", descriptors)

	if showOutput {
		errCount := len(files) - passCount

		var emoji string
		if errCount == 0 {
			emoji = "😎😎😎😎"
		} else {
			emoji = "💣💣💣💣"
		}
		fmt.Printf("\n%s FAIL = %d, PASS = %d %s\n", emoji, errCount, passCount, emoji)

		if showOutput {
			os.Exit(errCount)
		}
	}

	return recipes, cuisines, descriptors
}
