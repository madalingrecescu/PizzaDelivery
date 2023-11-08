package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users"
	"github.com/madalingrecescu/PizzaDelivery/internal/handlers/user_handlers"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
	"log"
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

	runGinServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := user_handlers.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server users: ", err)
	}

	err = server.Start(config.UsersServerAddress)
	if err != nil {
		log.Fatalln("cannot start server users: ", err)
	}
	log.Printf("starting users server at %s", config.UsersServerAddress)
}
