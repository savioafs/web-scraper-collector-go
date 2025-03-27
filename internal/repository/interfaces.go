package repository

import "github.com/savioafs/web-scraper-collector-go/internal/entity"

type UserStorer interface {
	Register(user *entity.User) error
	Verify(email string) (entity.User, error)
	//Delete(email string) error
}

type ProductStorer interface {
	Register(product *entity.Product) error
	Verify(link string) (*entity.Product, error)
	//Delete(id string) error // for control of system
}

type ProductHistoryStorer interface {
	Register(productHistory *entity.ProductHistory) error
	Delete(id string) error // for control of system
}

type WatchListStorer interface {
	Register(watcherList *entity.WatchList) error
	Verify(productID, userID string) (bool, error)
}
