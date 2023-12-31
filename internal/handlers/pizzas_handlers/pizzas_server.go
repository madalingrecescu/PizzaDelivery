package pizzas_handlers

import (
	"github.com/gin-gonic/gin"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_pizzas"
	"github.com/madalingrecescu/PizzaDelivery/internal/token"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/pizzas", server.getAllPizzas)
	router.GET("/pizzas/id/:id", server.getPizzaById)
	router.GET("/pizzas/name/:name", server.getPizzaByName)
	router.PUT("/pizzas/name/:name", server.updatePizza)
	router.POST("/pizzas", server.createPizza)
	router.DELETE("/pizzas/id/:id", server.deletePizza)

	router.POST("/addToOrder", server.addPizzaToShoppingCart)
	router.DELETE("/deletePizzaOrder", server.deletePizzaFromShoppingCart)
	router.PUT("/changeOrderQuantity", server.changeQuantityOfPizzas)

	router.POST("/shoppingCart/:id", server.createShoppingCart)

	router.POST("/order", server.order)

	server.router = router

}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
