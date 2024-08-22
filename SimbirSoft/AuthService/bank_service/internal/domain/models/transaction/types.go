package transaction

import "time"

const (
	TopUp = iota + 1
	Withdraw
	Transfer
)

type Transaction struct {
	ID                     int
	AccountID              int
	ParticipatingAccountID int
	TransactionType        int
	Amount                 int
	Date                   time.Time
}
