package entity

import (
	"github.com/google/uuid"
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"time"
)

type WatchList struct {
	ID        string    `json:"id"`
	ProductID string    `json:"book_id"`
	UserID    string    `json:"user_id"`
	Date      time.Time `json:"date"`
}

func NewWatchList(productID, userID string) (*WatchList, error) {
	w := &WatchList{
		ID:        uuid.New().String(),
		ProductID: productID,
		UserID:    userID,
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
