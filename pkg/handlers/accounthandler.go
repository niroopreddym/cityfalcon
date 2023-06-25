package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niroopreddym/cityfalcon/pkg/database"
	"github.com/niroopreddym/cityfalcon/pkg/models"
)

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
		"AccountUUID": *uniqueID,
	})
}

// GetAccountDetails gets account details
func (handler *BankAndAccountHandler) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID := params["uuid"]

	getAccDetailsRabbitPayload := models.GetAccountDetails{
		XCorrelationID: accountID,
		AccountID:      accountID,
	}

	byteArr, err := json.Marshal(getAccDetailsRabbitPayload)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = handler.RMQEventsService.PublishMessage(byteArr)
	if err != nil {
		responseController(w, http.StatusInternalServerError, "Fail to send data to RMQ producer "+err.Error())
		return
	}

	partialResponse := map[string]string{
		"trackingURL": "/account/getaccountdetails/asyncresponse/" + accountID,
	}

	responseController(w, http.StatusPartialContent, partialResponse)
}

func (handler *BankAndAccountHandler) GetAccountDetailsResponse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	correlationID := params["uuid"]

	//read thorugh cache type implementation
	accDetails, err := handler.Redis.ReadKey(correlationID)
	if err != nil {
		fmt.Println(err)
	} else {
		responseController(w, http.StatusOK, accDetails)
		return
	}

	accDetails, err = handler.Redis.ReadKey(correlationID)
	if err != nil {
		responseController(w, http.StatusPartialContent, "response is in progress")
		return
	}

	responseController(w, http.StatusOK, accDetails)
}

// UpdateAccountDetails updates the balance in account using write through cache
func (handler *BankAndAccountHandler) UpdateAccountDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountUUID := params["uuid"]

	_, err := handler.DatabaseService.GetAccountDetails(accountUUID)
	if err != nil {
		if err == database.NoRowError {
			responseController(w, http.StatusBadRequest, fmt.Sprintf("bank with bank_uuid %v is not found", accountUUID))
			return
		}

		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	requestedAccDetails := models.Account{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Println("reading from buffer")
		err := errors.New("error reading data from the response " + readErr.Error())
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	strBufferValue := string(bodyBytes)
	err = json.Unmarshal([]byte(strBufferValue), &requestedAccDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorMessages := []string{}
	updateAccountRequestBodyInitialValidation(requestedAccDetails, &errorMessages)
	if len(errorMessages) > 0 {
		responseController(w, http.StatusBadRequest, errorMessages)
		return
	}

	err = handler.DatabaseService.UpdateAccountDetails(accountUUID, requestedAccDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = handler.Redis.AddKey(accountUUID, fmt.Sprintf("%v", *requestedAccDetails.Balance))
	if err != nil {
		responseController(w, http.StatusPartialContent, "response is in progress")
		return
	}

	responseController(w, http.StatusNoContent, "Updated Sucessfully")
}

func postAccountRequestBodyInitialValidation(accountDetails models.Account, errorMessages *[]string) {
	if accountDetails.BankId == nil {
		errorMessage := "Attribute Missing: BankId in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}

func updateAccountRequestBodyInitialValidation(accountDetails models.Account, errorMessages *[]string) {
	if accountDetails.Balance == nil {
		errorMessage := "Attribute Missing: Balance in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}
