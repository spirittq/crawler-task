package core

import (
	"errors"
	"shared/grpc"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestBookDataParseAvailability(t *testing.T) {

	t.Run("availability successfully parsed", func(t *testing.T) {
		availability := "in store (22 available)"
		bookData := BookData{}
		expectedResult := 22

		bookData.parseAvailabilty(availability)

		assert.Equal(t, expectedResult, *bookData.Availability)
	})

	t.Run("availability could not be parsed", func(t *testing.T) {
		availability := "some random text"
		bookData := BookData{}
		bookData.parseAvailabilty(availability)
		assert.Nil(t, bookData.Availability)
	})
}

func TestBookDataParsePrice(t *testing.T) {

	t.Run("price & tax parsed succesfully", func(t *testing.T) {

		price := "$22.22"
		tax := "$0.11"

		expectedPrice := decimal.NewFromFloat(22.22)
		expectedTax := decimal.NewFromFloat(0.11)

		bookData := BookData{}
		bookData.parcePrice(price, false)
		bookData.parcePrice(tax, true)

		assert.True(t, expectedPrice.Equal(bookData.PriceWithoutTax.Decimal))
		assert.True(t, expectedTax.Equal(bookData.Tax.Decimal))
	})

	t.Run("price is not parsed", func(t *testing.T) {

		price := "not a number"
		bookData := BookData{}
		bookData.parcePrice(price, false)
		assert.Nil(t, bookData.PriceWithoutTax)
	})
}

func mockSaveToDBNoError(pd *BookData) error {
	return nil
}

func mockSaveToDBError(pd *BookData) error {
	return errors.New("error")
}

func TestParseAndValidate(t *testing.T) {

	name := "t"
	upc := "upc"
	availability := "22"
	price := "23.22"
	tax := "24.22"

	mockBookData := &grpc.CrawlerRequest{
		Name:            &name,
		Upc:             &upc,
		Availability:    &availability,
		PriceWithoutTax: &price,
		Tax:             &tax,
	}

	t.Run("book data save to db", func(t *testing.T) {
		val = BookDataSave{
			SaveToDB: mockSaveToDBNoError,
		}

		err := ParseAndValidate(mockBookData)
		assert.Nil(t, err)
	})

	t.Run("book data returns error from db", func(t *testing.T) {

		val = BookDataSave{
			SaveToDB: mockSaveToDBError,
		}

		err := ParseAndValidate(mockBookData)
		assert.Error(t, err)

	})

}
