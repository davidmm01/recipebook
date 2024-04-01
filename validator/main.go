package main

import (
	"flag"
	"log"
	"slices"
	"strings"
)

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
		getValidRecipes(true)
	case "generate":
		recipes := getValidRecipes(false)
		generateAll(recipes)
	}
}
