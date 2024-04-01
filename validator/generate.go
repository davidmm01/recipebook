package main

import (
	"fmt"
	"os"
	"strings"
)

const CONTENT_DIR = "content/posts/"

func generateAll(recipes []Recipe) {
	for _, recipe := range recipes {
		markdown := generateMarkdown(recipe)
		markdownBytes := []byte(markdown)
		filename := strings.ToLower(recipe.Name)
		filename = strings.ReplaceAll(filename, " ", "_")
		filename += ".md"
		filePath := fmt.Sprintf("%s%s", CONTENT_DIR, filename)
		os.WriteFile(filePath, markdownBytes, os.ModePerm)
	}
}

func generateMarkdown(recipe Recipe) string {
	markdown := "---\n"
	markdown += fmt.Sprintf("title: %s\n", recipe.Name)
	markdown += fmt.Sprintf("date: %s\n", recipe.DateAdded)
	markdown += "draft: false\n"
	markdown += fmt.Sprintf("cuisine: %s\n", recipe.Cuisine)

	tags := "["
	for _, descriptor := range recipe.Descriptors {
		tags += fmt.Sprintf("\"%s\", ", descriptor)
	}
	tags = strings.TrimRight(tags, ", ")
	tags += "]\n"
	markdown += fmt.Sprintf("tags: %s", tags)
	markdown += "---\n\n"

	markdown += "## Ingredients\n"
	for _, ingredient := range recipe.Ingredients {
		markdown += fmt.Sprintf("- %s\n", ingredient)
	}
	markdown += "\n"

	markdown += "## Instructions\n"
	for _, instruction := range recipe.Instructions {
		markdown += fmt.Sprintf("- %s\n", instruction)
	}
	markdown += "\n"

	return markdown
}
