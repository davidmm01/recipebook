package main

import (
	"flag"
	"log"
	"slices"
	"strings"
)

// relative path from makefile to the hugo sites recipes posts
const CONTENT_RECIPES_DIR = "content/recipes/"

// relative path from makefile to the hugo sites cocktails posts
const CONTENT_COCKTAILS_DIR = "content/cocktails/"

// relative path from makefile to the recipe source files
const RECIPES_DIR = "source/recipes/"

// relative path from makefile to the cocktails source files
const COCKTAILS_DIR = "source/cocktails/"

var supportedCommands = []string{
	"validate", "generate",
}

func main() {
	var command string
	flag.StringVar(&command, "command", "", strings.Join(supportedCommands[:], ", "))

	flag.Parse()

	if !slices.Contains(supportedCommands, command) {
		log.Fatalf("provided command='%s' not one of supported commands=%v", command, strings.Join(supportedCommands[:], ", "))
	}

	switch command {
	case "validate":
		getValidRecipes(RECIPES_DIR, true)
	case "generate":
		recipes := getValidRecipes(RECIPES_DIR, false)
		cocktails := getValidRecipes(COCKTAILS_DIR, false)
		writeRecipesAsMarkdownFile(recipes, CONTENT_RECIPES_DIR)
		writeRecipesAsMarkdownFile(cocktails, CONTENT_COCKTAILS_DIR)
	}
}
