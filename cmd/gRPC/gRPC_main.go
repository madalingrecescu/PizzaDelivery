package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users"
	"github.com/madalingrecescu/PizzaDelivery/internal/handlers/gRPC_handlers"
	"github.com/madalingrecescu/PizzaDelivery/internal/pb"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	config, err := util.LoadConfig("internal")
	if err != nil {
		log.Fatal("cannot load config users: ", err)
	}
	conn, err := sql.Open(config.DBDriverUsers, config.DBSourceUsers)
	if err != nil {
		log.Fatalln("cannot connect to db users: ", err)
	}

	store := db.NewStore(conn)

	runGrpcServer(config, store)

}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gRPC_handlers.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server gRPC: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPizzeriaServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener for grpc server")
	}

	log.Printf("start grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
