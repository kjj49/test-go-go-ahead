package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/kjj49/test-go-go-ahead/internal/entity"
	"github.com/kjj49/test-go-go-ahead/internal/usecase"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func currency(t *testing.T) (*usecase.CurrencyUseCase, *MockCurrencyRepository, *MockCurrencyWebAPI) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockCurrencyRepository(mockCtl)
	webAPI := NewMockCurrencyWebAPI(mockCtl)

	currency := usecase.New(repo, webAPI)

	return currency, repo, webAPI
}

func TestGetCurrency(t *testing.T) {
	t.Parallel()

	currency, repo, webAPI := currency(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				webAPI.EXPECT().GetCurrency(entity.CurrencyRequest{}).Return(entity.Currency{}, nil)
				repo.EXPECT().SaveRequest(context.Background(), entity.CurrencyRequest{}).Return(nil)
			},
			res: entity.Currency{},
			err: nil,
		},
		{
			name: "web API error",
			mock: func() {
				webAPI.EXPECT().GetCurrency(entity.CurrencyRequest{}).Return(entity.Currency{}, errInternalServErr)
			},
			res: entity.Currency{},
			err: errInternalServErr,
		},
		{
			name: "repository error",
			mock: func() {
				webAPI.EXPECT().GetCurrency(entity.CurrencyRequest{}).Return(entity.Currency{}, nil)
				repo.EXPECT().SaveRequest(context.Background(), entity.CurrencyRequest{}).Return(errInternalServErr)
			},
			res: entity.Currency{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := currency.GetCurrency(context.Background(), entity.CurrencyRequest{})

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
