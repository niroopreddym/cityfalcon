package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/niroopreddym/cityfalcon/pkg/database"
	mock "github.com/niroopreddym/cityfalcon/pkg/mocks"
	"github.com/niroopreddym/cityfalcon/pkg/models"
	"github.com/stretchr/testify/assert"
)

func Test_RemoveBankDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request, _ := http.NewRequest("DELETE", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", nil)
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(&models.Bank{
		BankID:   1,
		BankUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)

	sqlServiceMock.EXPECT().DeleteBank(gomock.Any()).AnyTimes().Return(nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.RemoveBank(response, request)
	assert.Equal(t, response.Code, 204)
}

func Test_RemoveBankDetails_DBFail_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request, _ := http.NewRequest("DELETE", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", nil)
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(&models.Bank{
		BankID:   1,
		BankUUID: "dcb50ac8-3420-49ca-9980-7986a3b6d5b8",
	}, nil)

	sqlServiceMock.EXPECT().DeleteBank(gomock.Any()).AnyTimes().Return(errors.New("error"))

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.RemoveBank(response, request)
	assert.Equal(t, response.Code, 500)
}

func Test_RemoveBankDetails_InvalidBankID_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request, _ := http.NewRequest("DELETE", "/api/bank/{dcb50ac8-3420-49ca-9980-7986a3b6d5b8}", nil)
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().GetBankDetails(gomock.Any()).AnyTimes().Return(nil, database.NoRowError)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.RemoveBank(response, request)
	assert.Equal(t, response.Code, 400)
}
