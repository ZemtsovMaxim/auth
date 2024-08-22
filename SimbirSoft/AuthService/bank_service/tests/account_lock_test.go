package tests

import (
	bank_v1 "bank_service/api/gen/bank"
	"bank_service/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_AccountLock_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(500)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)

	reqAccountLock := &bank_v1.AccountLockRequest{AccountId: respCreateAccount.AccountId}
	_, err = st.BankClient.AccountLock(ctx, reqAccountLock)
	require.NoError(t, err)
}

func Test_AccountLock_DoubleLock(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(500)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)

	reqAccountLock := &bank_v1.AccountLockRequest{AccountId: respCreateAccount.AccountId}
	_, err = st.BankClient.AccountLock(ctx, reqAccountLock)
	require.NoError(t, err)

	_, err = st.BankClient.AccountLock(ctx, reqAccountLock)
	require.Error(t, err)
	require.Contains(t, err.Error(), "account is already locked")
}

func Test_AccountLock_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(500)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)

	tests := []struct {
		name        string
		accountID   int64
		expectedErr string
	}{{
		name:        "empty account id",
		accountID:   int64(0),
		expectedErr: "incorrect or empty account id",
	},
		{
			name:        "negative account id",
			accountID:   int64(-1),
			expectedErr: "incorrect or empty account id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.BankClient.AccountLock(ctx, &bank_v1.AccountLockRequest{
				AccountId: tt.accountID,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
