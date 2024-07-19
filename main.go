package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Flags struct {
	dev  bool
	file string
	dir  string
}

func main() {
	// flag "--dev" to run the script in dev mode
	dev := flag.Bool("dev", false, "Run the script in dev mode")
	// flag "--file" to specify the file to sort
	file := flag.String("file", "", "Specify the file to sort")
	// flag "--dir" to specify the directory to sort
	dir := flag.String("dir", "", "Specify the directory to sort")
	flag.Parse()

	templCSSSort(Flags{
		dev:  *dev,
		file: *file,
		dir:  *dir,
	})
}

func assert(condition bool, msg string) {
	if !condition {
		log.Fatal(msg)
	}
}

func templCSSSort(flags Flags) {
	start := time.Now()

	// find all .templ files in directory and subdirectories
	var files []string
	var err error
	if flags.file != "" {
		// If the file flag is specified, only take in that file
		if !strings.HasSuffix(flags.file, ".templ") {
			log.Fatal("File must have .templ extension")
		}
		files = append(files, flags.file)
	} else if flags.dir != "" {
		// If the dir flag is specified, take in that dir and its subdirectories
		err = filepath.Walk(flags.dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".templ") {
				files = append(files, path)
			}
			return nil
		})
	} else {
		// If neither flag is specified, go through cwd and all subdirectories for any .templ file
		err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".templ") {
				files = append(files, path)
			}
			return nil
		})
	}
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Found", len(files), "files")

	// parse each file
	for _, file := range files {
		// read file
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		originalContent := string(content)
		assert(originalContent != "", "File is empty")

		newContent := processContent(originalContent)

		// write the modified content back to the file
		err = os.WriteFile(file, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Done in", time.Since(start))
}

func processContent(content string) string {
	// find all classes in content
	re := regexp.MustCompile(`class="([^"]+)"`)
	matches := re.FindAllStringSubmatch(content, -1)
	assert(len(matches) > 0, "No classes found")

	for _, match := range matches {
		classList := match[1]
		assert(classList != "", "Class list is empty")

		// trim in place
		classList = strings.TrimSpace(classList)
		assert(classList != "", "Class list is empty")

		// any whitespace bigger then 1 char, reduce to 1 char
		for strings.Contains(classList, "  ") {
			classList = strings.ReplaceAll(classList, "  ", " ")
		}
		assert(classList != "", "Class list is empty")

		// split
		classes := strings.Split(classList, " ")
		assert(len(classes) > 0, "No classes found")

		// sort
		sort.Strings(classes)

		// remove duplicates
		classes = removeDuplicates(classes)

		// create new class list string
		newClassList := strings.Join(classes, " ")
		assert(newClassList != "", "New class list is empty")

		// replace class list in file
		content = strings.Replace(content, match[0], "class=\""+newClassList+"\"", -1)
		assert(content != "", "New content is empty")
	}

	return content
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
	assert(len(list) > 0, "No classes found")
	return list
}
