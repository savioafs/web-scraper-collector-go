package collectors

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/savioafs/web-scraper-collector-go/internal/entity"
	"log"
)

//type Collector struct {
//	repo repository.CollectorStorer
//}
//
//func NewCollector(repo repository.CollectorStorer) *Collector {
//	return &Collector{
//		repo: repo,
//	}
//}

func Run(link string) error {
	cl := colly.NewCollector()

	var (
		name  string
		price float64
	)

	cl.OnHTML("div.ui-pdp-container__row.ui-pdp-with--separator--fluid.ui-pdp-with--separator--40-24", func(e *colly.HTMLElement) {
		name := e.ChildText("h1.ui-pdp-title")
		price := e.ChildAttr("div.ui-pdp-price__second-line meta[itemprop='price']", "content")

		fmt.Println("name on child", name)
		fmt.Println("price on child", price)
	})

	// debug
	cl.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting:", r.URL.String())
	})

	cl.OnScraped(func(r *colly.Response) {
		p, err := entity.NewProduct(name, price, link)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("name:", p.Name)
		fmt.Println("price:", p.Price)

		fmt.Println("link:", p.Url)
	})

	err := cl.Visit(link)
	if err != nil {
		log.Fatal(err)
	}

	cl.Wait()

	return nil
}
