package services

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/google/uuid"
	"github.com/niroopreddym/cityfalcon/pkg/database"
	"github.com/niroopreddym/cityfalcon/pkg/models"
)

// DatabaseService is the class implementation for ProductServicesIface interface
type DatabaseService struct {
	DatabaseService database.DbIface
}

// NewDatabaseServicesInstance instantiates the struct
func NewDatabaseServicesInstance() *DatabaseService {
	return &DatabaseService{
		DatabaseService: database.DBNewHandler(),
	}
}

// PostBankDetails posts the bank data from DB
func (service *DatabaseService) PostBankDetails(bankDetails *models.Bank) (*string, error) {
	defer service.DatabaseService.DbClose()

	uuid := uuid.New().String()
	query := `INSERT INTO Bank (bank_uuid, bank_name, ifsc_code, branch_name) VALUES ($1, $2, $3, $4)`
	data, err := service.DatabaseService.DbExecuteScalar(query, uuid, bankDetails.BankName, bankDetails.IFSCCode, bankDetails.BranchName)
	fmt.Println(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &uuid, nil
}

// ListAllBanks retrives all bank records from DB
func (service *DatabaseService) ListAllBanks() ([]*models.Bank, error) {
	defer service.DatabaseService.DbClose()
	// pagelimit := 1
	// query := "select * from Bank limit " + strconv.Itoa(pagelimit) + " offset " + strconv.Itoa(pagelimit*(pageNumber-1))
	query := "select * from Bank"

	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	txResult := []*models.Bank{}
	for rowsAffected.Next() {
		var id int
		var bankUUID string
		var bankName string
		var ifscCode string
		var branchName string

		if err := rowsAffected.Scan(&id, &bankUUID, &bankName, &ifscCode, &branchName); err != nil {
			log.Println(err)
			log.Fatal(err)
		}

		bank := &models.Bank{
			BankID:     id,
			BankUUID:   bankUUID,
			BankName:   bankName,
			IFSCCode:   ifscCode,
			BranchName: branchName,
		}

		txResult = append(txResult, bank)
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, err
	}

	return txResult, nil
}

// GetBankDetails fetches the bank details
func (service *DatabaseService) GetBankDetails(uuid string) (*models.Bank, error) {
	defer service.DatabaseService.DbClose()
	query := `select * from Bank where bank_uuid='` + uuid + `'`
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	txResult := &models.Bank{}
	for rowsAffected.Next() {
		var id int
		var bankUUID string
		var bankName string
		var ifscCode string
		var branchName string

		if err := rowsAffected.Scan(&id, &bankUUID, &bankName, &ifscCode, &branchName); err != nil {
			log.Println(err)
			return nil, err
		}

		txResult = &models.Bank{
			BankID:     id,
			BankUUID:   bankUUID,
			BankName:   bankName,
			IFSCCode:   ifscCode,
			BranchName: branchName,
		}

	}

	if txResult.BankUUID == "" {
		return nil, database.NoRowError
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, err
	}

	return txResult, nil
}

// PatchBankDetails updates the bank details
func (service *DatabaseService) PatchBankDetails(id string, bankDetails models.Bank) error {
	defer service.DatabaseService.DbClose()
	query := fmt.Sprintf("update Bank set bank_name = '%v', branch_name='%v' where bank_uuid='%v'", bankDetails.BankName, bankDetails.BranchName, id)

	tx, err := service.DatabaseService.TxBegin()
	_, err = service.DatabaseService.TxExecuteStmt(tx, query)

	if err != nil {
		log.Println(err)
		return err
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return err
	}

	return nil
}

// DeleteBank deletes the bank record
func (service *DatabaseService) DeleteBank(id string) error {
	defer service.DatabaseService.DbClose()
	query := fmt.Sprintf("delete from Bank where bank_uuid='%v'", id)

	tx, err := service.DatabaseService.TxBegin()
	_, err = service.DatabaseService.TxExecuteStmt(tx, query)

	if err != nil {
		log.Println(err)
		return err
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return err
	}

	return nil
}

// PostAccountDetails creates account data record
func (service *DatabaseService) PostAccountDetails(accountDetails *models.Account) (*string, error) {
	defer service.DatabaseService.DbClose()

	uuid := uuid.New().String()

	query := `INSERT INTO Account (account_uuid, account_holder_name, bank_id, first_name, last_name, balance) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := service.DatabaseService.DbExecuteScalar(query, uuid, accountDetails.AccountHolderName, accountDetails.BankId,
		accountDetails.FirstName, accountDetails.LastName, fmt.Sprintf("%v", *accountDetails.Balance))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &uuid, nil
}

// GetAccountDetails fetches acc details
func (service *DatabaseService) GetAccountDetails(id string) (*models.Account, error) {
	defer service.DatabaseService.DbClose()
	query := fmt.Sprintf("select * from Account where account_uuid= '%v'", id)
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	txResult := &models.Account{}
	for rowsAffected.Next() {
		var id int
		var accountUUID string
		var accountHolderName string
		var bankID int
		var firstName string
		var lastName string
		var balance string

		if err := rowsAffected.Scan(&id, &accountUUID, &accountHolderName, &bankID, &firstName, &lastName, &balance); err != nil {
			log.Println(err)
			return nil, err
		}

		fmtBalance, err := strconv.ParseFloat(balance, 32)
		if err != nil {
			fmt.Println(err.Error())
		}

		roundedOffBalance := math.Ceil(fmtBalance*100) / 100
		txResult = &models.Account{
			AccountID:         id,
			AccountUUID:       accountUUID,
			AccountHolderName: accountHolderName,
			BankId:            &bankID,
			FirstName:         firstName,
			LastName:          lastName,
			Balance:           &roundedOffBalance,
		}
	}

	if txResult.AccountUUID == "" {
		return nil, database.NoRowError
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, err
	}

	return txResult, nil
}

// UpdateAccountDetails updates the account balance details
func (service *DatabaseService) UpdateAccountDetails(id string, accDetails models.Account) error {
	defer service.DatabaseService.DbClose()
	query := fmt.Sprintf("update Account set balance = '%v' where account_uuid='%v'", fmt.Sprintf("%v", *accDetails.Balance), id)

	tx, err := service.DatabaseService.TxBegin()
	_, err = service.DatabaseService.TxExecuteStmt(tx, query)

	if err != nil {
		log.Println(err)
		return err
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return err
	}

	return nil
}
