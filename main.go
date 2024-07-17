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

		originalContent := string(content)

		// find all classes in templ file
		re := regexp.MustCompile(`class="([^"]+)"`)
		matches := re.FindAllStringSubmatch(originalContent, -1)

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

			// create new class list string
			newClassList := strings.Join(classes, " ")

			// log diff
			logDiff(file, classList, newClassList)

			// replace class list in file
			originalContent = strings.Replace(originalContent, match[0], "class=\""+newClassList+"\"", -1)
		}

		// write the modified content back to the file
		err = os.WriteFile(file, []byte(originalContent), 0644)
		if err != nil {
			log.Fatal(err)
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

func logDiff(file, oldClassList, newClassList string) {
	const (
		red   = "\033[31m"
		green = "\033[32m"
		reset = "\033[0m"
	)

	log.Println("File:", file)
	log.Println("Old:", red, oldClassList, reset)
	log.Println("New:", green, newClassList, reset)
}
