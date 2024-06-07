package usecase

import (
	"context"
	"kredit-plus/src/domain"
	"time"
)

type creditUsecase struct {
	database domain.Database
}

func New(database domain.Database) domain.Usecase {
	return &creditUsecase{
		database: database,
	}
}

func (c *creditUsecase) CreateCustomer(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	err := c.database.InsertCustomer(ctx, customer)
	if err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

func (c *creditUsecase) CreateCreditLimit(ctx context.Context, credit domain.CreditLimit) (domain.CreditLimit, error) {
	return domain.CreditLimit{}, nil
}

func (c *creditUsecase) CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	creditLimit, err := c.database.GetCreditLimitByCustomerId(ctx, transaction.CustomerID)
	if err != nil {
		return domain.Transaction{}, err
	}

	if transaction.OTR > creditLimit.LimitAmount {
		return domain.Transaction{}, nil
	}

	transaction.TransactionDate = time.Now().Format("2006-01-02")
	err = c.database.InsertTransaction(ctx, transaction)
	if err != nil {
		return domain.Transaction{}, err
	}
	
	contract := domain.Contract{
		ContractNumber: generateUniqueContractNumber(),
		TransactionID: transaction.TransactionID,
		ContractDate: time.Now().Format("2006-01-02"),
		Terms: "Lorem Ipsum Dolor Sit",
		Status: "active",
	}

	err = c.database.InsertContract(ctx, contract)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (c *creditUsecase) CreateContract(ctx context.Context, contract domain.Contract) (domain.Contract, error) {
	return domain.Contract{}, nil
}


