package entity

import (
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"time"
)

type WatchList struct {
	ProductID string    `json:"book_id"`
	UserID    string    `json:"user_id"`
	Date      time.Time `json:"date"`
}

func NewWatchList(productID, useID string) (*WatchList, error) {
	w := &WatchList{
		ProductID: productID,
		UserID:    useID,
		Date:      time.Now(),
	}

	err := w.Validate()
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *WatchList) Validate() error {
	if w.UserID == "" {
		return common.ErrUserIDIsRequired
	}

	if w.ProductID == "" {
		return common.ErrProductIDIsRequired
	}

	return nil
}
