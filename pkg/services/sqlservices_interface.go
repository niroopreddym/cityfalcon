package services

import "github.com/niroopreddym/cityfalcon/pkg/models"

//ISQLService provides the data interface
type ISQLService interface {
	IBankDetails
	IAccountDetails
}

//IBankDetails implements the DB calls for fetchig or amending Bank details
type IBankDetails interface {
	PostBankDetails(bankDetails *models.Bank) (*string, error)
	ListAllBanks() ([]*models.Bank, error)
	GetBankDetails(uuid string) (*models.Bank, error)
	PatchBankDetails(id string, bankDetails models.Bank) error
	DeleteBank(id string) error
}

//IAccountDetails implements the DB calls for fetchig or amending Account details
type IAccountDetails interface {
	PostAccountDetails(bankDetails *models.Account) (*string, error)
	GetAccountDetails(id string) (*models.Account, error)
	UpdateAccountDetails(id string, accDetails models.Account) error
}
