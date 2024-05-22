package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/kjj49/test-go-go-ahead/internal/entity"
	"github.com/kjj49/test-go-go-ahead/internal/usecase"
	"github.com/kjj49/test-go-go-ahead/internal/validator"
	"github.com/kjj49/test-go-go-ahead/pkg/logger"
)

type currencyRoutes struct {
	t usecase.Currency
	l logger.Interface
}

func newCurrencyRoutes(handler *gin.RouterGroup, t usecase.Currency, l logger.Interface) {
	r := &currencyRoutes{t, l}

	h := handler.Group("")
	{
		h.GET("", r.getCurrencyByParams)
	}
}

// @Summary Get currency
// @Tags currency
// @Description Get currency by params
// @ID get-currency-by-params
// @Accept  json
// @Produce  json
// @Param input query entity.CurrencyRequest true "request params"
// @Success 200 {object} entity.Currency
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /currency [get]
func (r *currencyRoutes) getCurrencyByParams(c *gin.Context) {
	val := c.Query("val")
	date := c.Query("date")

	// Validation for the presence of the val parameter
	if val == "" {
		newErrorResponse(c, http.StatusBadRequest, "Val parameter is required. For example USD, EUR, AUD, etc. The case does not matter")
		return
	}

	// Remove spaces from the val variable and convert to uppercase
	val = strings.ToUpper(strings.TrimSpace(val))

	// Check if val contains exactly three characters
	if len(val) != 3 {
		newErrorResponse(c, http.StatusBadRequest, "The currency code must consist of three characters. For example USD, EUR, AUD, etc. The case does not matter")
		return
	}

	// Checking the correctness of the currency code
	if !validator.ValidateCurrencyCode(val) {
		newErrorResponse(c, http.StatusBadRequest, "Unsupported currency code, try another one. For example USD, EUR, AUD, etc. The case does not matter")
		return
	}

	// Validation of the date parameter, if it has been passed
	if date != "" && !validator.ValidateDate(date) {
		newErrorResponse(c, http.StatusBadRequest, "Incorrect date format. Use the format 22.05.2024")
		return
	}

	request := entity.CurrencyRequest{
		Val:  val,
		Date: date,
	}

	currency, err := r.t.GetCurrency(c.Request.Context(), request)
	if err != nil {
		r.l.Error(err, "http - getCurrencyByParams")
		newErrorResponse(c, http.StatusInternalServerError, "cbr.ru API problems")
		return
	}

	c.JSON(http.StatusOK, currency)
}
