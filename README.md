# Torri Interpreter

**Torri** is a custom interpreter built in **Golang**, featuring a **REPL (Read-Eval-Print Loop)** for intuitive scripting. It includes essential programming functions, variable handling, and control flow, using both English and Nepali keywords, making it a great tool for scripting and educational use.

## Features

- **Tokenization**: Supports various operators, delimiters, and keywords in both English and Nepali.
  - Operators: `=`, `+`, `-`, `*`, `/`, `%`
  - Delimiters: `,`, `;`, `:`, `()`, `{}`, `[]`
  - Keywords:
    - `yedi` - if
    - `kita` - else if
    - `natra` - else
    - `firta` - return
    - `la` - let
- **Built-in Functions**:
  - `len`: Returns the length of an array.
  - `push`: Adds an element to an array.
  - `first` / `last`: Gets the first or last element of an array.
  - `lekh`: Prints output to the terminal (Nepali for "write").
- **Variable Management**: Flexible identifiers for defining and using variables.

## Getting Started

### Prerequisites

- **Golang** should be installed on your system.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Sarthak-Neupane/Torri.git

2. Navigate to the project directory:
   ```bash
   cd torri

3. Build the project:
   ```bash
    go build -o torri

4. Run the executable:
    ```bash
    ./torri run path/to/yourfile.tor
    ```
    or
    
    ```bash
    ./torri
    ```

## Usage

### REPL

- Run the executable without any arguments to start the REPL.
- Enter your code and press `Enter` to execute it.
- Use `Ctrl + C` to exit the REPL.

### Scripting

- Create a `.tor` file with your code.
- Run the executable with the path to your file to execute it.

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue if you have any suggestions or feedback.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.





