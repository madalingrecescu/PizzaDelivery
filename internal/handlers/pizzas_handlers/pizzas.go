package pizzas_handlers

import (
	"github.com/gin-gonic/gin"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_pizzas"
	"net/http"
)

type createPizzaRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func (server *Server) createPizza(ctx *gin.Context) {
	var req createPizzaRequest
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
	// TODO Implement me
}
func (server *Server) getPizzaById(ctx *gin.Context) {
	// TODO Implement me
}
func (server *Server) getPizzaByName(ctx *gin.Context) {
	// TODO Implement me
}
func (server *Server) updatePizza(ctx *gin.Context) {
	// TODO Implement me
}
func (server *Server) deletePizza(ctx *gin.Context) {
	// TODO Implement me
}
