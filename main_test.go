package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"
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

func TestSetupTemplFile2(t *testing.T) {
	// Create the templates directory if it doesn't exist
	err := os.MkdirAll("./templates", os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create templates directory: %v", err)
	}

	// Define the content to be written to the .templ file
	content := `
package billing

import "github.com/dexprism/data-processor/general"

type PlansPageProps struct {
	Success string
	Error   string
}

templ plans_page(props PlansPageProps) {
	@general.Base() {
		@general.Sidebar("", "", "Plans", "icon_profile", []general.NavItem{
			{Href: "/profile", Label: "Profile"},
			{Href: "/plans", Label: "Plans", Active: true},
		}) {
			<main class="mx-auto container px-4 py-4">
				<div class="empty:hidden bg-green-200 text-green-600 p-4 rounded-md">
					{ general.TernaryIf(props.Success != "", "🟢 " + props.Success, "") }
				</div>
				<div class="empty:hidden bg-red-200 text-red-6000 p-4 rounded-md">
					{ general.TernaryIf(props.Error != "", "🔴 " + props.Error, "") }
				</div>
				<div class="mx-auto max-w-4xl text-center">
					<h2 class="text-base font-semibold leading-7 text-blue-600 uppercase">Pre-paid plans</h2>
					<p class="mt-2 text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
						More data, more insights
					</p>
				</div>
				<p class="mx-auto mt-6 max-w-2xl text-center text-lg leading-8 text-gray-600">
					If you need data older than 3 months or very granular (per minute) data 
					to better understand Chainlink and your operation, you can use one of our pre-paid plans.
				</p>
				<div class="isolate mx-auto mt-10 grid max-w-md grid-cols-1 gap-8 lg:mx-0 lg:max-w-none lg:grid-cols-3">
					<div class="rounded-3xl p-8 ring-1 ring-gray-200 xl:p-10 transition-all duration-300 hover:shadow-lg">
						<h3 id="tier-month" class="text-lg font-semibold leading-8 text-gray-900">
							1 month
						</h3>
						<p class="mt-4 text-sm leading-6 text-gray-600">
							Best when you're not sure that you'll use the data regularly.
						</p>
						<p class="mt-6 flex items-baseline gap-x-1">
							<span class="text-4xl font-bold tracking-tight text-gray-900">$100</span>
						</p>
						<a href="/plans/1-month" aria-describedby="tier-freelancer" class="mt-6 block rounded-md bg-blue-600 px-3 py-2 text-center text-sm font-semibold leading-6 text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600">Buy 1-month plan</a>
						<ul role="list" class="mt-8 space-y-3 text-sm leading-6 text-gray-600 xl:mt-10">
							<li class="flex gap-x-3">
								<svg class="h-6 w-5 flex-none text-blue-600" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
									<path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd"></path>
								</svg>
								data going back more than 3 months
							</li>
							<li class="flex gap-x-3">
								<svg class="h-6 w-5 flex-none text-blue-600" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
									<path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd"></path>
								</svg>
								data as granular as "per hour" or "per minute"
							</li>
						</ul>
					</div>
					<div class="rounded-3xl p-8 ring-1 ring-gray-200 xl:p-10 transition-all duration-300 hover:shadow-lg">
						<h3 id="tier-year" class="text-lg font-semibold leading-8 text-gray-900">
							1 year
						</h3>
						<p class="mt-4 text-sm leading-6 text-gray-600">
							Best when you already love the product and know the data is critical to you.
						</p>
						<p class="mt-6 flex items-baseline gap-x-1">
							<span class="text-4xl font-bold tracking-tight text-gray-900">$1,000</span>
						</p>
						<a href="/plans/1-year" aria-describedby="tier-startup" class="mt-6 block rounded-md bg-blue-600 px-3 py-2 text-center text-sm font-semibold leading-6 text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600">Buy 1-year plan</a>
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
					</div>
					<div class="rounded-3xl bg-gray-900 p-8 ring-1 ring-gray-900 xl:p-10 transition-all duration-300 hover:shadow-lg">
						<h3 id="tier-enterprise" class="text-lg font-semibold leading-8 text-white">
							Custom
						</h3>
						<p class="mt-4 text-sm leading-6 text-gray-300">
							If you need custom data or anything else that's not here, we can build it.
						</p>
						<p class="mt-6 flex items-baseline gap-x-1">
							<span class="text-4xl font-bold tracking-tight text-white">
								Custom
							</span>
						</p>
						<a href="#" aria-describedby="tier-enterprise" class="mt-6 block rounded-md bg-white/10 px-3 py-2 text-center text-sm font-semibold leading-6 text-white hover:bg-white/20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white">Contact sales</a>
						<ul role="list" class="mt-8 space-y-3 text-sm leading-6 text-gray-300 xl:mt-10">
							<li class="flex gap-x-3">
								<svg class="h-6 w-5 flex-none text-white" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
									<path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd"></path>
								</svg>
								Unlimited data
							</li>
							<li class="flex gap-x-3">
								<svg class="h-6 w-5 flex-none text-white" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
									<path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd"></path>
								</svg>
								Unlimited granularity
							</li>
							<li class="flex gap-x-3">
								<svg class="h-6 w-5 flex-none text-white" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
									<path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd"></path>
								</svg>
								Dedicated support from core devs
							</li>
						</ul>
					</div>
				</div>
			</main>
		}
	}
}
`

	// Write the content to the .templ file
	err = os.WriteFile("./templates/test2.templ", []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write to test.templ file: %v", err)
	}
}

func TestTemplCSSSort(t *testing.T) {
	start := time.Now()
	templCSSSort(Flags{
		dev: true,
	})
	end := time.Now()
	fmt.Println("Time taken:", end.Sub(start))
}

func TestStdinToStdout(t *testing.T) {
	// Simulate stdin input
	input := "class=\"b a c\""
	r, w, _ := os.Pipe()
	w.Write([]byte(input))
	w.Close()
	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	// Capture stdout
	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut
	defer func() { os.Stdout = oldStdout }()

	// Run the function
	templCSSSort(Flags{stdin: true})

	// Close the writer and read the output
	wOut.Close()
	var buf bytes.Buffer
	buf.ReadFrom(rOut)

	// Check the output
	expectedOutput := "class=\"a b c\"\n"
	if buf.String() != expectedOutput {
		t.Errorf("expected %q, got %q", expectedOutput, buf.String())
	}
}
