package domain

import "context"

type (
	Database interface {
		InsertCustomer(ctx context.Context, customer Customer) error
		GetCreditLimitByCustomerId(ctx context.Context, customerId int64) (CreditLimit, error)
		InsertTransaction(ctx context.Context, transaction Transaction) error
		InsertContract(ctx context.Context, contract Contract) error
	}

	Usecase interface {
		CreateCustomer(ctx context.Context, customer Customer) (Customer, error)
		CreateCreditLimit(ctx context.Context, credit CreditLimit) (CreditLimit, error)
		CreateTransaction(ctx context.Context, transaction Transaction) (Transaction, error)
		CreateContract(ctx context.Context, contract Contract) (Contract, error)
	}
)

type (
	GetResponse struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

type (
	Customer struct {
		CustomerId int64  `json:"customerid"`
		NIK        string `json:"nik"`
		FullName   string `json:"fullname"`
		LegalName  string `json:"legalname"`
		BirthPlace string `json:"birthplace"`
		BirthDate  string `json:"birthdate"`
		Salary     int64  `json:"salary"`
		PhotoKTP   string `json:"photoktp"`
		PhotoSelfie string `json:"photoselfie"`
	}
	CreditLimit struct {
		LimitID     int64   `json:"limitid"`
		CustomerID  int64   `json:"customerid"`
		Tenor       int     `json:"tenor"`
		LimitAmount float64 `json:"limitamount"`
	}
	Transaction struct {
		TransactionID   int64   `json:"transactionid"`
		CustomerID      int64   `json:"customerid"`
		TransactionDate string  `json:"transactiondate"`
		OTR             float64 `json:"otr"`
		AdminFee        float64 `json:"adminfee"`
		InstallmentAmount float64 `json:"installmentamount"`
		InterestAmount  float64 `json:"interestamount"`
		AssetName       string  `json:"assetname"`
	}
	Contract struct {
		ContractNumber string `json:"contractnumber"`
		ContractDate   string `json:"contractdate"`
		TransactionID   int64   `json:"transactionid"`
		Terms          string `json:"terms"`
		Status         string `json:"status"`
	}
)