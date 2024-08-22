package tests

import (
	bank_v1 "bank_service/api/gen/bank"
	"bank_service/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_AccountTopUp_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(250)
	topUpAmount := int64(50)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)

	reqAccountTopUp := &bank_v1.AccountTopUpRequest{Jwt: validJWT, AccountId: respCreateAccount.AccountId, TopUpAmount: topUpAmount}
	respAccountTopUp, err := st.BankClient.AccountTopUp(ctx, reqAccountTopUp)
	require.NoError(t, err)

	assert.Equal(t, respAccountTopUp.Balance, balance+topUpAmount)
}

func Test_AccountTopUp_LockedAccount(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(500)
	topUpAmount := int64(50)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)

	reqAccountLock := &bank_v1.AccountLockRequest{AccountId: respCreateAccount.AccountId}
	_, err = st.BankClient.AccountLock(ctx, reqAccountLock)
	require.NoError(t, err)

	reqAccountTopUp := &bank_v1.AccountTopUpRequest{Jwt: validJWT, AccountId: respCreateAccount.AccountId, TopUpAmount: topUpAmount}
	_, err = st.BankClient.AccountTopUp(ctx, reqAccountTopUp)
	require.Error(t, err)
	require.Contains(t, err.Error(), "account is locked")
}

func Test_AccountTopUp_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(500)
	topUpAmount := int64(50)

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
			amount:      topUpAmount,
			expectedErr: "invalid token",
		},
		{
			name:        "nonexistent account id",
			jwt:         validJWT,
			accountID:   int64(1000000),
			amount:      topUpAmount,
			expectedErr: "account id does not exist",
		},
		{
			name:        "empty account id",
			jwt:         validJWT,
			accountID:   int64(0),
			amount:      topUpAmount,
			expectedErr: "incorrect or empty account id",
		},
		{
			name:        "negative account id",
			jwt:         validJWT,
			accountID:   int64(-1),
			amount:      topUpAmount,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.BankClient.AccountTopUp(ctx, &bank_v1.AccountTopUpRequest{
				Jwt:         tt.jwt,
				AccountId:   tt.accountID,
				TopUpAmount: tt.amount,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
