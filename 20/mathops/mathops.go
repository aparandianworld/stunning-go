package mathops

import (
	"errors"
)

// basic math operations

func Add(a, b float64) (float64, error) {
	return a + b, nil
}

func Sub(a, b float64) (float64, error) {
	return a - b, nil
}

func Mul(a, b float64) (float64, error) {
	return a * b, nil
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
