package main

import (
	"fmt"
	"os"
	"strings"
)

func writeRecipesAsMarkdownFile(recipes []Recipe, destDir string) {
	for _, recipe := range recipes {
		markdown := generateRecipeMarkdown(recipe)
		markdownBytes := []byte(markdown)
		filename := strings.ToLower(recipe.Name)
		filename = strings.ReplaceAll(filename, " ", "_")
		filename += ".md"
		filePath := fmt.Sprintf("%s%s", destDir, filename)
		os.WriteFile(filePath, markdownBytes, os.ModePerm)
	}
}

func generateRecipeMarkdown(recipe Recipe) string {
	markdown := "---\n"
	markdown += fmt.Sprintf("title: %s\n", recipe.Name)
	markdown += fmt.Sprintf("date: %s\n", recipe.DateAdded)
	markdown += "draft: false\n"
	markdown += fmt.Sprintf("cuisines: [\"%s\"]\n", recipe.Cuisine)

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
