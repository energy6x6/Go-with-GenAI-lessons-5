package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDo_SucceedsOnFirstAttempt(t *testing.T) {
	op := NewFlakyOperation(0, "ok")

	result, err := Do(op, 3, time.Millisecond)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if result != "ok" {
		t.Errorf("expected result %q, got %q", "ok", result)
	}
}

func TestDo_SucceedsAfterRetries(t *testing.T) {
	op := NewFlakyOperation(2, "ok")

	result, err := Do(op, 5, time.Millisecond)
	if err != nil {
		t.Fatalf("expected eventual success, got error: %v", err)
	}
	if result != "ok" {
		t.Errorf("expected result %q, got %q", "ok", result)
	}
}

func TestDo_ExhaustsAttemptsAndReturnsWrappedError(t *testing.T) {
	op := NewFlakyOperation(100, "ok") // always fails within the attempt budget

	_, err := Do(op, 3, time.Millisecond)
	if err == nil {
		t.Fatal("expected an error after exhausting all attempts, got nil")
	}
	if !errors.Is(err, ErrTemporary) {
		t.Errorf("expected the final error to still be recognizable via errors.Is(err, ErrTemporary), got: %v", err)
	}
}

func TestDo_RespectsMaxAttempts(t *testing.T) {
	calls := 0
	op := func() (string, error) {
		calls++
		return "", fmt.Errorf("always fails: %w", ErrTemporary)
	}

	_, err := Do(op, 4, time.Millisecond)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	if calls != 4 {
		t.Errorf("expected exactly 4 attempts, got %d", calls)
	}
}

func TestDo_DoesNotRetryNonTemporaryErrors(t *testing.T) {
	permanentErr := errors.New("permanent failure")
	calls := 0
	op := func() (string, error) {
		calls++
		return "", permanentErr
	}

	_, err := Do(op, 5, time.Millisecond)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	if calls != 1 {
		t.Errorf("expected exactly 1 attempt for a non-temporary error, got %d", calls)
	}
	if errors.Is(err, ErrTemporary) {
		t.Error("did not expect a permanent error to be recognized as ErrTemporary")
	}
}

func TestDo_WaitsBackoffBetweenAttempts(t *testing.T) {
	op := NewFlakyOperation(2, "ok")
	backoff := 20 * time.Millisecond

	start := time.Now()
	_, err := Do(op, 5, backoff)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("expected eventual success, got error: %v", err)
	}
	// 2 retries => at least 2 backoff pauses.
	minExpected := 2 * backoff
	if elapsed < minExpected {
		t.Errorf("expected Do to wait at least %v between retries, only took %v", minExpected, elapsed)
	}
}
