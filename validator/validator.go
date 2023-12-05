package main

import (
	"fmt"
	"os"
)

func main() {

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
