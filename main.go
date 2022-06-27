package main

import (
	"enigmacamp.com/go-db-fundamnetal/config"
	"enigmacamp.com/go-db-fundamnetal/sample"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

func main() {
	config := config.NewConfigDB()
	db := config.DbConn()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			utils.IsError(err)
		}
	}(db)

	sample.ShopRun(db)
	//sample.CustomerRun(db)
}

/**
# Fetching
-> Get All (x)
-> Get By Id (x)
-> Get By Name (x)
-> Aggregate:
   -> count
   -> sum
   -> etc...
-> Join
*/
