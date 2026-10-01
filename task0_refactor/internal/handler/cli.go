// Package handler handles CLI input and output.
package handler

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"
)

var ErrUsage = errors.New("usage: app <a> <b> <text> <output.json>")

type CLI struct {
	service *service.ReportService
	output  io.Writer
}

func New(svc *service.ReportService, output io.Writer) *CLI {
	return &CLI{service: svc, output: output}
}

func (h *CLI) Run(args []string) error {
	if len(args) != 4 {
		return ErrUsage
	}
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return fmt.Errorf("parse a: %w", err)
	}
	b, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return fmt.Errorf("parse b: %w", err)
	}
	report, err := h.service.Create(a, b, args[2], args[3])
	if err != nil {
		return fmt.Errorf("run report command: %w", err)
	}
	if _, err := fmt.Fprintf(h.output, "quotient=%g words=%d characters=%d\n", report.Quotient, report.Words, report.Characters); err != nil {
		return fmt.Errorf("print saved report: %w", err)
	}
	return nil
}
