package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"techinical/challenges/handlers"
	"techinical/challenges/infrastructura/dto"
	dbMock "techinical/test/mocks/db"
	sentencesMock "techinical/test/mocks/sentences"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetChallenge_Success(t *testing.T) {
	app := fiber.New()
	mockSentences := new(sentencesMock.MockSentences)
	mockDB := new(dbMock.MockDb) // Usamos el mock de db.DB
	handler := handlers.NewHandlerGetChallenge(mockSentences)
	app.Get("/test-challenge", handler.GetChallenge)

	query := "SELECT * FROM challenges"
	mockSentences.On("PaginateAndQuery", mock.Anything, "challenges").Return(query, nil)

	challenges := []dto.Challenge{
		{Id_Challenge: "1", Name_Challenge: "Challenge 1", Number_Participants: 10},
		{Id_Challenge: "2", Name_Challenge: "Challenge 2", Number_Participants: 20},
	}

	mockDB.On("Raw", query).Return(mockDB)
	mockDB.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*[]dto.Challenge)
		*dest = challenges
	}).Return(mockDB)

	// Simulamos una petición HTTP GET
	req := httptest.NewRequest(http.MethodGet, "/test-challenge", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)

	// Verificamos que el status code es 200 OK
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var responseChallenges []dto.Challenge
	err = json.NewDecoder(resp.Body).Decode(&responseChallenges)
	assert.Nil(t, err)
	assert.Equal(t, challenges, responseChallenges)
}
