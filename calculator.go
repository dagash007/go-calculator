package calculator

import "errors"

// Divide returns the quotient of a divided by b
// Returns an error if b is 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Square returns the square of a
func Square(a int) int {
	return a * a
}
