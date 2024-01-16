package core

import (
	"shared/utils"

	pb "shared/grpc"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

type tableRowName string

const (
	UPC               tableRowName = "UPC"
	PRICE_WITHOUT_TAX tableRowName = "Price (excl. tax)"
	TAX               tableRowName = "Tax"
	AVAILABILITY      tableRowName = "Availability"
)

var ALLOWED_DOMAIN = utils.GetEnvOrDefault("ALLOWED_DOMAIN", "")
var ASYNC_COUNT = utils.GetEnvAsIntOrDefault("ASYNC_COUNT", 5)
var SCRAPE_URL = utils.GetEnvOrDefault("SCRAPE_URL", "")

func Crawling(stream pb.Crawler_CrawlerDataIncomingClient) {

	c := colly.NewCollector(
		colly.AllowedDomains(ALLOWED_DOMAIN),
	)
	nested := c.Clone()
	c.Async = true
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: ASYNC_COUNT})

	c.OnRequest(func(r *colly.Request) {
		log.Info().Msgf("Visiting %v", r.URL)
	})

	c.OnHTML("article.product_pod", func(e *colly.HTMLElement) {
		url := e.ChildAttr("h3>a", "href")
		go nested.Visit(e.Request.AbsoluteURL(url))

	})

	c.OnHTML("ul.pager>li.next>a", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		if nextPage != "" {
			e.Request.Visit(nextPage)
		}
	})

	nested.OnRequest(func(r *colly.Request) {
		log.Info().Msgf("Accessing product page: %v", r.URL)
	})

	nested.OnHTML("article.product_page", func(h *colly.HTMLElement) {

		name := h.ChildText("div.row>div.product_main>h1")
		crawlerRequest := pb.CrawlerRequest{
			Name: &name,
		}

		h.ForEach("table.table>tbody>tr", func(_ int, e *colly.HTMLElement) {

			th := e.ChildText("th")
			field := e.ChildText("td")
			switch th {
			case string(UPC):
				crawlerRequest.Upc = &field
			case string(AVAILABILITY):
				crawlerRequest.Availability = &field
			case string(PRICE_WITHOUT_TAX):
				crawlerRequest.PriceWithoutTax = &field
			case string(TAX):
				crawlerRequest.Tax = &field
			}
		})
		err := stream.Send(&crawlerRequest)
		if err != nil {
			log.Err(err).Msg("failed to send stream to server")
			return
		}
	})

	c.Visit(SCRAPE_URL)
	c.Wait()
	nested.Wait()
}
