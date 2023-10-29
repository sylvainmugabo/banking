package domain

import (
	"github.com/sylvainmugabo/microservices-lib/errs"
	"github.com/sylvainmugabo/microservices/banking/dto"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	DateofBirth string `db:"date_of_birth"`
	City        string
	ZipCode     string
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	Find(string) (*Customer, *errs.AppError)
}

func (c Customer) StatusAsText() string {
	statusText := "active"
	if c.Status == "0" {
		statusText = "inactive"
	}
	return statusText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		DateofBirth: c.DateofBirth,
		City:        c.City,
		ZipCode:     c.ZipCode,
		Status:      c.StatusAsText(),
	}
}
