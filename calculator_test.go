package calculator

import (
	"testing"
)

// TestDivide checks the Divide function
func TestDivide(t *testing.T) {
	// Normal division case
	result, err := Divide(10, 2)
	if err != nil || result != 5 {
		t.Errorf("Divide(10, 2) failed: expected 5, got %d", result)
	}

	// Division by zero case
	_, err = Divide(10, 0)
	if err == nil {
		t.Errorf("Divide(10, 0) did not return an error as expected")
	}
}

// TestSquare checks the Square function
func TestSquare(t *testing.T) {
	result := Square(5)
	if result != 25 {
		t.Errorf("Square(5) failed: expected 25, got %d", result)
	}

	result = Square(-3)
	if result != 9 {
		t.Errorf("Square(-3) failed: expected 9, got %d", result)
	}
}

// BenchmarkDivide benchmarks the Divide function
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(100, 5)
	}
}

// BenchmarkSquare benchmarks the Square function
func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(10)
	}
}
