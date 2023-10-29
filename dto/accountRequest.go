package dto

import "github.com/sylvainmugabo/microservices-lib/errs"

type AccountRequest struct {
	CustomerId  string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Amount      float64
}

func (acc AccountRequest) Validate() *errs.AppError {
	if acc.Amount > 5000 {
		return errs.NewValidationError("Exceeded balance")
	}
	if acc.AccountType != "saving" && acc.AccountType != "checking" {
		return errs.NewValidationError("Only saving or checking accounts are allowed")
	}

	return nil

}
