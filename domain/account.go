package domain

import "github.com/sylvainmugabo/microservices-lib/errs"

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	AccountType string `db:"account_type"`
	Amount      float64
	OpeningDate string `db:"opening_date"`
	Status      string
}

type AccountRepository interface {
	Save(acc Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(id string) (*Account, *errs.AppError)
}

func (acc Account) CanWithdraw(amount float64) bool {
	return acc.Amount >= amount
}
