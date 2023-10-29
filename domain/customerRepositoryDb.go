package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sylvainmugabo/microservices-lib/errs"
	"github.com/sylvainmugabo/microservices-lib/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var query string
	var err error
	customers := make([]Customer, 0)
	if len(status) == 0 {
		query = "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		err = d.client.Select(&customers, query)

	} else {

		query = "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers where status=?"
		err = d.client.Select(&customers, query, status)

	}

	if err != nil {

		logger.Error("Error while connecting to database " + err.Error())
		return nil, errs.NewUnexpectedError("Unable to connect to the database")
	}

	return customers, nil

}

func (d CustomerRepositoryDb) Find(id string) (*Customer, *errs.AppError) {
	var c Customer
	query := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers where customer_id=?"

	err := d.client.Get(&c, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database connection error")
	}

	return &c, nil

}

func NewCustomerRepositoryDb(dbclient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{client: dbclient}
}
