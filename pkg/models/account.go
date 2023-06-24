package models

//Account ...
type Account struct {
	AccountID         int     `json:"accountid"`
	AccountUUID       string  `json:"accountuuid"`
	AccountHolderName string  `json:"accountholdername"`
	BankId            *int    `json:"bankid"`
	FirstName         string  `json:"firstname"`
	LastName          string  `json:"lastname"`
	Balance           float64 `json:"balance"`
}
