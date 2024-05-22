// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/kjj49/test-go-go-ahead/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Currency -.
	Currency interface {
		GetCurrency(context.Context, entity.CurrencyRequest) (entity.Currency, error)
	}

	// CurrencyRepository -.
	CurrencyRepository interface {
		SaveRequest(context.Context, entity.CurrencyRequest) error
		SaveAllCurrency(context.Context, []entity.Currency) error
	}

	// CurrencyWebAPI -.
	CurrencyWebAPI interface {
		GetCurrency(entity.CurrencyRequest) (entity.Currency, error)
		GetAllCurrency() ([]entity.Currency, error)
	}
)
