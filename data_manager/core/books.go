package core

import (
	"datamanager/database"
	"encoding/json"
	"regexp"
	pb "shared/grpc"
	"shared/utils"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

var reAvailability = regexp.MustCompile(`\d+`)
var rePrice = regexp.MustCompile(`\d+\.\d+`)

type BookData struct {
	Name            *string              `validate:"required"`
	Availability    *int                 `validate:"required"`
	Upc             *string              `validate:"required"`
	PriceWithoutTax *utils.CustomDecimal `validate:"required"`
	Tax             *utils.CustomDecimal `validate:"required"`
}

// bookData method to parse availability data
func (pd *BookData) parseAvailabilty(availability string) {
	digits := reAvailability.FindString(availability)
	digitsInt, err := strconv.Atoi(digits)
	if err != nil {
		log.Warn().Msgf("could not parse availaibility from '%v'", availability)
		return
	}
	pd.Availability = &digitsInt
}

// bookData method to parce price fields
func (pd *BookData) parcePrice(price string, isTax bool) {
	digits := rePrice.FindString(price)
	digitsDecimal, err := decimal.NewFromString(digits)
	if err != nil {
		log.Warn().Msgf("could not parse from %v", price)
		return
	}

	if isTax {
		pd.Tax = &utils.CustomDecimal{Decimal: digitsDecimal}
	} else {
		pd.PriceWithoutTax = &utils.CustomDecimal{Decimal: digitsDecimal}
	}
}

// parses and validates bookData and saves to db
func ParseAndValidate(in *pb.CrawlerRequest) error {

	bookData := BookData{
		Name: in.Name,
		Upc:  in.Upc,
	}
	bookData.parseAvailabilty(*in.Availability)
	bookData.parcePrice(*in.PriceWithoutTax, false)
	bookData.parcePrice(*in.Tax, true)

	err := utils.Validate[BookData](bookData)
	if err != nil {
		return err
	}

	err = database.SaveToDB[BookData](bookData, []byte(*bookData.Upc))
	return err
}

// fetches all books data from db
func GetBooks() ([]byte, error) {
	books := []BookData{}
	data, err := database.FetchAllFromDB()
	if err != nil {
		log.Err(err).Msg("failed to fetch from db")
		return []byte{}, err
	}
	for _, row := range data {
		book := BookData{}
		json.Unmarshal(row, &book)
		books = append(books, book)
	}

	response := struct {
		Data  []BookData `json:"data"`
		Total int        `json:"total"`
	}{
		Data:  books,
		Total: len(books),
	}

	jsonData, err := json.Marshal(response)
	return jsonData, err
}
