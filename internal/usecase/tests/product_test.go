package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/savioafs/web-scraper-collector-go/internal/entity"
	"github.com/savioafs/web-scraper-collector-go/internal/usecase"
)

func TestCollect_Success(t *testing.T) {
	as := assert.New(t)
	mockUserRepo := new(MockUserRepo)
	mockProductRepo := new(MockProductRepo)
	mockWatchListRepo := new(MockWatchListRepo)

	productUseCase := usecase.NewProductUseCase(mockProductRepo, mockUserRepo, mockWatchListRepo)

	link := "https://www.mercadolivre.com.br/creatina-monohidratada-pura-1kg-dark-lab-unidade-sem-sabor/p/MLB25929487#reco_item_pos=1&reco_backend=item_decorator&reco_backend_type=function&reco_client=home_items-decorator-legacy&reco_id=4fcfeaed-c163-41fc-9508-71d4a782c3ca&reco_model=&c_id=/home/navigation-trends-recommendations/element&c_uid=96fadc5c-ab6d-444b-9474-02ead8771d16&da_id=navigation_trend&da_position=3&id_origin=/home/dynamic_access&da_sort_algorithm=ranker"

	mockUser, err := entity.NewUser("nameuser", "emailuser@gmail.com", "passsworduser")
	as.NoError(err)
	mockProduct, err := entity.NewProduct("creatina monohidratada", 132.79, link)
	as.NoError(err)

	mockUserRepo.On("Verify", mockUser.Email).Return(*mockUser, nil)
	mockProductRepo.On("Verify", link).Return(mockProduct, nil)
	mockWatchListRepo.On("Verify", mockProduct.ID, mockUser.ID).Return(false, nil)
	mockWatchListRepo.On("Register", mock.Anything).Return(nil)

	err = productUseCase.Collect(mockUser.Email, link)

	as.NoError(err)
	mockUserRepo.AssertExpectations(t)
	mockProductRepo.AssertExpectations(t)
	mockWatchListRepo.AssertExpectations(t)
}
