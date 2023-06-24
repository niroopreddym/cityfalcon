package services

import "github.com/niroopreddym/cityfalcon/pkg/models"

//ISQLService provides the data interface
type ISQLService interface {
	PostBankDetails(bankDetails *models.Bank) (*string, error)
	ListAllBanks() ([]*models.Bank, error)
	GetBankDetails(uuid string) (*models.Bank, error)
	PatchBankDetails(id string, bankDetails models.Bank) error
	DeleteBank(id string) error

	PostAccountDetails(bankDetails *models.Account) (*string, error)
	GetAccountDetails(id string) (*models.Account, error)
}
