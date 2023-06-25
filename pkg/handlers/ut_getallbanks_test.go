package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/niroopreddym/cityfalcon/pkg/mocks"
	"github.com/niroopreddym/cityfalcon/pkg/models"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllBanks(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request, _ := http.NewRequest("GET", "/api/bank", nil)
	response := httptest.NewRecorder()

	sqlServiceMock := mock.NewMockISQLService(controller)
	sqlServiceMock.EXPECT().ListAllBanks().AnyTimes().Return([]*models.Bank{{
		BankID: 1,
	}}, nil)

	handler := BankAndAccountHandler{
		RMQEventsService: nil,
		DatabaseService:  sqlServiceMock,
		Redis:            nil,
	}

	handler.GetAllBanks(response, request)
	assert.Equal(t, response.Code, 200)
}
