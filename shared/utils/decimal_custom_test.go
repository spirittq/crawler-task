package utils

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCustomDecimalMarshalJSON(t *testing.T) {

	expectedValue := []byte("20.20")
	customDecimal := CustomDecimal{Decimal: decimal.NewFromFloat(20.20)}
	actualValue, err := customDecimal.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, expectedValue, actualValue)

}
