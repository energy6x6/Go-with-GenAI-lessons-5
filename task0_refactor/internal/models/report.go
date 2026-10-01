// Package models contains the report shared by the application layers.
package models

type Report struct {
	Quotient   float64 `json:"quotient"`
	Words      int     `json:"words"`
	Characters int     `json:"characters"`
}
