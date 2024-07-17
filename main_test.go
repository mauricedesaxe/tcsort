package main

import (
	"os"
	"testing"
)

func TestSetupTemplFile(t *testing.T) {
	// Create the templates directory if it doesn't exist
	err := os.MkdirAll("./templates", os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create templates directory: %v", err)
	}

	// Define the content to be written to the .templ file
	content := `
	package templates

	templ list() {
		<ul role="list" class="mt-8 space-y-3 text-sm leading-6 text-gray-600 xl:mt-10">
			<li class="flex gap-x-3">
				<svg class="h-6 w-5 flex-none text-blue-600" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
					<path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd"></path>
				</svg>
				you get all the features in the "1 month" plan
			</li>
			<li class="flex gap-x-3">
				<svg class="h-6 w-5 flex-none text-blue-600" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
					<path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd"></path>
				</svg>
				$200 discount
			</li>
		</ul>
	}
	`

	// Write the content to the .templ file
	err = os.WriteFile("./templates/test.templ", []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write to test.templ file: %v", err)
	}
}

func TestMain(t *testing.T) {
	main()
}
