package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
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

			// trim in place
			classList = strings.TrimSpace(classList)

			// any whitespace bigger then 1 char, reduce to 1 char
			for strings.Contains(classList, "  ") {
				classList = strings.ReplaceAll(classList, "  ", " ")
			}

			// split
			classes := strings.Split(classList, " ")

			// sort
			sort.Strings(classes)

			// remove duplicates
			classes = removeDuplicates(classes)

			log.Println("File:", file, "Classes:", classes)
		}
	}
}

func removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
