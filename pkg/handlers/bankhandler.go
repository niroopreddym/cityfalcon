package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/niroopreddym/cityfalcon/pkg/database"
	"github.com/niroopreddym/cityfalcon/pkg/models"
	"github.com/niroopreddym/cityfalcon/pkg/rabbitmq"
	"github.com/niroopreddym/cityfalcon/pkg/services"
)

// BankAndAccountHandler is the class implementation for CompositeIface Interface
type BankAndAccountHandler struct {
	RMQEventsService *rabbitmq.RabbitEvents
	DatabaseService  services.ISQLService
	Redis            *services.RedisService
}

// NewBankAndAccountsHandlerInstance instantiates the struct
func NewBankAndAccountsHandlerInstance() CompositeIface {
	rabbitConnection, err := rabbitmq.NewConnection()
	if err != nil {
		fmt.Println(err)
	}

	return &BankAndAccountHandler{
		RMQEventsService: rabbitConnection,
		DatabaseService:  services.NewDatabaseServicesInstance(),
		Redis:            services.NewRedisService(),
	}
}

//---------------------------------Bank related endpoints-------------------------------------------------

// CreateBank creates a bank in DB
func (handler *BankAndAccountHandler) CreateBank(w http.ResponseWriter, r *http.Request) {
	bankDetails := models.Bank{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err := json.Unmarshal([]byte(strBufferValue), &bankDetails)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	errorMessages := []string{}
	postRequestBodyInitialValidation(bankDetails, &errorMessages)
	if len(errorMessages) > 0 {
		errorResponse := map[string][]string{
			"errormessages": errorMessages,
		}

		responseController(w, http.StatusBadRequest, errorResponse)
		return
	}

	uniqueID, err := handler.DatabaseService.PostBankDetails(&bankDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"BankUUID": *uniqueID,
	})
}

// GetAllBanks gets all banks from DB
func (handler *BankAndAccountHandler) GetAllBanks(w http.ResponseWriter, r *http.Request) {
	lstBanks, err := handler.DatabaseService.ListAllBanks()
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, "Error occured while fetching the bank details"+err.Error())
		return
	}

	responseController(w, http.StatusOK, lstBanks)
}

// GetBankDetails gets the bank details
func (handler *BankAndAccountHandler) GetBankDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bankUUID := params["uuid"]

	bankDetails, err := handler.DatabaseService.GetBankDetails(bankUUID)
	if err != nil {
		responseController(w, http.StatusInternalServerError, "Error occured while fetching the bank details")
		return
	}

	responseController(w, http.StatusOK, bankDetails)
}

// UpdateBankDetails updates the bank details
func (handler *BankAndAccountHandler) UpdateBankDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bankUUID := params["uuid"]

	_, err := handler.DatabaseService.GetBankDetails(bankUUID)
	if err != nil {
		if err == database.NoRowError {
			responseController(w, http.StatusBadRequest, fmt.Sprintf("bank with bank_uuid %v is not found", bankUUID))
			return
		}

		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	requestedBankDetails := models.Bank{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Println("reading from buffer")
		err := errors.New("error reading data from the response " + readErr.Error())
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	strBufferValue := string(bodyBytes)
	err = json.Unmarshal([]byte(strBufferValue), &requestedBankDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorMessages := []string{}
	patchCallInitialValidation(requestedBankDetails, &errorMessages)
	if len(errorMessages) > 0 {
		responseController(w, http.StatusBadRequest, errorMessages)
		return
	}

	err = handler.DatabaseService.PatchBankDetails(bankUUID, requestedBankDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusNoContent, "Updated Sucessfully")
}

// RemoveBank deletes the bank record
func (handler *BankAndAccountHandler) RemoveBank(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bankUUID := params["uuid"]

	_, err := handler.DatabaseService.GetBankDetails(bankUUID)
	if err != nil {
		if err == database.NoRowError {
			responseController(w, http.StatusBadRequest, fmt.Sprintf("bank with bank_uuid %v is not found", bankUUID))
			return
		}

		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = handler.DatabaseService.DeleteBank(bankUUID)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err.Error())
	}

	responseController(w, http.StatusNoContent, "Successfully Deleted")
}

func postRequestBodyInitialValidation(bankDetails models.Bank, errorMessages *[]string) {
	if strings.TrimSpace(bankDetails.BankName) == "" {
		errorMessage := "Attribute Missing: Name in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}

	if strings.TrimSpace(bankDetails.IFSCCode) == "" {
		errorMessage := "Attribute Missing: IFSCCode in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}

	if strings.TrimSpace(bankDetails.BranchName) == "" {
		errorMessage := "Attribute Missing: BranchName in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}

func responseController(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func patchCallInitialValidation(bankDetails models.Bank, errorMessages *[]string) {
	if bankDetails.IFSCCode == "" {
		errorMessage := "Invalid Attribute: IFSCCode in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}
