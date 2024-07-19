# Templ CSS Sort

This project is a Go-based tool for sorting CSS classes in `.templ` fiels. 

The Makefile provided helps in compiling the binary for different operating systems and architectures.

## Prerequisites

- Go (I used v1.22.5 in development)
- Git

## Getting Started

### Cloning the Repository

First, clone the repository:

```sh
git clone https://github.com/yourusername/templ-css-sort.git
cd templ-css-sort
```

### Building the Project

The Makefile includes several targets to help with building the project. Below are the steps to compile the binary for different platforms.

#### Vendor Dependencies

Vendor the Go module dependencies:

```sh
make vendor
```

#### Compile the Binary

Compile the binary for different platforms:

```sh
make dist
```

This will create the following binaries in the `bin/` directory:

- `templ-css-sort-amd64` (Linux)
- `templ-css-sort-darwin` (macOS)
- `templ-css-sort-arm` (ARM Linux)
- `templ-css-sort-arm64` (ARM64 Linux)
- `templ-css-sort.exe` (Windows)

## Usage

After building the binary, you can use it to sort your CSS files. For example:

```sh
# sort all files in the current working directory and all subdirectories
./bin/templ-css-sort-amd64

# sort a given folder of .templ files
./bin/templ-css-sort-amd64 --dir path/to/your/templates/folder/

# sort a given .templ file
./bin/templ-css-sort-amd64 --file path/to/your/templates/folder/pages.templ

# sort from stdin
echo 'class="b a c"' | ./bin/templ-css-sort-amd64 --stdin
```

Replace `templ-css-sort-amd64` with the appropriate binary for your operating system.

You can also rename the binary to a shorter name and move it to a directory that's in your system's `PATH` to make it accessible from anywhere. For example:

### Linux and macOS

For Linux and macOS, you can use the `mv` command to move the binary to `/usr/local/bin`, which is typically in the system's `PATH`.

```sh
# linux
mv ./bin/templ-css-sort-amd64 /usr/local/bin/tcsort

# mac os
mv ./bin/templ-css-sort-darwin /usr/local/bin/tcsort
```

### Windows

For Windows, you can use the `move` command to move the binary to a directory in your `PATH`. One common directory is `C:\Windows\System32`.

```sh
move .\bin\templ-css-sort.exe C:\Windows\System32\tcsort.exe
```

Now you can use the tool from anywhere on your system:

```sh
tcsort --dir path/to/your/templates/folder/
```

## Run on Save

### VS Code

1. install tcsort and make sure it's available globally
2. install emeraldwalk.runonsave extension from VS Code marketplace
3. Configure it to run `tcsort` on the saved file

```json
{
    "emeraldwalk.runonsave": {
        "commands": [
            {
            "match": "\\.templ$",
            "cmd": "tcsort --file ${file}"
            }
        ]
    }
}
```

### Nvim with --file

1. install tcsort and make sure it's available globally
2. install emeraldwalk.runonsave extension from VS Code marketplace
3. Configure it to run `tcsort` on the saved file

Either in `init.vim`:

```
augroup tcsort
  autocmd!
  autocmd BufWritePre *.templ :silent! execute '!tcsort --file %'
augroup END
```

Or in `init.lua`:

```lua
vim.api.nvim_exec([[
  augroup tcsort
    autocmd!
    autocmd BufWritePre *.templ :silent! execute '!tcsort --file %'
  augroup END
]], false)
```

### Nvim with --stdin

You can also use stdin/stdout mode for nvim. This is an example of `init.lua`.

```lua
-- Create an augroup for tcsort
local tcsort_group = vim.api.nvim_create_augroup("tcsort", { clear = true })

-- Create an autocmd for BufWritePre event
vim.api.nvim_create_autocmd("BufWritePre", {
  group = tcsort_group,
  pattern = "*.templ",
  command = "silent! %!tcsort --stdin",
})
```

## Contributing

Feel free to open issues or submit pull requests if you find any bugs or have suggestions for improvements.

## License

This project is licensed under the MIT License.