package bank

import (
	"bank_service/internal/domain/models"
	"bank_service/internal/storage"
	"bank_service/internal/storage/postgres"
	"bank_service/pkg/logger/sl"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

var (
	ErrAccountLocked         = errors.New("account id is locked")
	ErrAccountIDDoesNotExist = errors.New("account id does not exist")
)

type Bank struct {
	log               *slog.Logger
	ownerModifier     OwnerModifier
	accountModifier   AccountModifier
	accountTransacter AccountTransacter
}

type OwnerModifier interface {
	CreateOwner(ctx context.Context, fullName, citizenship string) (models.Owner, error)
	Owner(ctx context.Context, fullName string) (models.Owner, error)
}

type AccountModifier interface {
	// Returning created account ID, balance and error
	CreateAccount(ctx context.Context, ownerID, balance int64) (models.Account, error)
	Account(ctx context.Context, id int64) (models.Account, error)
	AccountLock(ctx context.Context, id int64) error
}

type AccountTransacter interface {
	// Returning updated account balance and error
	AccountTopUp(ctx context.Context, accountID, amount int64) (int64, error)
	// Returning updated account balance and error
	AccountWithdraw(ctx context.Context, accountID, amount int64) (int64, error)
	// Returning updated accounts balances and error
	AccountTransfer(ctx context.Context, writeOfAccountID, beneficiaryAccountID, amount int64) (int64, int64, error)
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
) *Bank {
	return &Bank{
		log:               log,
		ownerModifier:     storage,
		accountModifier:   storage,
		accountTransacter: storage,
	}
}

// Returning created account ID, balance and error
func (b *Bank) CreateAccount(ctx context.Context, fullName, citizenship string, balance int64) (int64, int64, error) {
	const op = "internal.service.bank.CreateAccount"

	log := b.log.With(slog.String("op", op))

	// Check on owner and get his id if exist
	owner, err := b.ownerModifier.Owner(ctx, fullName)
	if err != nil {
		if !errors.Is(err, storage.ErrOwnerDoesNotExist) { // If owner exist but err was not nil return by DB error
			log.Error("failed to get owner", sl.Err(err))
			return 0, 0, fmt.Errorf("%s: %w", op, err)
		}

		owner, err = b.ownerModifier.CreateOwner(ctx, fullName, citizenship)
		if err != nil {
			log.Error("failed to create owner", sl.Err(err))
			return 0, 0, fmt.Errorf("%s: %w", op, err)
		}
	}

	account, err := b.accountModifier.CreateAccount(ctx, owner.ID, balance)
	if err != nil {
		log.Error("failed to create account", sl.Err(err))
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	return account.ID, account.Balance, nil
}

// Returning updated account balance and error
func (b *Bank) AccountTopUp(ctx context.Context, accountID, amount int64) (int64, error) {
	const op = "internal.service.bank.AccountTopUp"

	log := b.log.With(slog.String("op", op))

	// Check on account
	_, err := validateAccount(ctx, b, accountID)
	if err != nil {
		log.Error("failed to validate account id", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	balance, err := b.accountTransacter.AccountTopUp(ctx, accountID, amount)
	if err != nil {
		log.Error("failed to get account", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return balance, nil
}

// Returning updated account balance and error
func (b *Bank) AccountWithdraw(ctx context.Context, accountID, amount int64) (int64, error) {
	const op = "internal.service.bank.AccountWithdraw"

	log := b.log.With(slog.String("op", op))

	// Check on account
	_, err := validateAccount(ctx, b, accountID)
	if err != nil {
		log.Error("failed to validate account id", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	balance, err := b.accountTransacter.AccountWithdraw(ctx, accountID, amount)
	if err != nil {
		log.Error("failed to withdraw", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return balance, nil
}

// Returning updated accounts balances and error
func (b *Bank) AccountTransfer(ctx context.Context, writeOfAccountID, beneficiaryAccountID, amount int64) (int64, int64, error) {
	const op = "internal.service.bank.AccountTransfer"

	log := b.log.With(slog.String("op", op))

	// Check on account
	_, err := validateAccount(ctx, b, writeOfAccountID)
	if err != nil {
		log.Error("failed to validate write off account id", sl.Err(err))
		return 0, 0, fmt.Errorf("%s: write off %w", op, err)
	}

	// Check on beneficiary account
	_, err = validateAccount(ctx, b, beneficiaryAccountID)
	if err != nil {
		log.Error("failed to validate beneficiary account id", sl.Err(err))
		return 0, 0, fmt.Errorf("%s: beneficiary %w", op, err)
	}

	writeOffAccountBalance, bebeneficiaryAccountBalance, err := b.accountTransacter.AccountTransfer(ctx, writeOfAccountID, beneficiaryAccountID, amount)
	if err != nil {
		log.Error("failed to transfer", sl.Err(err))
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	return writeOffAccountBalance, bebeneficiaryAccountBalance, nil
}

func (b *Bank) AccountLock(ctx context.Context, accountID int64) error {
	const op = "internal.service.bank.AccountLock"

	log := b.log.With(slog.String("op", op))

	// Check on account
	_, err := validateAccount(ctx, b, accountID)
	if errors.Is(err, ErrAccountLocked) {
		log.Error("account id is already locked")
		return fmt.Errorf("%s: %w", op, ErrAccountLocked)
	}
	if err != nil {
		log.Error("failed to validate account id", sl.Err(err))
		return fmt.Errorf("%s: failed to validate account id", op)
	}

	err = b.accountModifier.AccountLock(ctx, accountID)
	if err != nil {
		log.Error("failed to lock", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Return true if account id exists and not locked
func validateAccount(ctx context.Context, b *Bank, accountID int64) (models.Account, error) {
	const op = "validateAccount"

	log := b.log.With(slog.String("op", op))

	account, err := b.accountModifier.Account(ctx, accountID)
	if errors.Is(err, storage.ErrAccountDoesNotExist) {
		log.Error("account id does not exist")
		return models.Account{}, fmt.Errorf("%s: %w", op, ErrAccountIDDoesNotExist)
	}
	if err != nil {
		log.Error("failed to get account id", sl.Err(err))
		return models.Account{}, fmt.Errorf("%s: %w", op, err)
	}

	if account.IsLocked {
		log.Error("account id is locked")
		return models.Account{}, fmt.Errorf("%s: %w", op, ErrAccountLocked)
	}

	return account, nil
}
