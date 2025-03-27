package common

import "errors"

var (
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrURLIsRequired   = errors.New("url is required")

	ErrProductNotFound = errors.New("product not found")

	ErrEmailIsRequired = errors.New("email is required")
	ErrInvalidEmail    = errors.New("invalid email")

	ErrProductIDIsRequired = errors.New("product id is required")

	ErrProductAlreadyWatcherByUser = errors.New("product already watcher by user")
	ErrUserNotFound                = errors.New("user not found")

	ErrPasswordIsRequired = errors.New("password is required")

	ErrUserIDIsRequired = errors.New("user id is required")
)
