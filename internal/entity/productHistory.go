package entity

import (
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"time"
)

type ProductHistory struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	Price     float64   `json:"price"`
	Date      time.Time `json:"date"`
}

func SaveProductPriceHistory(productID string, newPrice float64) error {

	p := &ProductHistory{
		ProductID: productID,
		Price:     newPrice,
		Date:      time.Now(),
	}

	if err := p.Validate(); err != nil {
		return err
	}

	return nil
}

func (p *ProductHistory) Validate() error {
	if p.ProductID == "" {
		return common.ErrProductIDIsRequired
	}

	if p.Price == 0.0 {
		return common.ErrPriceIsRequired
	}

	return nil
}
