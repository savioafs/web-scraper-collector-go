package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrNameIsRequired   = errors.New("name is required")
	ErrPriceIsRequired  = errors.New("price is required")
	ErrRatingIsRequired = errors.New("rating is required")
	ErrStockIsRequired  = errors.New("stock is required")
	ErrURLIsRequired    = errors.New("url is required")
)

type ProductHistory struct {
	ProductID string    `json:"book_id"`
	Date      time.Time `json:"date"`
	Price     float64   `json:"price"`
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
// todo: return response to stock low to (5 units)
// todo: active download in json and download report
// todo: get historic of price the books
// todo: the user select email qnd send via login
// todo: verify link in db and add new client with link note: not return error, only atribute on other client

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
		return ErrNameIsRequired
	}

	if p.Price == 0.0 {
		return ErrPriceIsRequired
	}

	if p.Url == "" {
		return ErrURLIsRequired
	}

	return nil
}
