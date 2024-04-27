package request

import (
	"errors"
	"slices"
	"strings"
)

var (
	ErrUnsupportedCurrency = errors.New("given currency is not supported")
	ErrEmptyCurrency       = errors.New("currency is required")

	SupportedCurrencies = []string{"USD", "GEL"}
)

type OpenBill struct {
	Currency string `json:"currency"`
}

func (req *OpenBill) Validate() error {
	if len(req.Currency) == 0 {
		return ErrEmptyCurrency
	}

	req.Currency = strings.ToUpper(req.Currency)

	if !slices.Contains(SupportedCurrencies, req.Currency) {
		return ErrUnsupportedCurrency
	}

	return nil
}
