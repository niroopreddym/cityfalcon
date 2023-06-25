package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/niroopreddym/cityfalcon/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_GetAccountDetailsResponse(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request, _ := http.NewRequest("GET", "/api/account/getaccountdetails/asyncresponse/dcb50ac8-3420-49ca-9980-7986a3b6d5b8", nil)
	response := httptest.NewRecorder()

	redisServiceMock := mock.NewMockIRedisService(controller)
	res := "123"
	redisServiceMock.EXPECT().ReadKey(gomock.Any()).AnyTimes().Return(&res, nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  nil,
		Redis:            redisServiceMock,
	}

	handler.GetAccountDetailsResponse(response, request)
	assert.Equal(t, response.Code, 200)
}
