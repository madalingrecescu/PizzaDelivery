package pizzas_handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_pizzas"
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
		UserID int32 `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shoppingCart, err := server.store.CreateShoppingCart(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, shoppingCart)
}

type deletePizzaFromShoppingCartRequest struct {
	ShoppingCartId int32  `json:"shopping_cart_id" binding:"required"`
	PizzaName      string `json:"name" binding:"required"`
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
	PizzaName string `uri:"name" binding:"required"`
}

func (server *Server) addPizzaToShoppingCart(ctx *gin.Context) {
	var req addPizzaToShoppingCartRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind the uri"})
		return
	}

	fmt.Println(req.PizzaName)
	pizza, err := server.store.GetPizzaByName(ctx, req.PizzaName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Pizza not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	fmt.Println(req.PizzaName)

	//TODO Hardcoded for example - replace this logic with actual userID retrieval
	shoppingCartId := int32(1) // TODO Replace with real logic

	arg := db.GetPizzaOrderByNameFromShoppingCartParams{
		ShoppingCartID: shoppingCartId,
		PizzaName:      pizza.Name,
	}
	pizzaOrder, err := server.store.GetPizzaOrderByNameFromShoppingCart(ctx, arg)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find the order"})
			return
		}
	}

	fmt.Println(pizzaOrder.PizzaOrderID, " id")
	if pizzaOrder.PizzaOrderID == 0 {
		_, err = server.store.CreatePizzaOrder(ctx, db.CreatePizzaOrderParams{
			ShoppingCartID: shoppingCartId,
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
			ShoppingCartID: shoppingCartId,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add to quantity"})
			return
		}
	}
	fmt.Println("4")

	ctx.JSON(http.StatusOK, gin.H{"message": "Pizza added to the shopping cart"})
}

func (server *Server) changeQuantityOfPizzas(ctx *gin.Context) {
	//todo implement
}
