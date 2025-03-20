package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type BookHistory struct {
	BookID string    `json:"book_id"`
	Date   time.Time `json:"date"`
	Price  float64   `json:"price"`
}

var (
	ErrTitleIsRequired  = errors.New("title is required")
	ErrPriceIsRequired  = errors.New("price is required")
	ErrRatingIsRequired = errors.New("rating is required")
	ErrStockIsRequired  = errors.New("stock is required")
)

type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	Rating      int       `json:"rating"`
	Stock       int       `json:"stock"`
	ExtractDate time.Time `json:"extract_date"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// todo: return response to value change
// todo: return response to stock low to (5 units)
// todo: active download in json and download report
// todo: get historic of price the books

func NewBook(title string, price float64, rating int, stock int) (*Book, error) {
	book := &Book{
		ID:          uuid.New().String(),
		Title:       title,
		Price:       price,
		Rating:      rating,
		Stock:       stock,
		ExtractDate: time.Now(),
	}

	if err := book.Validate(); err != nil {
		return nil, err
	}

	return book, nil
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return ErrTitleIsRequired
	}

	if b.Price == 0.0 {
		return ErrPriceIsRequired
	}

	if b.Rating == 0 {
		return ErrRatingIsRequired
	}

	if b.Stock == 0 {
		return ErrStockIsRequired
	}

	return nil
}
