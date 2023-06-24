package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/niroopreddym/cityfalcon/pkg/models"
)

func postAccountRequestBodyInitialValidation(accountDetails models.Account, errorMessages *[]string) {
	if strings.TrimSpace(accountDetails.BankUUID) == "" {
		errorMessage := "Attribute Missing: BankUUID in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}

//---------------------------------Account related endpoints-------------------------------------------------

// CreateAccount creates a account in bank
func (handler *BankAndAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	accountDetails := models.Account{}
	log.Println("request:", r)
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err := json.Unmarshal([]byte(strBufferValue), &accountDetails)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	errorMessages := []string{}
	postAccountRequestBodyInitialValidation(accountDetails, &errorMessages)
	if len(errorMessages) > 0 {
		responseController(w, http.StatusBadRequest, errorMessages)
		return
	}

	uniqueID, err := handler.DatabaseService.PostAccountDetails(&accountDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"AccountID": *uniqueID,
	})
}

// GetAccountDetails gets account details
func (handler *BankAndAccountHandler) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID := params["id"]

	time.Sleep(5 * time.Second)
	accDetails, err := handler.DatabaseService.GetAccountDetails(accountID)
	if accDetails.AccountID == nil {
		responseController(w, http.StatusNotFound, "Account Not Found")
		return
	}

	if err != nil {
		responseController(w, http.StatusInternalServerError, "Error occured while fetching the bank details")
		return
	}

	responseController(w, http.StatusOK, accDetails)
}
