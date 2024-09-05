package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"techinical/challenges/handlers"
	"techinical/challenges/infrastructura/dto"
	sentencesMock "techinical/test/mocks/sentences"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteChallenge(t *testing.T) {
	app := fiber.New()
	mockSentente := new(sentencesMock.MockSentences)
	handler := handlers.NewHandlerDeleteChallenge(mockSentente)
	app.Delete("/test-challenge", handler.DeleteChallenge)

	mockSentente.On("DeleteRecord", mock.Anything, "challenges", "id_challenge", "test-id-challenge", mock.Anything).Return(nil)

	challenge := dto.Challenge{
		Id_Challenge: "test-id-challenge",
	}

	challengeJSON, err := json.Marshal(challenge)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/test-challenge", bytes.NewBuffer(challengeJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestDeleteChallengeFail(t *testing.T) {
	app := fiber.New()
	mockSentente := new(sentencesMock.MockSentences)
	handler := handlers.NewHandlerDeleteChallenge(mockSentente)
	app.Delete("/test-challenge", handler.DeleteChallenge)

	mockSentente.On("DeleteRecord", mock.Anything, "challenges", "id_challenge", "test-id-challenge", mock.Anything).Return(errors.New("error deleting record"))

	challenge := dto.Challenge{
		Id_Challenge: "test-id-challenge",
	}

	challengeJSON, err := json.Marshal(challenge)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/test-challenge", bytes.NewBuffer(challengeJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	mockSentente.AssertCalled(t, "DeleteRecord", mock.Anything, "challenges", "id_challenge", "test-id-challenge", mock.Anything)
}
