package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/niroopreddym/cityfalcon/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_PostAccountDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka",
		"balance": 9256.89
	}`

	request, _ := http.NewRequest("POST", "/api/account", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	accID := "dcb50ac8-3420-49ca-9980-7986a3b6d5b8"
	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().PostAccountDetails(gomock.Any()).AnyTimes().Return(&accID, nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.CreateAccount(response, request)
	assert.Equal(t, response.Code, 200)
}

func Test_PostAccountDetails_InvalidBody_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := ``

	request, _ := http.NewRequest("POST", "/api/account", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  nil,
		Redis:            nil,
	}

	handler.CreateAccount(response, request)
	assert.Equal(t, 500, response.Code)
}

func Test_PostAccountDetails_MissingBankID_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"firstname":"yamuna",
		"lastname":"karuka",
		"balance": 9256.89
	}`

	request, _ := http.NewRequest("POST", "/api/account", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  nil,
		Redis:            nil,
	}

	handler.CreateAccount(response, request)
	assert.Equal(t, response.Code, 400)
}

func Test_PostAccountDetails_DBCallFail_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka",
		"balance": 9256.89
	}`

	request, _ := http.NewRequest("POST", "/api/account", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().PostAccountDetails(gomock.Any()).AnyTimes().Return(nil, errors.New("error"))

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.CreateAccount(response, request)
	assert.Equal(t, response.Code, 500)
}
