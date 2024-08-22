package tests

import (
	bank_v1 "bank_service/api/gen/bank"
	"bank_service/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	validJWT   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5vbiBmdWdpYXQgRXhjZXB0ZUB1ciBjb25zZXF1YXQgdmVuaWFtIiwiZXhwIjoxNzIyNDI0MDA2fQ.M_x8k9S2gaRCEHr1xLnmw0sqlim2TbeFTpGyxc70tDA"
	invalidJWT = "12345"
)

func Test_CreateAccount_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	fullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	citizenship := gofakeit.Country()
	balance := int64(500)

	reqCreateAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: fullName, Citizenship: citizenship, Balance: balance}
	respCreateAccount, err := st.BankClient.CreateAccount(ctx, reqCreateAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateAccount.AccountId)
	assert.Equal(t, respCreateAccount.Balance, balance)
}

func Test_CreateAccount_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		jwt         string
		fullName    string
		citizenship string
		balance     int64
		expectedErr string
	}{
		{
			name:        "invalid jwt",
			jwt:         invalidJWT,
			fullName:    gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName(),
			citizenship: gofakeit.Country(),
			balance:     int64(500),
			expectedErr: "invalid token",
		},
		{
			name:        "empty jwt",
			jwt:         "",
			fullName:    gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName(),
			citizenship: gofakeit.Country(),
			balance:     int64(500),
			expectedErr: "invalid token",
		},
		{
			name:        "empty full name",
			jwt:         validJWT,
			fullName:    "",
			citizenship: gofakeit.Country(),
			balance:     int64(500),
			expectedErr: "full name is required",
		},
		{
			name:        "empty citizenship",
			jwt:         validJWT,
			fullName:    gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName(),
			citizenship: "",
			balance:     int64(500),
			expectedErr: "citizenship is required",
		},
		{
			name:        "negative balance",
			jwt:         validJWT,
			fullName:    gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName(),
			citizenship: gofakeit.Country(),
			balance:     int64(-15),
			expectedErr: "incorrect balance",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.BankClient.CreateAccount(ctx, &bank_v1.CreateAccountRequest{
				Jwt:         tt.jwt,
				FullName:    tt.fullName,
				Citizenship: tt.citizenship,
				Balance:     tt.balance,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
