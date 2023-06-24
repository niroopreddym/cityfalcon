package models

//Bank ...
type Bank struct {
	BankID     int    `json:"bankid"`
	BankUUID   string `json:"bankuuid"`
	BankName   string `json:"bankname"`
	IFSCCode   string `json:"ifsccode"`
	BranchName string `json:"branchname"`
}
