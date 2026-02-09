package db

import (
	"context"
	"testing"

	"github.com/0abdullahbhutto0/bank_backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfers {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	account3, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account3)

	require.Equal(t, account1.ID, account3.FromAccountID)
	require.Equal(t, account2.ID, account3.ToAccountID)
	require.Equal(t, arg.Amount, account3.Amount)

	require.NotEmpty(t, account3.ID)
	require.NotEmpty(t, account3.CreatedAt)

	return account3

}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}
