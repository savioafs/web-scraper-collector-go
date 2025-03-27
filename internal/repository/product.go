package repository

import (
	"database/sql"
	"github.com/savioafs/web-scraper-collector-go/internal/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Register(product *entity.Product) error {

	return nil
}
func (r *ProductRepository) Verify(link string) (entity.Product, error) {

	return entity.Product{}, nil
}
func (r *ProductRepository) Delete(id string) error {

	return nil
}
