package core

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

func Crawling() {

	c := colly.NewCollector()
	nested := c.Clone()

	nested.OnHTML("article", func(h *colly.HTMLElement) {
		fmt.Println("here")
		titleName := h.ChildText("h1")
		fmt.Println(titleName)
	})

	nested.OnRequest(func(r *colly.Request) {
		log.Info().Msgf("Accessing product page: %v", r.URL)
	})

	c.OnHTML("article.product_pod", func(e *colly.HTMLElement) {

		url := e.ChildAttr("h3>a", "href")
		nested.Visit(e.Request.AbsoluteURL(url))

	})

	c.OnRequest(func(r *colly.Request) {
		log.Info().Msgf("Visiting %v", r.URL)
	})

	c.Visit("https://books.toscrape.com/")
}
