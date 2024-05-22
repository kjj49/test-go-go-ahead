package usecase

import (
	"context"
	"fmt"

	"github.com/kjj49/test-go-go-ahead/internal/entity"
)

// CurrencyUseCase -.
type CurrencyUseCase struct {
	repository CurrencyRepository
	webAPI     CurrencyWebAPI
}

// New -.
func New(r CurrencyRepository, w CurrencyWebAPI) *CurrencyUseCase {
	return &CurrencyUseCase{
		repository: r,
		webAPI:     w,
	}
}

// GetCurrency - getting currency from sbr api.
func (uc *CurrencyUseCase) GetCurrency(ctx context.Context, c entity.CurrencyRequest) (entity.Currency, error) {
	currency, err := uc.webAPI.GetCurrency(c)
	if err != nil {
		return entity.Currency{}, fmt.Errorf("CurrencyUseCase - GetCurrency - s.webAPI.GetCurrency: %w", err)
	}

	err = uc.repository.SaveRequest(context.Background(), c)
	if err != nil {
		return entity.Currency{}, fmt.Errorf("CurrencyUseCase - GetCurrency - s.repository.SaveRequest: %w", err)
	}

	return currency, nil
}

// GetAllCurrency - getting all currency from sbr api.
func (uc *CurrencyUseCase) GetAllCurrency(ctx context.Context) error {
	allCurrency, err := uc.webAPI.GetAllCurrency()
	if err != nil {
		return fmt.Errorf("CurrencyUseCase - GetAllCurrency - s.webAPI.GetCurrency: %w", err)
	}

	err = uc.repository.SaveAllCurrency(context.Background(), allCurrency)
	if err != nil {
		return fmt.Errorf("CurrencyUseCase - GetCurrency - s.repository.SaveAllCurrency: %w", err)
	}

	return nil
}
