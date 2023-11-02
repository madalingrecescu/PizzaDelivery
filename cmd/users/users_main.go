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
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriverUsers, config.DBSourceUsers)
	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := user_handlers.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	err = server.Start(config.UsersServerAddress)
	if err != nil {
		log.Fatalln("cannot start server: ", err)
	}
}
