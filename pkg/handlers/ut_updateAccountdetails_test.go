package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/niroopreddym/cityfalcon/pkg/database"
	mock "github.com/niroopreddym/cityfalcon/pkg/mocks"
	"github.com/niroopreddym/cityfalcon/pkg/models"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateAccountDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka",
		"balance": 9256.89
	}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetAccountDetails(gomock.Any()).AnyTimes().Return(&models.Account{
		AccountID:   1,
		AccountUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)
	sqlServiceMock.EXPECT().UpdateAccountDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	redisMock := mock.NewMockIRedisService(controller)
	redisMock.EXPECT().AddKey(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            redisMock,
	}

	handler.UpdateAccountDetails(response, request)
	assert.Equal(t, response.Code, 204)
}

func Test_UpdateAccountDetails_BalanceMissingInBody_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka"
	}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  nil,
		Redis:            nil,
	}

	handler.UpdateAccountDetails(response, request)
	assert.Equal(t, response.Code, 400)
}

func Test_UpdateAccountDetails_InvalidBody_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka",
	}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  nil,
		Redis:            nil,
	}

	handler.UpdateAccountDetails(response, request)
	assert.Equal(t, response.Code, 500)
}

func Test_UpdateAccountDetails_RedisCallFail_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka",
		"balance": 9256.89
	}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetAccountDetails(gomock.Any()).AnyTimes().Return(&models.Account{
		AccountID:   1,
		AccountUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)
	sqlServiceMock.EXPECT().UpdateAccountDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	redisMock := mock.NewMockIRedisService(controller)
	redisMock.EXPECT().AddKey(gomock.Any(), gomock.Any()).AnyTimes().Return(errors.New("error"))

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            redisMock,
	}

	handler.UpdateAccountDetails(response, request)
	assert.Equal(t, response.Code, 206)
}

func Test_UpdateAccountDetails_UpdateCallFail_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka",
		"balance": 9256.89
	}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetAccountDetails(gomock.Any()).AnyTimes().Return(&models.Account{
		AccountID:   1,
		AccountUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)
	sqlServiceMock.EXPECT().UpdateAccountDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(errors.New("error"))

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateAccountDetails(response, request)
	assert.Equal(t, response.Code, 500)
}

func Test_UpdateAccountDetails_AccounDetailsDBFail_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka",
		"balance": 9256.89
	}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetAccountDetails(gomock.Any()).AnyTimes().Return(nil, errors.New("error"))

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateAccountDetails(response, request)
	assert.Equal(t, response.Code, 500)
}

func Test_UpdateAccountDetails_AccounDetailsNotThere_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"accountholdername":"Niroop",
		"bankid":1,
		"firstname":"yamuna",
		"lastname":"karuka",
		"balance": 9256.89
	}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetAccountDetails(gomock.Any()).AnyTimes().Return(nil, database.NoRowError)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateAccountDetails(response, request)
	assert.Equal(t, response.Code, 400)
}
