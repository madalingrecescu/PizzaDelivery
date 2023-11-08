package pizzas_handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_pizzas"
	"io/ioutil"
	"net/http"
)

type createOrUpdatePizzaRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func (server *Server) createPizza(ctx *gin.Context) {
	var req createOrUpdatePizzaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePizzaParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	pizza, err := server.store.CreatePizza(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, pizza)
}

func (server *Server) getAllPizzas(ctx *gin.Context) {
	pizzas, err := server.store.GetAllPizzas(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pizzas)
}

type getPizzaByIdRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPizzaById(ctx *gin.Context) {
	var req getPizzaByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pizza, err := server.store.GetPizzaById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errors.New(404, "Pizza id not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pizza)
}

type getPizzaByNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) getPizzaByName(ctx *gin.Context) {
	var req getPizzaByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pizza, err := server.store.GetPizzaByName(ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errors.New(404, "Pizza name not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pizza)
}

func (server *Server) updatePizza(ctx *gin.Context) {
	var reqGet getPizzaByNameRequest
	if err := ctx.ShouldBindUri(&reqGet); err != nil {
		ctx.JSON(http.StatusBadRequest, "1")
		return
	}

	// Check if the pizza exists before attempting an update
	pizza, err := server.store.GetPizzaByName(ctx, reqGet.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Pizza name not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var reqPut createOrUpdatePizzaRequest
	if err := ctx.ShouldBindJSON(&reqPut); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	arg := db.UpdatePizzaParams{
		PizzaID:     pizza.PizzaID,
		Name:        reqPut.Name,
		Description: reqPut.Description,
		Price:       reqPut.Price,
	}

	// Update the pizza only if it exists
	pizza, err = server.store.UpdatePizza(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pizza"})
		return
	}

	ctx.JSON(http.StatusOK, pizza)
}

type deletePizzaByIdRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deletePizza(ctx *gin.Context) {
	var req deletePizzaByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeletePizza(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, "Pizza deleted successfully")
}

//	type createOrUpdatePizzaRequest struct {
//		Name        string  `json:"name" binding:"required"`
//		Description string  `json:"description" binding:"required"`
//		Price       float64 `json:"price" binding:"required"`
//	}

//type createShoppingCartRequest struct {
//	UserID string `json:"user_id" binding:"required"`
//}

func (server *Server) createShoppingCart(ctx *gin.Context) {
	var req struct {
		Username string `uri:"username" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shoppingCart, err := server.store.CreateShoppingCart(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, shoppingCart)
}

type deletePizzaFromShoppingCartRequest struct {
	ShoppingCartId int32  `json:"shopping_cart_id" binding:"required"`
	PizzaName      string `json:"pizza_name" binding:"required"`
}

func (server *Server) deletePizzaFromShoppingCart(ctx *gin.Context) {

	var req deletePizzaFromShoppingCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind the json"})
		return
	}

	arg := db.DeleteAllPizzasWithTheSameNameFromShoppingCartParams{
		ShoppingCartID: req.ShoppingCartId,
		PizzaName:      req.PizzaName,
	}
	err := server.store.DeleteAllPizzasWithTheSameNameFromShoppingCart(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pizzas deleted successfully"})
}

type addPizzaToShoppingCartRequest struct {
	PizzaName string `json:"name" binding:"required"`
	Token     string `json:"token" binding:"required"`
}
type User struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (server *Server) addPizzaToShoppingCart(ctx *gin.Context) {
	var req addPizzaToShoppingCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind the uri"})
		return
	}
	if req.Token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You need to authenticate first"})
		return
	}
	userServiceUrl := "http://localhost:3000/user"
	userReq, err := http.NewRequest("GET", userServiceUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	userReq.Header.Set("Authorization", "Bearer "+req.Token)

	client := &http.Client{}
	resp, err := client.Do(userReq)
	if err != nil {
		fmt.Println("Error 1 request:", err)
		return
	}
	defer resp.Body.Close()

	var user User
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Received non-OK status:", resp.StatusCode)
		// Read the response body to identify the issue
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("Response body:", string(responseBody))
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Error 3 request:", err)

		return
	}

	shoppingCart, err := server.store.GetShoppingCartByUsername(ctx, user.Username)
	if err == sql.ErrNoRows {
		shoppingCart1, err1 := server.store.CreateShoppingCart(ctx, user.Username)
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error at creating shopping cart"})
			return
		}
		shoppingCart = shoppingCart1
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error at getting shopping cart"})
	}

	pizza, err := server.store.GetPizzaByName(ctx, req.PizzaName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Pizza not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	arg := db.GetPizzaOrderByNameFromShoppingCartParams{
		ShoppingCartID: shoppingCart.ShoppingCartID,
		PizzaName:      pizza.Name,
	}
	pizzaOrder, err := server.store.GetPizzaOrderByNameFromShoppingCart(ctx, arg)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find the order"})
			return
		}
	}

	if pizzaOrder.PizzaOrderID == 0 {
		_, err = server.store.CreatePizzaOrder(ctx, db.CreatePizzaOrderParams{
			ShoppingCartID: shoppingCart.ShoppingCartID,
			PizzaName:      pizza.Name,
			PizzaPrice:     pizza.Price,
			Quantity:       int32(1),
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
			return
		}
	} else {
		_, err = server.store.AddQuantityToOrderForExistingPizza(ctx, db.AddQuantityToOrderForExistingPizzaParams{
			PizzaName:      req.PizzaName,
			ShoppingCartID: shoppingCart.ShoppingCartID,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add to quantity"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pizza added to the shopping cart"})
}

type changeQuantityOfPizzasRequest struct {
	ShoppingCartId int32  `json:"shopping_cart_id" binding:"required"`
	PizzaName      string `json:"pizza_name" binding:"required"`
	Quantity       int32  `json:"quantity" binding:"required,min=1"`
}

func (server *Server) changeQuantityOfPizzas(ctx *gin.Context) {

	var req changeQuantityOfPizzasRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind the json"})
		return
	}

	arg := db.GetPizzaOrderByNameFromShoppingCartParams{
		ShoppingCartID: req.ShoppingCartId,
		PizzaName:      req.PizzaName,
	}
	_, err := server.store.GetPizzaOrderByNameFromShoppingCart(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Pizza order not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.UpdatePizzaQuantityInShoppingCart(ctx, db.UpdatePizzaQuantityInShoppingCartParams{
		Quantity:       req.Quantity,
		PizzaName:      req.PizzaName,
		ShoppingCartID: req.ShoppingCartId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pizza's quantity updated successfully"})

}

type orderRequest struct {
	Token string `json:"token" binding:"required"`
}
type orderResponse struct {
	Username    string                `json:"username" binding:"required"`
	Email       string                `json:"Email" binding:"required"`
	PhoneNumber string                `json:"phone_number" binding:"required"`
	Orders      []pizzasOrderResponse `json:"orders"`
	TotalCost   float64               `json:"total-cost"`
}
type pizzasOrderResponse struct {
	PizzaName string `json:"pizza_name"`
	Quantity  int32  `json:"quantity"`
}

func (server *Server) order(ctx *gin.Context) {
	var req orderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind the uri"})
		return
	}
	if req.Token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You need to authenticate first"})
		return
	}
	userServiceUrl := "http://localhost:3000/user"
	userReq, err := http.NewRequest("GET", userServiceUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	userReq.Header.Set("Authorization", "Bearer "+req.Token)

	client := &http.Client{}
	resp, err := client.Do(userReq)
	if err != nil {
		fmt.Println("Error 1 request:", err)
		return
	}
	defer resp.Body.Close()

	var user User
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Received non-OK status:", resp.StatusCode)
		// Read the response body to identify the issue
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("Response body:", string(responseBody))
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Error 3 request:", err)

		return
	}

	shoppingCart, err := server.store.GetShoppingCartByUsername(ctx, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error at getting shopping cart"})
	}
	orders, err := server.store.GetAllOrdersByShoppingCartID(ctx, shoppingCart.ShoppingCartID)

	var userOrdersResponse orderResponse
	userOrdersResponse.Username = user.Username
	userOrdersResponse.Email = user.Email
	userOrdersResponse.PhoneNumber = user.PhoneNumber

	var allOrders []pizzasOrderResponse
	var totalCost float64

	for _, order := range orders {
		orderDetails := pizzasOrderResponse{
			PizzaName: order.PizzaName,
			Quantity:  order.Quantity,
		}
		allOrders = append(allOrders, orderDetails)
		// Calculate total cost - modify this based on how you store the pizza prices in orders
		totalCost += order.PizzaPrice * float64(order.Quantity)
	}
	userOrdersResponse.Orders = allOrders
	userOrdersResponse.TotalCost = totalCost

	_, err = server.store.CreateShoppingCart(ctx, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error at creating shopping cart"})
		return
	}
	ctx.JSON(http.StatusOK, userOrdersResponse)
}
