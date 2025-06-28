# Fizz - Interactive Fuzzy File Finder

Fizz is a fast, interactive file finder for the command line that uses fuzzy search to help you quickly locate files in a directory.

## Features

- **Fast File Scanning**: Uses parallel processing to quickly scan directories
- **Fuzzy Search**: Find files by typing partial matches
- **Real-time Results**: Results update as you type
- **Interactive Interface**: Simple keyboard-based navigation

## Installation

### Prerequisites

- Go 1.24 or higher

### Building from Source

```bash
git clone https://github.com/pyljain/fizz.git
cd fizz
go build
```

## Usage

```bash
./fizz [directory]
```

If no directory is specified, Fizz will search in the current directory.

### Keyboard Controls

- Type to search
- `Backspace` or `Delete`: Remove the last character from the search
- `Esc` or `Ctrl+C`: Exit the application

## How It Works

Fizz scans the specified directory for files and distributes them across multiple lists for parallel processing. As you type, it performs fuzzy matching on these lists and displays the best matches in real-time.

The application uses:
- `fastwalk` for efficient directory traversal
- `fuzzysearch` for fuzzy string matching
- `pterm` for terminal UI rendering
- `keyboard` for handling keyboard input

## Dependencies

- [atomicgo.dev/keyboard](https://github.com/atomicgo/keyboard) - Keyboard event handling
- [github.com/charlievieth/fastwalk](https://github.com/charlievieth/fastwalk) - Fast directory walking
- [github.com/lithammer/fuzzysearch](https://github.com/lithammer/fuzzysearch) - Fuzzy search algorithms
- [github.com/pterm/pterm](https://github.com/pterm/pterm) - Terminal UI library

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.