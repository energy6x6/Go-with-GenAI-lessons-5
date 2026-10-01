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
//
// TODO(Завдання 0): реалізуйте.
func Add(a, b float64) float64 {
	panic("not implemented")
}

// Subtract повертає різницю a - b.
//
// TODO(Завдання 0): реалізуйте.
func Subtract(a, b float64) float64 {
	panic("not implemented")
}

// Multiply повертає добуток a та b.
//
// TODO(Завдання 0): реалізуйте.
func Multiply(a, b float64) float64 {
	panic("not implemented")
}

// Divide повертає частку a / b. Якщо b дорівнює 0, повертає помилку,
// що через %w обгортає ErrDivisionByZero (щоб виклики могли перевіряти
// її через errors.Is).
//
// TODO(Завдання 0): реалізуйте, застосувавши патерни error wrapping
// із заняття 5 (fmt.Errorf + %w).
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("calculator: divide %v by %v: %w", a, b, ErrDivisionByZero)
	}
	panic("not implemented")
}
