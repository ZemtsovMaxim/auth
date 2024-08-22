package models

type Account struct {
	ID       int64
	Balance  int64
	OwnerID  int64
	IsLocked bool
}
