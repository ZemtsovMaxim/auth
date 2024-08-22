package tests

import (
	bank_v1 "bank_service/api/gen/bank"
	"bank_service/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_AccountWithdraw_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(250)
	withdrawAmount := int64(50)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)

	reqAccountWithdraw := &bank_v1.AccountWithdrawRequest{Jwt: validJWT, AccountId: respCreateAccount.AccountId, WithdrawAmount: withdrawAmount}
	respAccountWithdraw, err := st.BankClient.AccountWithdraw(ctx, reqAccountWithdraw)
	require.NoError(t, err)

	assert.Equal(t, respAccountWithdraw.Balance, balance-withdrawAmount)
}

func Test_AccountWithdraw_LockedAccount(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(250)
	withdrawAmount := int64(50)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)

	reqAccountLock := &bank_v1.AccountLockRequest{AccountId: respCreateAccount.AccountId}
	_, err = st.BankClient.AccountLock(ctx, reqAccountLock)
	require.NoError(t, err)

	reqAccountWithdraw := &bank_v1.AccountWithdrawRequest{Jwt: validJWT, AccountId: respCreateAccount.AccountId, WithdrawAmount: withdrawAmount}
	_, err = st.BankClient.AccountWithdraw(ctx, reqAccountWithdraw)
	require.Error(t, err)
	require.Contains(t, err.Error(), "account is locked")
}

func Test_AccountWithdraw_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(500)
	withdrawAmount := int64(50)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)

	tests := []struct {
		name        string
		jwt         string
		accountID   int64
		amount      int64
		expectedErr string
	}{
		{
			name:        "invalid jwt",
			jwt:         invalidJWT,
			accountID:   respCreateAccount.AccountId,
			amount:      withdrawAmount,
			expectedErr: "invalid token",
		},
		{
			name:        "nonexistent account id",
			jwt:         validJWT,
			accountID:   int64(10000000),
			amount:      withdrawAmount,
			expectedErr: "account id does not exist",
		},
		{
			name:        "empty account id",
			jwt:         validJWT,
			accountID:   int64(0),
			amount:      withdrawAmount,
			expectedErr: "incorrect or empty account id",
		},
		{
			name:        "negative account id",
			jwt:         validJWT,
			accountID:   int64(-1),
			amount:      withdrawAmount,
			expectedErr: "incorrect or empty account id",
		},
		{
			name:        "empty amount",
			jwt:         validJWT,
			accountID:   respCreateAccount.AccountId,
			amount:      int64(0),
			expectedErr: "incorrect or empty amount",
		},
		{
			name:        "negative amount",
			jwt:         validJWT,
			accountID:   respCreateAccount.AccountId,
			amount:      int64(-15),
			expectedErr: "incorrect or empty amount",
		},
		{
			name:        "too much amount",
			jwt:         validJWT,
			accountID:   respCreateAccount.AccountId,
			amount:      balance + withdrawAmount,
			expectedErr: "not enough money",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.BankClient.AccountWithdraw(ctx, &bank_v1.AccountWithdrawRequest{
				Jwt:            tt.jwt,
				AccountId:      tt.accountID,
				WithdrawAmount: tt.amount,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
