package core

import (
	"crawler-task/database"
	"crawler-task/utils"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

var reAvailability = regexp.MustCompile(`\d+`)
var rePrice = regexp.MustCompile(`\d+\.\d+`)

type ProductData struct {
	Name            string           `validate:"required"`
	Availability    *int             `validate:"required"`
	UPC             string           `validate:"required"`
	PriceWithoutTax *decimal.Decimal `validate:"required"`
	Tax             *decimal.Decimal `validate:"required"`
}

func (pd *ProductData) parseAvailabilty(availability string) {
	digits := reAvailability.FindString(availability)
	digitsInt, err := strconv.Atoi(digits)
	if err != nil {
		log.Warn().Msgf("could not parse availaibility from '%v'", availability)
		return
	}
	pd.Availability = &digitsInt
}

func (pd *ProductData) parcePrice(price string, isTax bool) {
	digits := rePrice.FindString(price)
	digitsDecimal, err := decimal.NewFromString(digits)
	if err != nil {
		log.Warn().Msgf("could not parse from %v", price)
		return
	}

	if isTax {
		pd.Tax = &digitsDecimal
	} else {
		pd.PriceWithoutTax = &digitsDecimal
	}
}

type tableRowName string

const (
	upc             tableRowName = "UPC"
	priceWithoutTax tableRowName = "Price (excl. tax)"
	tax             tableRowName = "Tax"
	availability    tableRowName = "Availability"
)

func Crawling() {

	c := colly.NewCollector(
		colly.AllowedDomains(utils.GetEnvOrDefault("ALLOWED_DOMAIN", "")),
	)
	nested := c.Clone()

	nested.OnHTML("article.product_page", func(h *colly.HTMLElement) {
		// TODO: should be goroutined
		productData, err := parseAndValidateData(h)
		if err != nil {
			log.Err(err).Msgf("unable to parse product data")
		}
		database.SaveToDB[*ProductData](&productData)
	})

	nested.OnRequest(func(r *colly.Request) {
		log.Info().Msgf("Accessing product page: %v", r.URL)
	})

	c.OnHTML("article.product_pod", func(e *colly.HTMLElement) {

		url := e.ChildAttr("h3>a", "href")
		nested.Visit(e.Request.AbsoluteURL(url))

	})

	c.OnHTML("ul.pager>li.next>a", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		if nextPage != "" {
			e.Request.Visit(nextPage)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Info().Msgf("Visiting %v", r.URL)
	})

	c.Visit(utils.GetEnvOrDefault("SCRAPE_URL", ""))
}

func parseAndValidateData(h *colly.HTMLElement) (ProductData, error) {

	productData := ProductData{
		Name: h.ChildText("div.row>div.product_main>h1"),
	}
	h.ForEach("table.table>tbody>tr", func(_ int, e *colly.HTMLElement) {

		th := e.ChildText("th")
		switch th {
		case string(upc):
			productData.UPC = e.ChildText("td")
		case string(availability):
			productData.parseAvailabilty(e.ChildText("td"))
		case string(priceWithoutTax):
			productData.parcePrice(e.ChildText("td"), false)
		case string(tax):
			productData.parcePrice(e.ChildText("td"), true)
		}
	})

	err := utils.Validate[ProductData](productData)
	return productData, err
}
