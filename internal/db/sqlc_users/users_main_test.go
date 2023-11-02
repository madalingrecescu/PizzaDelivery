package users_db

import (
	"database/sql"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}
	testDB, err = sql.Open(config.DBDriverUsers, config.DBSourceUsers)
	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
