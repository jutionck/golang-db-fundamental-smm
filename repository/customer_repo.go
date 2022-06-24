package repository

import (
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	Insert(customer *model.Customer)
	Update(customer *model.Customer)
	Delete(id string)
}

type customerRepository struct {
	db *sqlx.DB
}

func (c *customerRepository) Insert(customer *model.Customer) {
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, customer)
	err = utils.IsError(err)
	if err != nil {
		return
	}
}

func (c *customerRepository) Update(customer *model.Customer) {
	_, err := c.db.NamedExec(utils.UPDATE_CUSTOMER, customer)
	err = utils.IsError(err)
	if err != nil {
		return
	}
}

func (c *customerRepository) Delete(id string) {
	c.db.MustExec(utils.DELETE_CUSTOMER, id)
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}
