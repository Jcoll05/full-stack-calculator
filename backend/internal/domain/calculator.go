package domain

import (
	"errors"
	"math"
)

type Operation string

// Supported operations
const (
	Add        Operation = "add"
	Subtract   Operation = "subtract"
	Multiply   Operation = "multiply"
	Divide     Operation = "divide"
	Power      Operation = "power" // Exponentiation
	SquareRoot Operation = "sqrt"
	Percentage Operation = "percentage"
)

// Validation errors
var (
	ErrUnsupportedOperation = errors.New("unsupported operation")
	ErrDivisionByZero       = errors.New("cannot divide by zero")
	ErrNegativeSquareRoot   = errors.New("cannot calculate square root of a negative number")
)

// Calculate performs the specified arithmetic operation and returns
// the result or an error if the operation is unsupported or invalid.
func Calculate(operation Operation, a, b float64) (float64, error) {
	switch operation {
	case Add:
		return a + b, nil

	case Subtract:
		return a - b, nil

	case Multiply:
		return a * b, nil

	case Divide:
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil

	case Power:
		return math.Pow(a, b), nil

	case SquareRoot:
		if a < 0 {
			return 0, ErrNegativeSquareRoot
		}
		return math.Sqrt(a), nil

	case Percentage:
		return a / 100, nil

	default:
		return 0, ErrUnsupportedOperation
	}
}
