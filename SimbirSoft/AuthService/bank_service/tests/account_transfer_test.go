package tests

import (
	bank_v1 "bank_service/api/gen/bank"
	"bank_service/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_AccountTransfer_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	writeOffFullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	writeOffCitizenship := gofakeit.Country()
	writeOffBalance := int64(250)
	beneficiaryFullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	beneficiaryCitizenship := gofakeit.Country()
	beneficiaryBalance := int64(250)
	transferAmount := int64(50)

	reqCreateWriteOffAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: writeOffFullName, Citizenship: writeOffCitizenship, Balance: writeOffBalance}
	respCreateWriteOffAccount, err := st.BankClient.CreateAccount(ctx, reqCreateWriteOffAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateWriteOffAccount.AccountId)
	assert.Equal(t, respCreateWriteOffAccount.Balance, writeOffBalance)

	reqCreateBeneficiaryAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: beneficiaryFullName, Citizenship: beneficiaryCitizenship, Balance: beneficiaryBalance}
	respCreateBeneficiaryAccount, err := st.BankClient.CreateAccount(ctx, reqCreateBeneficiaryAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateBeneficiaryAccount.AccountId)
	assert.Equal(t, respCreateBeneficiaryAccount.Balance, beneficiaryBalance)

	reqAccountTransfer := &bank_v1.AccountTransferRequest{Jwt: validJWT, WriteOffAccountId: respCreateWriteOffAccount.AccountId, BeneficiaryAccountId: respCreateBeneficiaryAccount.AccountId, TransferAmount: transferAmount}
	respAccountTransfer, err := st.BankClient.AccountTransfer(ctx, reqAccountTransfer)
	require.NoError(t, err)

	assert.Equal(t, respAccountTransfer.WriteOffAccountBalance, writeOffBalance-transferAmount)
	assert.Equal(t, respAccountTransfer.BeneficiaryAccountBalance, beneficiaryBalance+transferAmount)
}

func Test_AccountTransfer_LockedWriteOffAccount(t *testing.T) {
	ctx, st := suite.New(t)

	writeOffFullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	writeOffCitizenship := gofakeit.Country()
	writeOffBalance := int64(250)
	beneficiaryFullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	beneficiaryCitizenship := gofakeit.Country()
	beneficiaryBalance := int64(250)
	transferAmount := int64(50)

	reqCreateWriteOffAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: writeOffFullName, Citizenship: writeOffCitizenship, Balance: writeOffBalance}
	respCreateWriteOffAccount, err := st.BankClient.CreateAccount(ctx, reqCreateWriteOffAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateWriteOffAccount.AccountId)
	assert.Equal(t, respCreateWriteOffAccount.Balance, writeOffBalance)

	reqCreateBeneficiaryAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: beneficiaryFullName, Citizenship: beneficiaryCitizenship, Balance: beneficiaryBalance}
	respCreateBeneficiaryAccount, err := st.BankClient.CreateAccount(ctx, reqCreateBeneficiaryAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateBeneficiaryAccount.AccountId)
	assert.Equal(t, respCreateBeneficiaryAccount.Balance, beneficiaryBalance)

	reqWriteOffAccountLock := &bank_v1.AccountLockRequest{AccountId: respCreateWriteOffAccount.AccountId}
	_, err = st.BankClient.AccountLock(ctx, reqWriteOffAccountLock)
	require.NoError(t, err)

	reqAccountTransfer := &bank_v1.AccountTransferRequest{Jwt: validJWT, WriteOffAccountId: respCreateWriteOffAccount.AccountId, BeneficiaryAccountId: respCreateBeneficiaryAccount.AccountId, TransferAmount: transferAmount}
	_, err = st.BankClient.AccountTransfer(ctx, reqAccountTransfer)
	require.Error(t, err)
	require.Contains(t, err.Error(), "account is locked")
}

func Test_AccountTransfer_LockedBeneficiaryAccount(t *testing.T) {
	ctx, st := suite.New(t)

	writeOffFullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	writeOffCitizenship := gofakeit.Country()
	writeOffBalance := int64(250)
	beneficiaryFullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	beneficiaryCitizenship := gofakeit.Country()
	beneficiaryBalance := int64(250)
	transferAmount := int64(50)

	reqCreateWriteOffAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: writeOffFullName, Citizenship: writeOffCitizenship, Balance: writeOffBalance}
	respCreateWriteOffAccount, err := st.BankClient.CreateAccount(ctx, reqCreateWriteOffAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateWriteOffAccount.AccountId)
	assert.Equal(t, respCreateWriteOffAccount.Balance, writeOffBalance)

	reqCreateBeneficiaryAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: beneficiaryFullName, Citizenship: beneficiaryCitizenship, Balance: beneficiaryBalance}
	respCreateBeneficiaryAccount, err := st.BankClient.CreateAccount(ctx, reqCreateBeneficiaryAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateBeneficiaryAccount.AccountId)
	assert.Equal(t, respCreateBeneficiaryAccount.Balance, beneficiaryBalance)

	reqBeneficiaryAccountLock := &bank_v1.AccountLockRequest{AccountId: respCreateBeneficiaryAccount.AccountId}
	_, err = st.BankClient.AccountLock(ctx, reqBeneficiaryAccountLock)
	require.NoError(t, err)

	reqAccountTransfer := &bank_v1.AccountTransferRequest{Jwt: validJWT, WriteOffAccountId: respCreateWriteOffAccount.AccountId, BeneficiaryAccountId: respCreateBeneficiaryAccount.AccountId, TransferAmount: transferAmount}
	_, err = st.BankClient.AccountTransfer(ctx, reqAccountTransfer)
	require.Error(t, err)
	require.Contains(t, err.Error(), "account is locked")
}

func Test_AccountTransfer_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	writeOffFullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	writeOffCitizenship := gofakeit.Country()
	writeOffBalance := int64(250)
	beneficiaryFullName := gofakeit.FirstName() + " " + gofakeit.MiddleName() + " " + gofakeit.LastName()
	beneficiaryCitizenship := gofakeit.Country()
	beneficiaryBalance := int64(250)
	transferAmount := int64(50)

	reqCreateWriteOffAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: writeOffFullName, Citizenship: writeOffCitizenship, Balance: writeOffBalance}
	respCreateWriteOffAccount, err := st.BankClient.CreateAccount(ctx, reqCreateWriteOffAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateWriteOffAccount.AccountId)
	assert.Equal(t, respCreateWriteOffAccount.Balance, writeOffBalance)

	reqCreateBeneficiaryAccount := &bank_v1.CreateAccountRequest{Jwt: validJWT, FullName: beneficiaryFullName, Citizenship: beneficiaryCitizenship, Balance: beneficiaryBalance}
	respCreateBeneficiaryAccount, err := st.BankClient.CreateAccount(ctx, reqCreateBeneficiaryAccount)
	require.NoError(t, err)

	assert.NotEmpty(t, respCreateBeneficiaryAccount.AccountId)
	assert.Equal(t, respCreateBeneficiaryAccount.Balance, beneficiaryBalance)

	tests := []struct {
		name                 string
		jwt                  string
		writeOffAccountID    int64
		beneficiaryAccountID int64
		amount               int64
		expectedErr          string
	}{
		{
			name:                 "invalid jwt",
			jwt:                  invalidJWT,
			writeOffAccountID:    respCreateWriteOffAccount.AccountId,
			beneficiaryAccountID: respCreateBeneficiaryAccount.AccountId,
			amount:               transferAmount,
			expectedErr:          "invalid token",
		},
		{
			name:                 "nonexistent write off account id",
			jwt:                  validJWT,
			writeOffAccountID:    int64(1000000),
			beneficiaryAccountID: respCreateBeneficiaryAccount.AccountId,
			amount:               transferAmount,
			expectedErr:          "account id does not exist",
		},
		{
			name:                 "nonexistent beneficiary account id",
			jwt:                  validJWT,
			writeOffAccountID:    respCreateWriteOffAccount.AccountId,
			beneficiaryAccountID: int64(1000000),
			amount:               transferAmount,
			expectedErr:          "account id does not exist",
		},
		{
			name:                 "nonexistent both account id",
			jwt:                  validJWT,
			writeOffAccountID:    int64(1000000),
			beneficiaryAccountID: int64(1000000),
			amount:               transferAmount,
			expectedErr:          "account id does not exist",
		},
		{
			name:                 "empty write off account id",
			jwt:                  validJWT,
			writeOffAccountID:    int64(-1),
			beneficiaryAccountID: respCreateBeneficiaryAccount.AccountId,
			amount:               transferAmount,
			expectedErr:          "incorrect or empty write off account id",
		},
		{
			name:                 "negative beneficiary account id",
			jwt:                  validJWT,
			writeOffAccountID:    respCreateWriteOffAccount.AccountId,
			beneficiaryAccountID: int64(-1),
			amount:               transferAmount,
			expectedErr:          "incorrect or empty beneficiary account id",
		},
		{
			name:                 "negative both account id",
			jwt:                  validJWT,
			writeOffAccountID:    int64(-1),
			beneficiaryAccountID: int64(-1),
			amount:               transferAmount,
			expectedErr:          "incorrect or empty write off account id",
		},
		{
			name:                 "empty amount",
			jwt:                  validJWT,
			writeOffAccountID:    respCreateWriteOffAccount.AccountId,
			beneficiaryAccountID: respCreateBeneficiaryAccount.AccountId,
			amount:               int64(0),
			expectedErr:          "incorrect or empty amount",
		},
		{
			name:                 "negative amount",
			jwt:                  validJWT,
			writeOffAccountID:    respCreateWriteOffAccount.AccountId,
			beneficiaryAccountID: respCreateBeneficiaryAccount.AccountId,
			amount:               int64(-50),
			expectedErr:          "incorrect or empty amount",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.BankClient.AccountTransfer(ctx, &bank_v1.AccountTransferRequest{
				Jwt:                  tt.jwt,
				WriteOffAccountId:    tt.writeOffAccountID,
				BeneficiaryAccountId: tt.beneficiaryAccountID,
				TransferAmount:       tt.amount,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
