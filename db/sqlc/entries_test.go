package db

import (
	"context"
	"testing"

	"github.com/0abdullahbhutto0/bank_backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entries {
	account1 := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	account2, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.AccountID)
	require.Equal(t, arg.Amount, account2.Amount)

	return account2

}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}
