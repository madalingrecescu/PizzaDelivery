package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"pizzeria/internal/api"
	db "pizzeria/internal/sqlc_users"
	"pizzeria/internal/util"
)

func main() {
	config, err := util.LoadConfig("internal")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.UsersServerAdress)
	if err != nil {
		log.Fatalln("cannot start server: ", err)
	}
}
