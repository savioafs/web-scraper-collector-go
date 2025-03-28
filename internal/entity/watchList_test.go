package entity

import (
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWatchList(t *testing.T) {
	type watchListInput struct {
		productID string
		userID    string
	}
	tests := []struct {
		name      string
		watchList watchListInput
		expectErr error
	}{
		{
			name: "Create watch list with success",
			watchList: watchListInput{
				productID: "productID",
				userID:    "userID",
			},
			expectErr: nil,
		},
		{
			name: "Create watch list with error product id is required",
			watchList: watchListInput{
				userID: "userID",
			},
			expectErr: common.ErrProductIDIsRequired,
		},
		{
			name: "Create watch list with error user id is required",
			watchList: watchListInput{
				productID: "productID",
			},
			expectErr: common.ErrUserIDIsRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := assert.New(t)

			watchList, err := NewWatchList(tt.watchList.productID, tt.watchList.userID)
			if err != nil {
				as.Equal(err, tt.expectErr)
				return
			}
			as.Nil(err)
			as.NotNil(watchList)
			as.Equal(watchList.ProductID, tt.watchList.productID)
			as.Equal(watchList.UserID, tt.watchList.userID)

		})
	}
}
