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

func Test_UpdateBankDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"bankname":"hdfc",
		"ifsccode":"HDFC00008017",
		"branchname":"stonehousepet"
		}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(&models.Bank{
		BankID:   1,
		BankUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)
	sqlServiceMock.EXPECT().PatchBankDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateBankDetails(response, request)
	assert.Equal(t, response.Code, 204)
}

func Test_UpdateBankDetails_InvalidBody_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"bankname":"hdfc",
		"ifsccode":"HDFC00008017"
		"branchname":"stonehousepet"
		}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(&models.Bank{
		BankID:   1,
		BankUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)
	sqlServiceMock.EXPECT().PatchBankDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateBankDetails(response, request)
	assert.Equal(t, response.Code, 500)
}

func Test_UpdateBankDetails_IFSCMissing_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"bankname":"hdfc",
		"branchname":"stonehousepet"
		}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(&models.Bank{
		BankID:   1,
		BankUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateBankDetails(response, request)
	assert.Equal(t, response.Code, 400)
}

func Test_UpdateBankDetails_UpdateDBCallFail_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"bankname":"hdfc",
		"ifsccode":"HDFC00008017",
		"branchname":"stonehousepet"
		}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(&models.Bank{
		BankID:   1,
		BankUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)
	sqlServiceMock.EXPECT().PatchBankDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(errors.New("error"))

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateBankDetails(response, request)
	assert.Equal(t, response.Code, 500)
}

func Test_UpdateBankDetails_InvalidBankUUID_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"bankname":"hdfc",
		"ifsccode":"HDFC00008017",
		"branchname":"stonehousepet"
		}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(nil, database.NoRowError)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateBankDetails(response, request)
	assert.Equal(t, response.Code, 400)
}

func Test_UpdateBankDetails_GetbankDetailDBCallFail_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestBody := `{
		"bankname":"hdfc",
		"ifsccode":"HDFC00008017",
		"branchname":"stonehousepet"
		}`

	request, _ := http.NewRequest("PUT", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", strings.NewReader(requestBody))
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(nil, errors.New("error"))

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.UpdateBankDetails(response, request)
	assert.Equal(t, response.Code, 500)
}
