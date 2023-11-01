package user_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users"
	"github.com/madalingrecescu/PizzaDelivery/internal/token"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
)

// Server serves HTTP requests for our users service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/signup", server.createAccount)
	router.GET("/user/:id", server.getAccount)
	router.POST("/login", server.loginUser)

	server.router = router

}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
