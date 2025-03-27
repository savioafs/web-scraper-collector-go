package usecase

import (
	"errors"
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"github.com/savioafs/web-scraper-collector-go/internal/entity"
	"github.com/savioafs/web-scraper-collector-go/internal/repository"
	"github.com/savioafs/web-scraper-collector-go/internal/services/collectors"
)

type ProductUseCase struct {
	productRepo   repository.ProductStorer
	userRepo      repository.UserStorer
	watchListRepo repository.WatchListStorer
}

func NewProductUseCase(productRepo repository.ProductStorer, userRepo repository.UserStorer, watchListRepo repository.WatchListStorer) *ProductUseCase {
	return &ProductUseCase{
		productRepo:   productRepo,
		userRepo:      userRepo,
		watchListRepo: watchListRepo,
	}
}

func (u *ProductUseCase) Collect(emailUser, linkProduct string) error {

	user, err := u.userRepo.Verify(emailUser)
	if err != nil {
		return err
	}

	product, err := u.productRepo.Verify(linkProduct)
	if err != nil && !errors.Is(err, common.ErrProductNotFound) {
		return err
	}

	existsWatcherListByUser, err := u.watchListRepo.Verify(product.ID, user.ID)
	if err != nil {
		return err
	}
	if existsWatcherListByUser {
		return common.ErrProductAlreadyWatcherByUser
	}

	if product.ID == "" {
		product, err = u.registerNewProduct(linkProduct)
		if err != nil {
			return err
		}
	}

	watchList, err := entity.NewWatchList(product.ID, user.ID)
	if err != nil {
		return err
	}

	err = u.watchListRepo.Register(watchList)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductUseCase) registerNewProduct(linkProduct string) (*entity.Product, error) {

	name, price, err := collectors.Run(linkProduct)
	if err != nil {
		return nil, errors.New("cannot connect collector")
	}

	product, err := entity.NewProduct(name, price, linkProduct)
	if err != nil {
		return nil, err
	}

	err = u.productRepo.Register(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
