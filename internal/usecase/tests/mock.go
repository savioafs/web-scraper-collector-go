package tests

import (
	"github.com/stretchr/testify/mock"

	"github.com/savioafs/web-scraper-collector-go/internal/entity"
)

// User Storer
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Register(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) Verify(email string) (entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(entity.User), args.Error(1)
}

// Product Storer
type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) Register(product *entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepo) Verify(link string) (*entity.Product, error) {
	args := m.Called(link)
	return args.Get(0).(*entity.Product), args.Error(1)
}

// Watch list Storer
type MockWatchListRepo struct {
	mock.Mock
}

func (m *MockWatchListRepo) Register(watcherList *entity.WatchList) error {
	args := m.Called(watcherList)
	return args.Error(0)
}

func (m *MockWatchListRepo) Verify(productID, userID string) (bool, error) {
	args := m.Called(productID, userID)
	return args.Bool(0), args.Error(1)
}
