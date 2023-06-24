package models

//Account ...
type Account struct {
	AccountID         *int   `json:"accountid"`
	AccountHolderName string `json:"accountholdername"`
	BankUUID          string `json:"bankguid"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	Balance           string `json:"balance"`
}
