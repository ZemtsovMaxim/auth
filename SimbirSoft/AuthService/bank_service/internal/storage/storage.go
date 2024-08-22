package storage

import "errors"

var (
	ErrOwnerDoesNotExist   = errors.New("owner does not exist")
	ErrAccountDoesNotExist = errors.New("account does not exist")
	ErrNotEnoughMoney      = errors.New("not enough money")
)
