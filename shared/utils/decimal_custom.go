package utils

import "github.com/shopspring/decimal"

type CustomDecimal struct {
	decimal.Decimal
}

// CustomDecimal method to present number with 2 decimals after the point when marshalling to JSON
func (cd *CustomDecimal) MarshalJSON() ([]byte, error) {
	return []byte(cd.StringFixed(2)), nil
}
