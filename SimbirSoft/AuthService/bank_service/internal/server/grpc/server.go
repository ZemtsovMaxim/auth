package server

import (
	auth_v1 "bank_service/api/gen/auth"
	bank_v1 "bank_service/api/gen/bank"
	"bank_service/internal/service/bank"
	"bank_service/internal/storage"
	"bank_service/pkg/grpc/client"
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	authGRPCAddress = "auth:8080"
)

type Bank interface {
	// Returning created account ID, balance and error
	CreateAccount(ctx context.Context, fullName, citizenship string, balance int64) (int64, int64, error)
	// Returning updated account balance and error
	AccountTopUp(ctx context.Context, accountID, amount int64) (int64, error)
	// Returning updated account balance and error
	AccountWithdraw(ctx context.Context, accountID, amount int64) (int64, error)
	// Returning updated accounts balances and error
	AccountTransfer(ctx context.Context, writeOfAccountID, beneficiaryAccountID, amount int64) (int64, int64, error)
	AccountLock(ctx context.Context, accountID int64) error
}

type serverAPI struct {
	bank_v1.UnimplementedBankServer
	bank       Bank
	cl_auth_v1 *client.ClientGRPC
}

func Register(gRPC *grpc.Server, bank Bank) error {
	cl, err := client.NewClientGRPC(authGRPCAddress)
	if err != nil {
		log.Fatalf("couldn't make grpc auth client: %w", err)
		return err
	}

	bank_v1.RegisterBankServer(gRPC, &serverAPI{bank: bank, cl_auth_v1: cl})

	return nil
}

func (s *serverAPI) CreateAccount(ctx context.Context, req *bank_v1.CreateAccountRequest) (*bank_v1.CreateAccountResponse, error) {
	reqValidateToken := &auth_v1.ValidateTokenRequest{Token: req.Jwt}
	_, err := s.cl_auth_v1.ValidateToken(ctx, reqValidateToken)
	st, ok := status.FromError(err)
	if !ok {
		return nil, status.Error(codes.Internal, "couldn't proceed jwt validation by auth service")
	}
	if err != nil {
		return nil, status.Error(st.Code(), st.Message())
	}

	err = validateCreateAccount(req)
	if err != nil {
		return nil, err
	}

	account_id, balance, err := s.bank.CreateAccount(ctx, req.FullName, req.Citizenship, req.Balance)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create account")
	}

	return &bank_v1.CreateAccountResponse{AccountId: account_id, Balance: balance}, nil
}

func (s *serverAPI) AccountTopUp(ctx context.Context, req *bank_v1.AccountTopUpRequest) (*bank_v1.AccountTopUpResponse, error) {
	reqValidateToken := &auth_v1.ValidateTokenRequest{Token: req.Jwt}
	_, err := s.cl_auth_v1.ValidateToken(ctx, reqValidateToken)
	st, ok := status.FromError(err)
	if !ok {
		return nil, status.Error(codes.Internal, "couldn't proceed jwt validation by auth service")
	}
	if err != nil {
		return nil, status.Error(st.Code(), st.Message())
	}

	err = validateAccountTopUp(req)
	if err != nil {
		return nil, err
	}

	balance, err := s.bank.AccountTopUp(ctx, req.AccountId, req.TopUpAmount)
	if errors.Is(err, bank.ErrAccountIDDoesNotExist) {
		return nil, status.Error(codes.InvalidArgument, "account id does not exist")
	}
	if errors.Is(err, bank.ErrAccountLocked) {
		return nil, status.Error(codes.FailedPrecondition, "account is locked")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to top up account")
	}

	return &bank_v1.AccountTopUpResponse{Balance: balance}, nil
}

func (s *serverAPI) AccountWithdraw(ctx context.Context, req *bank_v1.AccountWithdrawRequest) (*bank_v1.AccountWithdrawResponse, error) {
	reqValidateToken := &auth_v1.ValidateTokenRequest{Token: req.Jwt}
	_, err := s.cl_auth_v1.ValidateToken(ctx, reqValidateToken)
	st, ok := status.FromError(err)
	if !ok {
		return nil, status.Error(codes.Internal, "couldn't proceed jwt validation by auth service")
	}
	if err != nil {
		return nil, status.Error(st.Code(), st.Message())
	}

	err = validateAccountWithdraw(req)
	if err != nil {
		return nil, err
	}

	balance, err := s.bank.AccountWithdraw(ctx, req.AccountId, req.WithdrawAmount)
	if errors.Is(err, bank.ErrAccountIDDoesNotExist) {
		return nil, status.Error(codes.InvalidArgument, "account id does not exist")
	}
	if errors.Is(err, bank.ErrAccountLocked) {
		return nil, status.Error(codes.FailedPrecondition, "account is locked")
	}
	if errors.Is(err, storage.ErrNotEnoughMoney) {
		return nil, status.Error(codes.ResourceExhausted, "not enough money")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to withdraw from account")
	}

	return &bank_v1.AccountWithdrawResponse{Balance: balance}, nil
}

func (s *serverAPI) AccountTransfer(ctx context.Context, req *bank_v1.AccountTransferRequest) (*bank_v1.AccountTransferResponse, error) {
	reqValidateToken := &auth_v1.ValidateTokenRequest{Token: req.Jwt}
	_, err := s.cl_auth_v1.ValidateToken(ctx, reqValidateToken)
	st, ok := status.FromError(err)
	if !ok {
		return nil, status.Error(codes.Internal, "couldn't proceed jwt validation by auth service")
	}
	if err != nil {
		return nil, status.Error(st.Code(), st.Message())
	}

	err = validateAccountTransfer(req)
	if err != nil {
		return nil, err
	}

	writeOffAccountBalance, beneficiaryAccountBalance, err := s.bank.AccountTransfer(ctx, req.WriteOffAccountId, req.BeneficiaryAccountId, req.TransferAmount)
	if errors.Is(err, bank.ErrAccountIDDoesNotExist) {
		return nil, status.Error(codes.InvalidArgument, "account id does not exist")
	}
	if errors.Is(err, bank.ErrAccountLocked) {
		return nil, status.Error(codes.FailedPrecondition, "account is locked")
	}
	if errors.Is(err, storage.ErrNotEnoughMoney) {
		return nil, status.Error(codes.ResourceExhausted, "not enough money")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to transfer")
	}

	return &bank_v1.AccountTransferResponse{
		WriteOffAccountBalance:    writeOffAccountBalance,
		BeneficiaryAccountBalance: beneficiaryAccountBalance}, nil
}

func (s *serverAPI) AccountLock(ctx context.Context, req *bank_v1.AccountLockRequest) (*bank_v1.AccountLockResponse, error) {
	// Forgot to add jwt in request

	err := validateAccounLock(req)
	if err != nil {
		return nil, err
	}

	err = s.bank.AccountLock(ctx, req.AccountId)
	if errors.Is(err, bank.ErrAccountIDDoesNotExist) {
		return nil, status.Error(codes.InvalidArgument, "account id does not exist")
	}
	if errors.Is(err, bank.ErrAccountLocked) {
		return nil, status.Error(codes.FailedPrecondition, "account is already locked")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to lock account")
	}

	return &bank_v1.AccountLockResponse{}, nil
}

func validateCreateAccount(req *bank_v1.CreateAccountRequest) error {
	if req.FullName == "" {
		return status.Error(codes.InvalidArgument, "full name is required")
	}

	if req.Citizenship == "" {
		return status.Error(codes.InvalidArgument, "citizenship is required")
	}

	if req.Balance < 0 {
		return status.Error(codes.InvalidArgument, "incorrect balance")
	}

	return nil
}

func validateAccountTopUp(req *bank_v1.AccountTopUpRequest) error {
	if req.AccountId <= 0 {
		return status.Error(codes.InvalidArgument, "incorrect or empty account id")
	}

	if req.TopUpAmount <= 0 {
		return status.Error(codes.InvalidArgument, "incorrect or empty amount")
	}

	return nil
}

func validateAccountWithdraw(req *bank_v1.AccountWithdrawRequest) error {
	if req.AccountId <= 0 {
		return status.Error(codes.InvalidArgument, "incorrect or empty account id")
	}

	if req.WithdrawAmount <= 0 {
		return status.Error(codes.InvalidArgument, "incorrect or empty amount")
	}

	return nil
}
func validateAccountTransfer(req *bank_v1.AccountTransferRequest) error {
	if req.WriteOffAccountId <= 0 {
		return status.Error(codes.InvalidArgument, "incorrect or empty write off account id")
	}

	if req.BeneficiaryAccountId <= 0 {
		return status.Error(codes.InvalidArgument, "incorrect or empty beneficiary account id")
	}

	if req.TransferAmount <= 0 {
		return status.Error(codes.InvalidArgument, "incorrect or empty amount")
	}

	return nil
}
func validateAccounLock(req *bank_v1.AccountLockRequest) error {
	if req.AccountId <= 0 {
		return status.Error(codes.InvalidArgument, "incorrect or empty account id")
	}

	return nil
}
