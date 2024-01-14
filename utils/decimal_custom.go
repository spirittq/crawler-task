package utils

import "github.com/shopspring/decimal"

type CustomDecimal struct {
	decimal.Decimal
}

func (cd *CustomDecimal) MarshalJSON() ([]byte, error) {
	return []byte(cd.StringFixed(2)), nil
}
