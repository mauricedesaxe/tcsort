package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// find all .templ files in directory and subdirectories
	files, err := filepath.Glob("./templates/*.templ")
	if err != nil {
		log.Fatal(err)
	}

	// parse each file
	for _, file := range files {
		os.ReadFile(file)
	}
}
