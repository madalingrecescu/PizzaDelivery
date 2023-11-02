package pizzas_db

import (
	"context"
	"database/sql"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomPizza(t *testing.T) Pizza {

	arg := CreatePizzaParams{
		Name:        util.RandomNameOrEmail(4, false),
		Description: util.RandomPass(6),
		Price:       5,
	}

	account, err := testQueries.CreatePizza(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.Description, account.Description)
	require.Equal(t, arg.Price, account.Price)

	require.NotZero(t, account.PizzaID)

	return account
}

func TestCreatePizza(t *testing.T) {
	createRandomPizza(t)
}

func TestGetPizzaById(t *testing.T) {
	//create account
	pizza1 := createRandomPizza(t)
	pizza2, err := testQueries.GetPizzaById(context.Background(), pizza1.PizzaID)
	require.NoError(t, err)
	require.NotEmpty(t, pizza2)

	require.Equal(t, pizza1.PizzaID, pizza2.PizzaID)
	require.Equal(t, pizza1.Name, pizza2.Name)
	require.Equal(t, pizza1.Description, pizza2.Description)
	require.Equal(t, pizza1.Price, pizza2.Price)
}

func TestGetPizzaByName(t *testing.T) {
	//create account
	pizza1 := createRandomPizza(t)
	pizza2, err := testQueries.GetPizzaByName(context.Background(), pizza1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, pizza2)

	require.Equal(t, pizza1.PizzaID, pizza2.PizzaID)
	require.Equal(t, pizza1.Name, pizza2.Name)
	require.Equal(t, pizza1.Description, pizza2.Description)
	require.Equal(t, pizza1.Price, pizza2.Price)
}

func TestUpdatePizza(t *testing.T) {
	pizza1 := createRandomPizza(t)

	arg := UpdatePizzaParams{
		PizzaID:     pizza1.PizzaID,
		Name:        "updated",
		Description: "updated",
		Price:       8,
	}

	pizza2, err := testQueries.UpdatePizza(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pizza2)

	require.Equal(t, arg.PizzaID, pizza2.PizzaID)
	require.Equal(t, arg.Name, pizza2.Name)
	require.Equal(t, arg.Description, pizza2.Description)
	require.Equal(t, arg.Price, pizza2.Price)
}

func TestDeletePizza(t *testing.T) {
	pizza1 := createRandomPizza(t)
	err := testQueries.DeletePizza(context.Background(), pizza1.PizzaID)
	require.NoError(t, err)

	pizza2, err := testQueries.GetPizzaById(context.Background(), pizza1.PizzaID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, pizza2)
}

func TestGetAllPizzas(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomPizza(t)
	}

	accounts, err := testQueries.GetAllPizzas(context.Background())
	require.NoError(t, err)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
