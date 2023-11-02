package pizzas_handlers

import (
	"database/sql"
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
