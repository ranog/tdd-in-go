package calculator

import "fmt"

type Engine struct{}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}

func (e *Engine) Subtract(x, y float64) float64 {
	return x - y
}

func (e *Engine) Multiply(x, y float64) float64 {
	return x * y
}

func (e *Engine) Divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return x / y, nil
}
