package api_test

import (
	"datamanager/core"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shared/utils"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBooks(t *testing.T) {

	t.Run("data is returned if call is successfull", func(t *testing.T) {

		testName := "testName"
		testAvailability := 5
		testUPC := "testUpc"
		testPriceWithoutTax := utils.CustomDecimal{Decimal: decimal.NewFromFloat(12.34)}
		testTax := utils.CustomDecimal{Decimal: decimal.NewFromFloat(56.78)}

		bookData := []core.BookData{
			{
				Name:            &testName,
				Availability:    &testAvailability,
				Upc:             &testUPC,
				PriceWithoutTax: &testPriceWithoutTax,
				Tax:             &testTax,
			},
		}

		response := struct {
			Data  []core.BookData `json:"data"`
			Total int             `json:"total"`
		}{
			Data:  bookData,
			Total: len(bookData),
		}

		jsonData, _ := json.Marshal(response)

		core.GetBooks = func() ([]byte, error) {

			return jsonData, nil
		}

		resp, _ := http.Get("http://localhost:3000/books")
		require.Equal(t, http.StatusOK, resp.StatusCode)
		bodyBytes, err := io.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, bodyBytes, jsonData)
	})

	t.Run("500 is returned if core returns error", func(t *testing.T) {

		core.GetBooks = func() ([]byte, error) {
			return []byte{}, errors.New("test")
		}

		resp, _ := http.Get("http://localhost:3000/books")
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}
