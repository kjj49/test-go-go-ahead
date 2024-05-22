package validator

import (
	"strings"
	"time"
)

var validCurrencies = map[string]bool{
	"USD": true, "EUR": true, "GBP": true, "AUD": true, "AZN": true, "BGN": true, "CAD": true,
	"BRL": true, "HUF": true, "VND": true, "HKD": true, "GEL": true, "DKK": true, "AED": true,
	"INR": true, "IDR": true, "KZT": true, "BYN": true, "QAR": true, "KGS": true, "CNY": true,
	"MDL": true, "NZD": true, "NOK": true, "PLN": true, "RON": true, "XDR": true, "SGD": true,
	"TJS": true, "THB": true, "TRY": true, "TMT": true, "UZS": true, "UAH": true, "CZK": true,
	"SEK": true, "CHF": true, "RSD": true, "ZAR": true, "KRW": true, "JPY": true,
}

func ValidateCurrencyCode(code string) bool {
	return validCurrencies[strings.ToUpper(code)]
}

func ValidateDate(date string) bool {
	_, err := time.Parse("02.01.2006", date)
	return err == nil
}
