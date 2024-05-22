// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"time"
)

// Currency
type Currency struct {
	Val     string    `json:"val"       example:"USD"`
	Nominal int       `json:"nominal"       example:"1"`
	Value   float64   `json:"value"  example:"90.22"`
	Date    time.Time `json:"date"  example:"2024-05-22T00:00:00Z"`
}

// CurrencyRequest
type CurrencyRequest struct {
	Val  string `json:"val"  binding:"required"  example:"USD"`
	Date string `json:"date"      example:"22.05.2024"`
}
