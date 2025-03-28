package entity

import (
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	type productInput struct {
		name  string
		price float64
		url   string
	}
	tests := []struct {
		name      string
		product   productInput
		expectErr error
	}{
		{
			name: "Create product with success",
			product: productInput{
				name:  "Caixa Caixinha De Som Potente Para Computador Subwoofer Potente Notebook Pc Usb 6w P2 Kaizu",
				price: 34.10,
				url:   "https://www.mercadolivre.com.br/caixa-caixinha-de-som-potente-para-computador-subwoofer-potente-notebook-pc-usb-6w-p2-kaizu/p/MLB46071639#polycard_client=recommendations_home_second-best-navigation-trend-recommendations&reco_backend=machinalis-homes-univb&wid=MLB3979944453&reco_client=home_second-best-navigation-trend-recommendations&reco_item_pos=3&reco_backend_type=function&reco_id=ad49502f-9dd9-4a0e-944c-bad5166456f5&sid=recos&c_id=/home/second-best-navigation-trend-recommendations/element&c_uid=d7bf204e-40a1-43ec-a3ac-62823d7665b0",
			},
			expectErr: nil,
		},
		{
			name: "Create product with error name is required",
			product: productInput{
				price: 34.10,
				url:   "https://www.mercadolivre.com.br/caixa-caixinha-de-som-potente-para-computador-subwoofer-potente-notebook-pc-usb-6w-p2-kaizu/p/MLB46071639#polycard_client=recommendations_home_second-best-navigation-trend-recommendations&reco_backend=machinalis-homes-univb&wid=MLB3979944453&reco_client=home_second-best-navigation-trend-recommendations&reco_item_pos=3&reco_backend_type=function&reco_id=ad49502f-9dd9-4a0e-944c-bad5166456f5&sid=recos&c_id=/home/second-best-navigation-trend-recommendations/element&c_uid=d7bf204e-40a1-43ec-a3ac-62823d7665b0",
			},
			expectErr: common.ErrNameIsRequired,
		},
		{
			name: "Create product with error price is required",
			product: productInput{
				name: "Caixa Caixinha De Som Potente Para Computador Subwoofer Potente Notebook Pc Usb 6w P2 Kaizu",
				url:  "https://www.mercadolivre.com.br/caixa-caixinha-de-som-potente-para-computador-subwoofer-potente-notebook-pc-usb-6w-p2-kaizu/p/MLB46071639#polycard_client=recommendations_home_second-best-navigation-trend-recommendations&reco_backend=machinalis-homes-univb&wid=MLB3979944453&reco_client=home_second-best-navigation-trend-recommendations&reco_item_pos=3&reco_backend_type=function&reco_id=ad49502f-9dd9-4a0e-944c-bad5166456f5&sid=recos&c_id=/home/second-best-navigation-trend-recommendations/element&c_uid=d7bf204e-40a1-43ec-a3ac-62823d7665b0",
			},
			expectErr: common.ErrPriceIsRequired,
		},
		{
			name: "Create product with err url is required",
			product: productInput{
				name:  "Caixa Caixinha De Som Potente Para Computador Subwoofer Potente Notebook Pc Usb 6w P2 Kaizu",
				price: 34.10,
			},
			expectErr: common.ErrURLIsRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := assert.New(t)

			product, err := NewProduct(tt.product.name, tt.product.price, tt.product.url)
			if err != nil {
				as.Equal(err, tt.expectErr)
				return
			}

			as.NotNil(product)
			as.Equal(product.Name, tt.product.name)
			as.Equal(product.Price, tt.product.price)
			as.Equal(product.Url, tt.product.url)
		})

	}
}
