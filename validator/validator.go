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

// fileNameRegexStr only matches strings that are lower case, snake case, and end in the `.yaml`
const fileNameRegexStr = `[a-z]+(_[a-z]+)*.yaml$`

func main() {
	files, err := os.ReadDir(RECIPES_DIR)
	if err != nil {
		log.Fatalf("error trying to read directory: %v", err)
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	passCount := 0

	fileNameRegex := regexp.MustCompile(fileNameRegexStr)

	for i, f := range files {
		fileName := f.Name()
		filePath := fmt.Sprintf("%s%s", RECIPES_DIR, fileName)
		fmt.Printf("%d. %s\n", i+1, filePath)

		// TODO: collect up more errors before erroring out

		if !fileNameRegex.Match([]byte(fileName)) {
			fmt.Printf("💣 filename '%s' does not conform to lower case, snake case, or missing extension '.yaml'\n", fileName)
			continue
		}

		contents, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("💣error trying to read file: %v", err)
		}

		recipe := Recipe{}
		err = yaml.Unmarshal([]byte(contents), &recipe)
		if err != nil {
			fmt.Printf("💣error unmarhsalling %s: %v\n", filePath, err)
			err = nil
			continue
		}

		err = validate.Struct(recipe)
		if err != nil {
			fmt.Printf("💣error validating struct for %s\n%v\n", filePath, err)
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
			fmt.Printf("💣recipe name '%s' should be the title case variation of file name '%s'\n", recipe.Name, fileName)
			fmt.Println("        recipe.Name:", recipe.Name)
			fmt.Println("fileNameAsTitleCase:", fileNameAsTitleCase)
			err = nil
			continue
		}

		passCount += 1
		fmt.Println("💚 PASS")
	}

	errCount := len(files) - passCount

	var emoji string
	if errCount == 0 {
		emoji = "😎😎😎😎"
	} else {
		emoji = "💣💣💣💣"
	}
	fmt.Printf("\n%s FAIL = %d, PASS = %d %s\n", emoji, errCount, passCount, emoji)
	os.Exit(errCount)
}
