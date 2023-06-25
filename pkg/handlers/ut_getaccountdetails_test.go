package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/niroopreddym/cityfalcon/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_GetAccountDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request, _ := http.NewRequest("GET", "/api/account/dcb50ac8-3420-49ca-9980-7986a3b6d5b8", nil)
	response := httptest.NewRecorder()

	rmqMock := mock.NewMockIRMQService(controller)
	rmqMock.EXPECT().PublishMessage(gomock.Any()).AnyTimes().Return(nil)

	handler := BankAndAccountHandler{
		RMQEventsService: rmqMock,
		DatabaseService:  nil,
		Redis:            nil,
	}

	handler.GetAccountDetails(response, request)
	assert.Equal(t, response.Code, 206)
}

func Test_GetAccountDetails_PublishRMQFail_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request, _ := http.NewRequest("GET", "/api/account/dcb50ac8-3420-49ca-9980-7986a3b6d5b8", nil)
	response := httptest.NewRecorder()

	rmqMock := mock.NewMockIRMQService(controller)
	rmqMock.EXPECT().PublishMessage(gomock.Any()).AnyTimes().Return(errors.New("error"))

	handler := BankAndAccountHandler{
		RMQEventsService: rmqMock,
		DatabaseService:  nil,
		Redis:            nil,
	}

	handler.GetAccountDetails(response, request)
	assert.Equal(t, response.Code, 500)
}
