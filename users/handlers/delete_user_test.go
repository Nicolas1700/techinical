package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	sentencesMock "techinical/test/mocks/sentences"
	"techinical/users/handlers"
	"techinical/users/infrastructura/dto"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteUser(t *testing.T) {
	app := fiber.New()
	mockSentente := new(sentencesMock.MockSentences)
	handler := handlers.NewHandlerDeleteUser(mockSentente)
	app.Delete("/test-user", handler.DeleteUser)

	mockSentente.On("DeleteRecord", mock.Anything, "users", "Id_User", "test-id-user", mock.Anything).
		Return(nil)

	user := dto.User{
		Id_User: "test-id-user",
	}

	userJSON, err := json.Marshal(user)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/test-user", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestDeleteUserFail(t *testing.T) {
	app := fiber.New()
	mockSentences := new(sentencesMock.MockSentences)
	h := handlers.NewHandlerDeleteUser(mockSentences)
	app.Delete("/test-user", h.DeleteUser)

	mockSentences.On("DeleteRecord", mock.Anything, "users", "Id_User", mock.Anything, mock.Anything).
		Return(errors.New("delete failed"))

	user := dto.User{
		Id_User: "test-id-user",
	}
	userJSON, err := json.Marshal(user)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/test-user", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.Nil(t, err)


	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	mockSentences.AssertExpectations(t)
}
