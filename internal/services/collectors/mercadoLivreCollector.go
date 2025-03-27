package collectors

import (
	"github.com/gocolly/colly/v2"
	"log"
	"strconv"
)

func Run(link string) (string, float64, error) {
	c := colly.NewCollector()

	var (
		name, priceText string
		price           float64
		err             error
	)

	c.OnHTML("div.ui-pdp-container__row.ui-pdp-with--separator--fluid.ui-pdp-with--separator--40-24", func(e *colly.HTMLElement) {
		name = e.ChildText("h1.ui-pdp-title")
		priceText = e.ChildAttr("div.ui-pdp-price__second-line meta[itemprop='price']", "content")

		price, err = strconv.ParseFloat(priceText, 64)
		if err != nil {
			log.Fatalf("cannot convert price %v", err)
		}
	})

	err = c.Visit(link)
	if err != nil {
		return "", 0.0, err
	}

	c.Wait()

	return name, price, nil
}
