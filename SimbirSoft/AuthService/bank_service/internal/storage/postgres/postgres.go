package postgres

import (
	"bank_service/internal/domain/models"
	"bank_service/internal/domain/models/transaction"
	"bank_service/internal/storage"
	"context"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
)

const (
	qrCreateOwner      = `INSERT INTO owner(full_name, citizenship) VALUES ($1, $2) RETURNING id;`
	qrOwner            = `SELECT * FROM owner WHERE full_name = $1;`
	qrCreateAccount    = `INSERT INTO account(owner_id, balance) VALUES ($1, $2) RETURNING id;`
	qrAccount          = `SELECT * FROM account WHERE id = $1;`
	qrAccountLock      = `UPDATE account SET is_locked = TRUE WHERE id = $1;`
	qrAccountTopUp     = `UPDATE account SET balance = balance + $1 WHERE id = $2;`
	qrAccountBalance   = `SELECT balance FROM account WHERE id = $1 FOR UPDATE;`
	qrAccountWithdraw  = `UPDATE account SET balance = balance - $1 WHERE id = $2;`
	qrCeateTransaction = `INSERT INTO transaction(account_id, participating_account_id, transaction_type, amount, date)
						  VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP);`
	qrTransaction = `SELECT * FROM transaction WHERE account_id = $1;`
)

type Storage struct {
	db *sql.DB
}

func New(storageDriver, storageInfo string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open(storageDriver, storageInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

func (s *Storage) MigrationUp(storageURL, migrationPath string) error {
	const op = "storage.postgres.MigrationUp"

	migration, err := migrate.New(migrationPath, storageURL)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = migration.Up()
	if err != nil && migration.Up().Error() != "no change" {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) MigrationDown(storageURL, migrationPath string) error {
	const op = "storage.postgres.MigrationDown"

	migration, err := migrate.New(migrationPath, storageURL)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = migration.Down()
	if err != nil && migration.Down().Error() != "no change" {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Returning created owner and error
func (s *Storage) CreateOwner(ctx context.Context, fullName, citizenship string) (models.Owner, error) {
	const op = "storage.postgres.CreateOwner"

	var owner models.Owner

	tx, err := s.db.Begin()
	if err != nil {
		return models.Owner{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrCreateOwner, fullName, citizenship).Scan(&owner.ID)
	if err != nil {
		tx.Rollback()
		return models.Owner{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrOwner, fullName).Scan(&owner.ID, &owner.FullName, &owner.Citizenship)
	if err != nil {
		tx.Rollback()
		return models.Owner{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return models.Owner{}, fmt.Errorf("%s: %w", op, err)
	}

	return owner, nil
}

func (s *Storage) Owner(ctx context.Context, fullName string) (models.Owner, error) {
	const op = "storage.postgres.Owner"

	var owner models.Owner

	err := s.db.QueryRowContext(ctx, qrOwner, fullName).Scan(&owner.ID, &owner.FullName, &owner.Citizenship)
	if err != nil {
		return models.Owner{}, fmt.Errorf("%s: %w", op, storage.ErrOwnerDoesNotExist)
	}

	return owner, nil
}

// Returning created account ID, balance and error
func (s *Storage) CreateAccount(ctx context.Context, ownerID, balance int64) (models.Account, error) {
	const op = "storage.postgres.CreateAccount"

	var account models.Account

	tx, err := s.db.Begin()
	if err != nil {
		return models.Account{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrCreateAccount, ownerID, balance).Scan(&account.ID)
	if err != nil {
		tx.Rollback()
		return models.Account{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrAccount, account.ID).Scan(&account.ID, &account.Balance, &account.OwnerID, &account.IsLocked)
	if err != nil {
		tx.Rollback()
		return models.Account{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return models.Account{}, fmt.Errorf("%s: %w", op, err)
	}

	return account, nil
}

func (s *Storage) Account(ctx context.Context, id int64) (models.Account, error) {
	const op = "storage.postgres.Account"

	var account models.Account

	err := s.db.QueryRowContext(ctx, qrAccount, id).Scan(&account.ID, &account.Balance, &account.OwnerID, &account.IsLocked)
	if err != nil {
		return models.Account{}, fmt.Errorf("%s: %w", op, storage.ErrAccountDoesNotExist)
	}

	return account, nil
}

func (s *Storage) AccountLock(ctx context.Context, id int64) error {
	const op = "storage.postgres.AccountLock"

	_, err := s.db.ExecContext(ctx, qrAccountLock, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Returning updated account balance and error
func (s *Storage) AccountTopUp(ctx context.Context, accountID, amount int64) (int64, error) {
	const op = "storage.postgres.AccountTopUp"

	var balance int64

	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.ExecContext(ctx, qrAccountTopUp, amount, accountID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrAccountBalance, accountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.ExecContext(ctx, qrCeateTransaction, accountID, nil, transaction.TopUp, amount)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return balance, nil
}

// Returning updated account balance and error
func (s *Storage) AccountWithdraw(ctx context.Context, accountID, amount int64) (int64, error) {
	const op = "storage.postgres.AccountWithdraw"

	var balance int64

	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrAccountBalance, accountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	if amount > balance {
		return 0, fmt.Errorf("%s: %w", op, storage.ErrNotEnoughMoney)
	}

	_, err = tx.ExecContext(ctx, qrAccountWithdraw, amount, accountID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrAccountBalance, accountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.ExecContext(ctx, qrCeateTransaction, accountID, nil, transaction.Withdraw, amount)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return balance, nil
}

// Returning updated wtite off and beneficiary accounts balances and error
func (s *Storage) AccountTransfer(ctx context.Context, writeOfAccountID, beneficiaryAccountID, amount int64) (int64, int64, error) {
	const op = "storage.postgres.AccountTransfer"

	var writeOffBalance, beneficiaryBalance int64

	tx, err := s.db.Begin()
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrAccountBalance, writeOfAccountID).Scan(&writeOffBalance)
	if err != nil {
		tx.Rollback()
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	if amount > writeOffBalance {
		return 0, 0, fmt.Errorf("%s: %w", op, storage.ErrNotEnoughMoney)
	}

	_, err = tx.ExecContext(ctx, qrAccountWithdraw, amount, writeOfAccountID)
	if err != nil {
		tx.Rollback()
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrAccountBalance, writeOfAccountID).Scan(&writeOffBalance)
	if err != nil {
		tx.Rollback()
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.ExecContext(ctx, qrAccountTopUp, amount, beneficiaryAccountID)
	if err != nil {
		tx.Rollback()
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.QueryRowContext(ctx, qrAccountBalance, beneficiaryAccountID).Scan(&beneficiaryBalance)
	if err != nil {
		tx.Rollback()
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.ExecContext(ctx, qrCeateTransaction, writeOfAccountID, beneficiaryAccountID, transaction.Transfer, amount)
	if err != nil {
		tx.Rollback()
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	return writeOffBalance, beneficiaryBalance, nil
}
