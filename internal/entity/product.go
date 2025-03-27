package entity

import (
	"github.com/google/uuid"
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"time"
)

type WatchList struct {
	ProductID string    `json:"book_id"`
	ClientID  time.Time `json:"client_id"`
}

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Url         string    `json:"url"`
	ExtractDate time.Time `json:"extract_date"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// todo: return response to value change
// todo: active download in json and download report
// todo: get historic of price the books
// todo: the user select email qnd send via login
// todo: verify link in db and add new client with link note: not return error, only atribute on other client
// your name, your email, link product
// register your email
// register your product

func NewProduct(name string, price float64, url string) (*Product, error) {
	p := &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Price:       price,
		Url:         url,
		ExtractDate: time.Now(),
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return common.ErrNameIsRequired
	}

	if p.Price == 0.0 {
		return common.ErrPriceIsRequired
	}

	if p.Url == "" {
		return common.ErrURLIsRequired
	}

	return nil
}
