package storage

import "errors"

var (
	ErrUserExists = errors.New("duplicate key value")

	ErrUserNotFound  = errors.New("user not found")
	ErrUserNotFound2 = errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")
)
