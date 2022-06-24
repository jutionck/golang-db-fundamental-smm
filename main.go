package main

import (
	"enigmacamp.com/go-db-fundamnetal/config"
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/repository"
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
	cstRepo := repository.NewCustomerRepository(db)

	// INSERT
	cstId := utils.GenerateId()
	customer := model.Customer{
		Id:      cstId,
		Name:    "Jution Candra Kirana",
		Address: "Ragunan",
		Phone:   "08292929",
		Email:   "jutionck@gmail.com",
		Balance: 150000,
	}
	cstRepo.Insert(&customer)

	// DELETE
	//customerId := "C004"
	//cstRepo.Delete(customerId)

	// UPDATE
	//customerUpdate := model.Customer{
	//	Name:    "Jution Aja",
	//	Address: "Ragunan",
	//	Phone:   "08292929",
	//	Email:   "jutionck@gmail.com",
	//	Id:      "1c63f19d-e896-4116-a847-2dbb75d7eae8",
	//}
	//cstRepo.Update(&customerUpdate)
}
