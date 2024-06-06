package calculator_test

import (
	"github.com/PacktPublishing/Test-Driven-Development-in-Go/chapter02/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	x, y := 2.5, 3.5
	want := x + y
	e := calculator.Engine{}

	got := e.Add(x, y)

	if got != want {
		t.Errorf("Add(%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}

func TestSubtract(t *testing.T) {
	x, y := 2.5, 3.5
	want := x - y
	e := calculator.Engine{}

	got := e.Subtract(x, y)

	if got != want {
		t.Errorf("Subtract(%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}

func TestMultiply(t *testing.T) {
	x, y := 2.5, 3.5
	want := x * y
	e := calculator.Engine{}

	got := e.Multiply(x, y)

	if got != want {
		t.Errorf("Multiply(%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}

func TestDivide(t *testing.T) {
	x, y := 2.5, 3.5
	want := x / y
	e := calculator.Engine{}

	got, _ := e.Divide(x, y)

	if got != want {
		t.Errorf("Divide(%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}

func TestDivideByZero(t *testing.T) {
	x, y := 2.5, 0.0
	e := calculator.Engine{}

	_, err := e.Divide(x, y)

	if err == nil {
		t.Error("Division by zero should return an error")
	}
}
