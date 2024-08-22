package suite

import (
	bank_v1 "bank_service/api/gen/bank"
	"bank_service/internal/config"
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	BankClient bank_v1.BankClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadPath("../config/default.yaml")

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.MyGRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.NewClient(cfg.MyGRPC.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("can't make client connection", err)
	}

	return ctx, &Suite{T: t, Cfg: cfg, BankClient: bank_v1.NewBankClient(cc)}
}
