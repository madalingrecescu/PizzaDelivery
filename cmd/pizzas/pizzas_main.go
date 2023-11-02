package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_pizzas"
	"github.com/madalingrecescu/PizzaDelivery/internal/handlers/pizzas_handlers"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
	"log"
)

func main() {
	config, err := util.LoadConfig("internal")
	if err != nil {
		log.Fatal("cannot load config pizzas: ", err)
	}
	conn, err := sql.Open(config.DBDriverPizzas, config.DBSourcePizzas)
	if err != nil {
		log.Fatalln("cannot connect to db pizzas: ", err)
	}

	store := db.NewStore(conn)
	server, err := pizzas_handlers.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server pizzas: ", err)
	}

	err = server.Start(config.PizzasServerAddress)
	if err != nil {
		log.Fatalln("cannot start server pizzas: ", err)
	}
}
