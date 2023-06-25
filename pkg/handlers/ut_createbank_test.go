package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/niroopreddym/cityfalcon/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_PostBankDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"bankname":"hdfc",
		"ifsccode":"HDFC00008017",
		"branchname":"stonehousepet"
		}`

	request, _ := http.NewRequest("POST", "/api/bank", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	bankID := "dcb50ac8-3420-49ca-9980-7986a3b6d5b8"
	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().PostBankDetails(gomock.Any()).AnyTimes().Return(&bankID, nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.CreateBank(response, request)
	assert.Equal(t, response.Code, 200)
}

func Test_CallConstructor(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	instance := NewBankAndAccountsHandlerInstance()
	assert.NotNil(t, instance)
	instance = nil
}
