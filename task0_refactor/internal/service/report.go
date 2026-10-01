// Package service coordinates calculation, text analysis, and persistence.
package service

import (
	"fmt"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/calculator"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/models"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/repository"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/textanalyzer"
)

type ReportService struct {
	repository *repository.ReportRepository
}

func New(repo *repository.ReportRepository) *ReportService {
	return &ReportService{repository: repo}
}

func (s *ReportService) Create(a, b float64, text, path string) (models.Report, error) {
	quotient, err := calculator.Divide(a, b)
	if err != nil {
		return models.Report{}, fmt.Errorf("create report: %w", err)
	}
	words, err := textanalyzer.WordCount(text)
	if err != nil {
		return models.Report{}, fmt.Errorf("create report: %w", err)
	}
	report := models.Report{
		Quotient: quotient, Words: words, Characters: textanalyzer.CharCount(text),
	}
	if err := s.repository.Save(path, report); err != nil {
		return models.Report{}, fmt.Errorf("create report: %w", err)
	}
	return report, nil
}
