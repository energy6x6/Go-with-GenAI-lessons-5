package calculator

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %v, want 5", got)
	}
}

func TestSubtract(t *testing.T) {
	if got := Subtract(5, 3); got != 2 {
		t.Errorf("Subtract(5, 3) = %v, want 2", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := Multiply(4, 3); got != 12 {
		t.Errorf("Multiply(4, 3) = %v, want 12", got)
	}
}

func TestDivide_Normal(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("Divide(10, 2) = %v, want 5", got)
	}
}

func TestDivide_ByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Fatal("expected an error when dividing by zero, got nil")
	}
	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("expected error to be recognizable via errors.Is(err, ErrDivisionByZero), got: %v", err)
	}
}
