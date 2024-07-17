package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// find all .templ files in directory and subdirectories
	files, err := filepath.Glob("./templates/*.templ")
	if err != nil {
		log.Fatal(err)
	}

	// parse each file
	for _, file := range files {
		// read file
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		// find all classes in templ file
		re := regexp.MustCompile(`class="([^"]+)"`)
		matches := re.FindAllStringSubmatch(string(content), -1)

		for _, match := range matches {
			classList := match[1]
			log.Println("File:", file, "ClassList:", classList)
		}
	}
}
