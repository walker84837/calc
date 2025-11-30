# Calculator

A simple GUI calculator application built with Go, Fyne, and govaluate.

## Features

* Basic arithmetic operations (+, -, *, /)
* Number input (0-9)
* Decimal point (.)
* Parentheses support
* Clear button (C)
* Display for input and results
* Basic error handling for invalid expressions

### Technologies Used

* **Go**: The primary programming language.
* **Fyne**: A cross-platform GUI toolkit for Go.
* **govaluate**: A library for evaluating mathematical and logical expressions.

## Installation & Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/walker84837/calc
   cd calc
   ```
2. **Install Go**: If you don't have Go installed, follow the official guide: <https://go.dev/doc/install>
3. **Build the project**:
   ```bash
   go build -o calculator ./cmd/calc/
   ```
   This will create an executable file named `calculator` in the project's root directory.

## Running

After building, you can run the application:
```bash
./calculator
```
This will launch the GUI calculator window.

## License

This project is licensed under the BSD Zero Clause License (see the [LICENSE](LICENSE) file for more details). This means you can do practically anything you want with the code, including commercial use.
