package calculator

import (
	"fmt"
	"math"
)

// Factorial calculates the factorial of a non-negative integer
// Returns an error if the input is negative
func Factorial(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}

	if n == 0 || n == 1 {
		return 1, nil
	}

	var result uint64 = 1
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}

	return result, nil
}

// Logarithm calculates the logarithm of a number with the specified base
// Returns an error for invalid inputs
func Logarithm(value, base float64) (float64, error) {
	if value <= 0 {
		return 0, fmt.Errorf("logarithm is not defined for non-positive numbers")
	}

	if base <= 0 || base == 1 {
		return 0, fmt.Errorf("invalid logarithm base")
	}

	return math.Log(value) / math.Log(base), nil
}

// Sine calculates the sine of an angle in radians
func Sine(angle float64) float64 {
	return math.Sin(angle)
}

// Cosine calculates the cosine of an angle in radians
func Cosine(angle float64) float64 {
	return math.Cos(angle)
}

// Tangent calculates the tangent of an angle in radians
func Tangent(angle float64) float64 {
	return math.Tan(angle)
}

// DegreesToRadians converts an angle from degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// RadiansToDegrees converts an angle from radians to degrees
func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
