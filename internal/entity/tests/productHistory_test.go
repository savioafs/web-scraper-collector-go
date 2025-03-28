package tests

import (
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"github.com/savioafs/web-scraper-collector-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveProductPriceHistory(t *testing.T) {
	type productHistoryInput struct {
		productID string
		newPrice  float64
	}

	tests := []struct {
		name           string
		productHistory productHistoryInput
		expectErr      error
	}{
		{
			name: "Create product list with success",
			productHistory: productHistoryInput{
				productID: "idone",
				newPrice:  390.99,
			},
			expectErr: nil,
		},
		{
			name: "Create product list with error product id is required",
			productHistory: productHistoryInput{
				newPrice: 390.99,
			},
			expectErr: common.ErrProductIDIsRequired,
		},
		{
			name: "Create product list with error price is required",
			productHistory: productHistoryInput{
				productID: "idone",
			},
			expectErr: common.ErrPriceIsRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := assert.New(t)

			err := entity.SaveProductPriceHistory(tt.productHistory.productID, tt.productHistory.newPrice)
			if err != nil {
				as.Equal(err, tt.expectErr)
				return
			}
			as.Nil(err)
		})

	}
}
