package common

import "errors"

var (
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrURLIsRequired   = errors.New("url is required")

	ErrEmailIsRequired = errors.New("email is required")

	ErrProductIDIsRequired = errors.New("product id is required")
)
