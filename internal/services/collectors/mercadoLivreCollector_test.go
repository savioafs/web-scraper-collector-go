package collectors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct {
		name         string
		productName  string
		productPrice float64
		link         string
		expectErr    error
	}{
		{
			name:         "Get product caixinha de som kaizu",
			productName:  "Caixa Caixinha De Som Potente Para Computador Subwoofer Potente Notebook Pc Usb 6w P2 Kaizu",
			productPrice: 34.10,
			link:         "https://www.mercadolivre.com.br/caixa-caixinha-de-som-potente-para-computador-subwoofer-potente-notebook-pc-usb-6w-p2-kaizu/p/MLB46071639#polycard_client=recommendations_home_second-best-navigation-trend-recommendations&reco_backend=machinalis-homes-univb&wid=MLB3979944453&reco_client=home_second-best-navigation-trend-recommendations&reco_item_pos=3&reco_backend_type=function&reco_id=ad49502f-9dd9-4a0e-944c-bad5166456f5&sid=recos&c_id=/home/second-best-navigation-trend-recommendations/element&c_uid=d7bf204e-40a1-43ec-a3ac-62823d7665b0",
			expectErr:    nil,
		},
		{
			name:         "Get product creatina dark lab",
			productName:  "Creatina Monohidratada Pura 1kg Dark Lab Unidade Sem sabor",
			productPrice: 132.70,
			link:         "https://www.mercadolivre.com.br/creatina-monohidratada-pura-1kg-dark-lab-unidade-sem-sabor/p/MLB25929487?pdp_filters=item_id%3AMLB4812130742#polycard_client=offers&deal_print_id=7c33365d-6cf9-4bef-a0ea-e7200ac41e77&position=1&tracking_id=&wid=MLB4812130742&sid=offers",
			expectErr:    nil,
		},
		{
			name:         "Get product with error not found",
			productName:  "",
			productPrice: 0.0,
			link:         "https://mercadolivre.com.br/MLB-3604451061-coqueteleira-2-doses-com-mola-500ml-duo-pack-rythmoon-_JM#polycard_client=recommendations_pdp-p2p&reco_backend=ranker_compl&reco_model=retrieval-ranker-complementarios&reco_client=pdp-p2p&reco_item_pos=3&reco_backend_type=low_level&reco_id=77e65252-fb17-4a98-a4e7-ea862a52395b",
			expectErr:    errors.New("Not Found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := assert.New(t)

			name, price, err := Run(tt.link)
			if err != nil {
				as.Equal(err, tt.expectErr)
				return
			}
			as.Equal(name, tt.productName)
			as.Equal(price, tt.productPrice)
		})
	}
}
