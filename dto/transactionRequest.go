package dto

import "github.com/sylvainmugabo/microservices-lib/errs"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

func (r TransactionRequest) Validate() *errs.AppError {

	if r.TransactionType == WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal ")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero ")
	}

	return nil

}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}
