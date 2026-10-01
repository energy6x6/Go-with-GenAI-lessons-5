// Package repository stores reports as JSON files.
package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/models"
)

type ReportRepository struct{}

// Save replaces the file at path after successfully encoding the report.
func (r *ReportRepository) Save(path string, report models.Report) error {
	data, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("encode report: %w", err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("write report %q: %w", path, err)
	}
	return nil
}
