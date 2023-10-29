package services

import (
	"time"

	"github.com/sylvainmugabo/microservices-lib/errs"
	"github.com/sylvainmugabo/microservices/banking/domain"
	"github.com/sylvainmugabo/microservices/banking/dto"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(acc dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(acc dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {
	err := acc.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  acc.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: acc.AccountType,
		Amount:      acc.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)

	if err != nil {
		return nil, errs.NewUnexpectedError("Unable to save an account")
	}

	response := newAccount.ToNewAccountResponseDto()
	return &response, nil

}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
