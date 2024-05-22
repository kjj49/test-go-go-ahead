package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/kjj49/test-go-go-ahead/internal/entity"
	"github.com/kjj49/test-go-go-ahead/pkg/postgres"
)

// CurrencyRepository -.
type CurrenycRepository struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *CurrenycRepository {
	return &CurrenycRepository{pg}
}

// SaveRequest -.
func (r *CurrenycRepository) SaveRequest(ctx context.Context, req entity.CurrencyRequest) error {
	date := time.Now()

	if len(req.Date) != 0 {
		parsedDate, err := time.Parse("02.01.2006", req.Date)
		if err != nil {
			return fmt.Errorf("CurrencyRepository - SaveRequest - time.Parse: %w", err)
		}
		date = parsedDate
	}

	query := "INSERT INTO currency_requests (val, date, request_date) VALUES ($1, $2, $3)"
	_, err := r.Pool.Exec(ctx, query, req.Val, date, time.Now())
	if err != nil {
		return fmt.Errorf("CurrencyRepository - SaveRequest - r.Pool.Exec: %w", err)
	}

	return nil
}

// SaveRequest -.
func (r *CurrenycRepository) SaveAllCurrency(ctx context.Context, allCurrency []entity.Currency) error {
	for _, c := range allCurrency {
		query := `
            INSERT INTO currency (val, nominal, value, updated_at) 
            VALUES ($1, $2, $3, $4) 
            ON CONFLICT (val) DO UPDATE 
            SET nominal = EXCLUDED.nominal, value = EXCLUDED.value, updated_at = EXCLUDED.updated_at
        `
		_, err := r.Pool.Exec(ctx, query, c.Val, c.Nominal, c.Value, c.Date)
		if err != nil {
			return fmt.Errorf("CurrencyRepository - SaveAllCurrency - r.Pool.Exec: %w", err)
		}
	}

	return nil
}
