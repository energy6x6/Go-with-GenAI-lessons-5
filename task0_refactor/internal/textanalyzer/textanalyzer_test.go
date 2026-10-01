package textanalyzer

import (
	"errors"
	"testing"
)

func TestWordCount_Normal(t *testing.T) {
	got, err := WordCount("the quick brown fox")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 4 {
		t.Errorf("WordCount(...) = %d, want 4", got)
	}
}

func TestWordCount_Empty(t *testing.T) {
	_, err := WordCount("   ")
	if err == nil {
		t.Fatal("expected an error for empty text, got nil")
	}
	if !errors.Is(err, ErrEmptyText) {
		t.Errorf("expected error to be recognizable via errors.Is(err, ErrEmptyText), got: %v", err)
	}
}

func TestCharCount(t *testing.T) {
	got := CharCount("  hello  ")
	if got != 5 {
		t.Errorf("CharCount(...) = %d, want 5", got)
	}
}
