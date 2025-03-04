package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-playground/calculator"
)

func main() {
	fmt.Println("Welcome to Go Calculator!")
	fmt.Println("Available operations: add, subtract, multiply, divide, power, sqrt, sin, cos, tan, log, factorial")
	fmt.Println("Enter 'exit' to quit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if input == "exit" {
			break
		}

		result, err := processInput(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Result: %v\n", result)
		}
	}

	fmt.Println("Thank you for using Go Calculator!")
}

func processInput(input string) (interface{}, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("no input provided")
	}

	operation := strings.ToLower(parts[0])

	switch operation {
	case "add":
		if len(parts) != 3 {
			return nil, fmt.Errorf("usage: add <number> <number>")
		}
		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Add(a, b), nil

	case "subtract":
		if len(parts) != 3 {
			return nil, fmt.Errorf("usage: subtract <number> <number>")
		}
		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Subtract(a, b), nil

	case "multiply":
		if len(parts) != 3 {
			return nil, fmt.Errorf("usage: multiply <number> <number>")
		}
		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Multiply(a, b), nil

	case "divide":
		if len(parts) != 3 {
			return nil, fmt.Errorf("usage: divide <number> <number>")
		}
		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Divide(a, b)

	case "power":
		if len(parts) != 3 {
			return nil, fmt.Errorf("usage: power <base> <exponent>")
		}
		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Power(a, b), nil

	case "sqrt":
		if len(parts) != 2 {
			return nil, fmt.Errorf("usage: sqrt <number>")
		}
		a, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.SquareRoot(a)

	case "sin":
		if len(parts) != 2 {
			return nil, fmt.Errorf("usage: sin <angle_in_radians>")
		}
		a, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Sine(a), nil

	case "cos":
		if len(parts) != 2 {
			return nil, fmt.Errorf("usage: cos <angle_in_radians>")
		}
		a, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Cosine(a), nil

	case "tan":
		if len(parts) != 2 {
			return nil, fmt.Errorf("usage: tan <angle_in_radians>")
		}
		a, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Tangent(a), nil

	case "log":
		if len(parts) != 3 {
			return nil, fmt.Errorf("usage: log <number> <base>")
		}
		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number format")
		}
		return calculator.Logarithm(a, b)

	case "factorial":
		if len(parts) != 2 {
			return nil, fmt.Errorf("usage: factorial <non-negative_integer>")
		}
		a, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid integer format")
		}
		return calculator.Factorial(a)

	default:
		return nil, fmt.Errorf("unknown operation: %s", operation)
	}
}
