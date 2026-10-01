// Package calculator реалізує прості арифметичні операції для Завдання 0
// (об'єднання calculator із Заняття 1 в multi-package проєкт).
package calculator

import (
	"errors"
	"fmt"
)

// ErrDivisionByZero — sentinel error для ділення на нуль.
var ErrDivisionByZero = errors.New("calculator: division by zero")

// Add повертає суму a та b.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract повертає різницю a - b.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply повертає добуток a та b.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide повертає частку a / b. Якщо b дорівнює 0, повертає помилку,
// що через %w обгортає ErrDivisionByZero (щоб виклики могли перевіряти
// її через errors.Is).
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("calculator: divide %v by %v: %w", a, b, ErrDivisionByZero)
	}
	return a / b, nil
}
