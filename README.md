# Go Calculator

A simple command-line calculator application written in Go.

## Features

- Basic arithmetic operations: addition, subtraction, multiplication, division
- Advanced mathematical functions: power, square root, logarithm, factorial
- Trigonometric functions: sine, cosine, tangent
- Angle conversion: degrees to radians and radians to degrees

## Usage

1. Build the application:
   ```
   go build
   ```

2. Run the application:
   ```
   ./go-playground
   ```

3. Enter commands in the format:
   ```
   operation [arguments]
   ```

   Examples:
   - `add 5 3` (Result: 8)
   - `subtract 10 4` (Result: 6)
   - `multiply 2.5 3` (Result: 7.5)
   - `divide 10 2` (Result: 5)
   - `power 2 3` (Result: 8)
   - `sqrt 16` (Result: 4)
   - `sin 1.5708` (Result: ~1, which is sin(π/2))
   - `cos 0` (Result: 1)
   - `tan 0.7854` (Result: ~1, which is tan(π/4))
   - `log 100 10` (Result: 2, which is log₁₀(100))
   - `factorial 5` (Result: 120)

4. Type `exit` to quit the application.

## Project Structure

- `main.go`: Contains the main application logic and user interface
- `calculator/calculator.go`: Implements basic calculator operations
- `calculator/advanced.go`: Implements advanced mathematical functions
