package webapi

import (
	"encoding/xml"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kjj49/test-go-go-ahead/internal/entity"
	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

// CurrencyWebAPI -.
type CurrencyWebAPI struct {
	UserAgent  []string
	ServiceUrl []string
}

// New -.
func New() *CurrencyWebAPI {
	return &CurrencyWebAPI{
		UserAgent:  []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrl: []string{"https://cbr.ru/scripts/XML_daily.asp?"},
	}
}

// GetCurrency -.
func (c *CurrencyWebAPI) GetCurrency(request entity.CurrencyRequest) (entity.Currency, error) {
	url := c.ServiceUrl[0]

	switch {
	case len(request.Date) != 0:
		url += "date_req=" + request.Date
	default:
		// If there is no date, we just use the base URL
	}

	resp, err := c.makeRequest(url)
	if err != nil {
		return entity.Currency{}, errors.Wrap(err, "CurrencyWebAPI - GetCurrency - makeRequest")
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var CBRResponse entity.CBRResponse
	err = decoder.Decode(&CBRResponse)
	if err != nil {
		return entity.Currency{}, errors.Wrap(err, "CurrencyWebAPI - GetCurrency - xml.Decode")
	}

	// Find the required currency in XML and fill the entity.Currency
	for _, valute := range CBRResponse.Valutes {
		if strings.EqualFold(strings.ToUpper(valute.CharCode), strings.ToUpper(request.Val)) {
			value, err := parseValue(valute.Value)
			if err != nil {
				return entity.Currency{}, errors.Wrap(err, "CurrencyWebAPI - GetCurrency - parseValue")
			}

			date, err := time.Parse("02.01.2006", CBRResponse.Date)
			if err != nil {
				return entity.Currency{}, errors.Wrap(err, "CurrencyWebAPI - GetCurrency - time.Parse")
			}

			return entity.Currency{
				Val:     valute.CharCode,
				Nominal: valute.Nominal,
				Value:   value,
				Date:    date,
			}, nil
		}
	}

	return entity.Currency{}, errors.New("Currency not found in XML")
}

// GetAllCurrency -.
func (c *CurrencyWebAPI) GetAllCurrency() ([]entity.Currency, error) {
	url := c.ServiceUrl[0]
	allCurrency := []entity.Currency{}

	resp, err := c.makeRequest(url)
	if err != nil {
		return []entity.Currency{}, errors.Wrap(err, "CurrencyWebAPI - GetAllCurrency - makeRequest")
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var CBRResponse entity.CBRResponse
	if err := decoder.Decode(&CBRResponse); err != nil {
		return []entity.Currency{}, errors.Wrap(err, "CurrencyWebAPI - GetAllCurrency - xml.Decode")
	}

	date, err := time.Parse("02.01.2006", CBRResponse.Date)
	if err != nil {
		return []entity.Currency{}, errors.Wrap(err, "CurrencyWebAPI - GetAllCurrency - time.Parse")
	}

	// Get all currency in XML and fill []entity.Currency
	for _, valute := range CBRResponse.Valutes {
		value, err := parseValue(valute.Value)
		if err != nil {
			return []entity.Currency{}, errors.Wrap(err, "CurrencyWebAPI - GetAllCurrency - parseValue")
		}
		allCurrency = append(allCurrency, entity.Currency{
			Val:     valute.CharCode,
			Nominal: valute.Nominal,
			Value:   value,
			Date:    date,
		})
	}

	return allCurrency, nil
}

func (c *CurrencyWebAPI) makeRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent[0])

	return client.Do(req)
}

func parseValue(value string) (float64, error) {
	value = strings.ReplaceAll(value, ",", ".")
	return strconv.ParseFloat(value, 64)
}
