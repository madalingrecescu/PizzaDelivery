package gRPC_handlers

import (
	"fmt"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users"
	"github.com/madalingrecescu/PizzaDelivery/internal/pb"
	"github.com/madalingrecescu/PizzaDelivery/internal/token"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
)

// Server serves gRPC requests for our users service
type Server struct {
	pb.UnimplementedPizzeriaServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server and setup routing
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

	return server, nil
}
