package user_handlers

import (
	"github.com/gin-gonic/gin"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users"
)

// Server serves HTTP requests for our users service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to router
	router.POST("/signup", server.createAccount)
	router.GET("/user/:id", server.getAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
