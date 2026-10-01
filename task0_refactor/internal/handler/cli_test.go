package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/calculator"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/handler"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/models"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/repository"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/textanalyzer"
)

func newCLI(out io.Writer) *handler.CLI {
	return handler.New(service.New(&repository.ReportRepository{}), out)
}

func TestRunSavesReport(t *testing.T) {
	path := filepath.Join(t.TempDir(), "report.json")
	var out bytes.Buffer
	if err := newCLI(&out).Run([]string{"10", "2", "  Привіт Go  ", path}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var report models.Report
	if err := json.Unmarshal(data, &report); err != nil {
		t.Fatal(err)
	}
	if want := (models.Report{Quotient: 5, Words: 2, Characters: 9}); report != want {
		t.Fatalf("report = %+v, want %+v", report, want)
	}
	if got := out.String(); got != "quotient=5 words=2 characters=9\n" {
		t.Fatalf("unexpected output: %q", got)
	}
}

func TestRunRejectsInvalidInputWithoutOverwritingReport(t *testing.T) {
	cases := []struct {
		name  string
		args  []string
		cause error
	}{
		{"usage", nil, handler.ErrUsage},
		{"invalid a", []string{"oops", "2", "hello"}, strconv.ErrSyntax},
		{"invalid b", []string{"10", "oops", "hello"}, strconv.ErrSyntax},
		{"division by zero", []string{"10", "0", "hello"}, calculator.ErrDivisionByZero},
		{"empty text", []string{"10", "2", " \t"}, textanalyzer.ErrEmptyText},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "report.json")
			if err := os.WriteFile(path, []byte("keep"), 0600); err != nil {
				t.Fatal(err)
			}
			args := append(append([]string{}, tc.args...), path)
			var out bytes.Buffer
			err := newCLI(&out).Run(args)
			if !errors.Is(err, tc.cause) {
				t.Fatalf("error = %v, want cause %v", err, tc.cause)
			}
			if errors.Is(err, strconv.ErrSyntax) {
				var numErr *strconv.NumError
				if !errors.As(err, &numErr) {
					t.Fatal("parse error type was lost")
				}
			}
			data, readErr := os.ReadFile(path)
			if readErr != nil || string(data) != "keep" || out.Len() != 0 {
				t.Fatalf("invalid input changed file or printed success: %q, %v, %q", data, readErr, out.String())
			}
		})
	}
}

func TestRunPreservesStorageError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing", "report.json")
	err := newCLI(io.Discard).Run([]string{"10", "2", "hello", path})
	var pathErr *os.PathError
	if !errors.Is(err, fs.ErrNotExist) || !errors.As(err, &pathErr) {
		t.Fatalf("storage error chain lost: %v", err)
	}
}

func TestRunPreservesEncodingErrorWithoutOverwritingReport(t *testing.T) {
	path := filepath.Join(t.TempDir(), "report.json")
	if err := os.WriteFile(path, []byte("keep"), 0600); err != nil {
		t.Fatal(err)
	}
	err := newCLI(io.Discard).Run([]string{"NaN", "2", "hello", path})
	var valueErr *json.UnsupportedValueError
	if !errors.As(err, &valueErr) {
		t.Fatalf("encoding error type lost: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "keep" {
		t.Fatalf("encoding failure changed file: %q, %v", data, err)
	}
}

type brokenWriter struct{}

func (brokenWriter) Write([]byte) (int, error) { return 0, io.ErrClosedPipe }

func TestRunReportsOutputFailureAfterSaving(t *testing.T) {
	path := filepath.Join(t.TempDir(), "report.json")
	err := newCLI(brokenWriter{}).Run([]string{"10", "2", "hello", path})
	if !errors.Is(err, io.ErrClosedPipe) {
		t.Fatalf("output error lost: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("report should already be saved: %v", err)
	}
}
