package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Flags struct {
	dev   bool
	file  string
	dir   string
	stdin bool
}

func main() {
	// flag "--dev" to run the script in dev mode
	dev := flag.Bool("dev", false, "Run the script in dev mode")
	// flag "--file" to specify the file to sort
	file := flag.String("file", "", "Specify the file to sort")
	// flag "--dir" to specify the directory to sort
	dir := flag.String("dir", "", "Specify the directory to sort")
	// flag "--stdin" to read from stdin
	stdin := flag.Bool("stdin", false, "Read from stdin")
	flag.Parse()

	templCSSSort(Flags{
		dev:   *dev,
		file:  *file,
		dir:   *dir,
		stdin: *stdin,
	})
}

// The main function that runs the templCSSSort.
// Takes in flags to determine what to sort and where to sort it.
// If no flags are set, it will sort all .templ files in the current directory and its subdirectories.
// With the --stdin flag, it will read from stdin and write to stdout.
// With the --file flag, it will only sort the specified file.
// With the --dir flag, it will only sort the specified directory and its subdirectories.
func templCSSSort(flags Flags) {
	start := time.Now()

	// if stdin flag is set, read from stdin and write to stdout
	if flags.stdin {
		content, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		newContent, err := processContent(string(content))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(newContent)
		return
	}

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
			log.Println("couldn't read file: ", file)
			continue
		}
		if len(content) == 0 {
			log.Println("file empty: ", file)
			continue
		}

		newContent, err := processContent(string(content))
		if err != nil {
			if err.Error() == "no_classes_found" {
				continue
			}
			log.Fatal(err)
		}

		// write the modified content back to the file
		err = os.WriteFile(file, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Done in", time.Since(start))
}

// Takes a string of content and returns a string of content with the classes sorted
func processContent(content string) (string, error) {
	// find all classes in content
	re := regexp.MustCompile(`class="([^"]+)"`)
	matches := re.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return "", errors.New("no_classes_found")
	}

	for _, match := range matches {
		classList := match[1]
		if classList == "" {
			continue
		}

		// trim in place
		classList = strings.TrimSpace(classList)
		if classList == "" {
			continue
		}

		// any whitespace bigger than 1 char, reduce to 1 char
		for strings.Contains(classList, "  ") {
			classList = strings.ReplaceAll(classList, "  ", " ")
		}
		if classList == "" {
			continue
		}

		// split
		classes := strings.Split(classList, " ")
		if len(classes) == 0 {
			continue
		}

		// sort
		sort.Strings(classes)

		// remove duplicates
		classes = removeDuplicates(classes)

		// create new class list string
		newClassList := strings.Join(classes, " ")
		if newClassList == "" {
			continue
		}

		// replace class list in file
		content = strings.Replace(content, match[0], `class="`+newClassList+`"`, -1)
		if content == "" {
			continue
		}
	}

	return content, nil
}

// Takes a slice of strings and returns a slice of strings with duplicates removed
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
