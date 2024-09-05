package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	sentencesMock "techinical/test/mocks/sentences"
	"techinical/video/handlers"
	"techinical/video/infrastructura/dto"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteVideo(t *testing.T) {
	app := fiber.New()
	mockSentente := new(sentencesMock.MockSentences)
	handler := handlers.NewHandlerDeleteVideo(mockSentente)
	app.Delete("/test-video", handler.DeleteVideo)

	mockSentente.On("DeleteRecord", mock.Anything, "video", "id_video", "test-id-video", mock.Anything).
		Return(nil)

	video := dto.Video{
		Id_Video: "test-id-video",
	}

	videoJSON, err := json.Marshal(video)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/test-video", bytes.NewBuffer(videoJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestDeleteVideoFail(t *testing.T) {
	app := fiber.New()
	mockSentente := new(sentencesMock.MockSentences)
	handler := handlers.NewHandlerDeleteVideo(mockSentente)
	app.Delete("/test-video", handler.DeleteVideo)

	mockSentente.On("DeleteRecord", mock.Anything, "video", "id_video", "test-id-video", mock.Anything).Return(errors.New("error deleting"))

	video := dto.Video{
		Id_Video: "test-id-video",
	}

	videoJSON, err := json.Marshal(video)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/test-video", bytes.NewBuffer(videoJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	mockSentente.AssertCalled(t, "DeleteRecord", mock.Anything, "video", "id_video", "test-id-video", mock.Anything)
}
