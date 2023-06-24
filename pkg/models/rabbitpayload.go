package models

type GetAccountDetails struct {
	XCorrelationID string `json:"xcorrelationid"`
	AccountID      string `json:"accountID"`
}
