package users_db

import (
	"context"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomPass(6))
	require.NoError(t, err)

	arg := CreateAccountParams{
		Username:       util.RandomNameOrEmail(5, false),
		Email:          util.RandomNameOrEmail(5, true),
		HashedPassword: hashedPassword,
		PhoneNumber:    util.RandomPhoneNumber(10),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.HashedPassword, account.HashedPassword)
	require.Equal(t, arg.PhoneNumber, account.PhoneNumber)

	require.NotZero(t, account.UserID)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	//create account
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.PhoneNumber, account2.PhoneNumber)
	require.Equal(t, account1.HashedPassword, account2.HashedPassword)
	require.Equal(t, account1.Role, account2.Role)
}
