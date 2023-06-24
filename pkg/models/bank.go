package models

//Bank ...
type Bank struct {
	BankUUID   string `json:"bankuuid"`
	BankName   string `json:"bankname"`
	IFSCCode   string `json:"ifsccode"`
	BranchName string `json:"branchname"`
}
