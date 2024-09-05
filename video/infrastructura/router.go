package infrastructura

import (
	"techinical/sentences"
	"techinical/video/handlers"

	"github.com/gofiber/fiber/v2"

	sharedRepo "techinical/shared/repository"
)

func SetupRoutes(apiBase fiber.Router, chatGptApi sharedRepo.ChatGptApi, sentences sentences.Sentences) {

	handlerGet := handlers.NewHandlerGetVideo(sentences)
	handlerPost := handlers.NewHandlerPostVideo(chatGptApi)
	handlerPatch := handlers.NewHandlerPatchVideo(chatGptApi)
	handlerDelete := handlers.NewHandlerDeleteVideo(sentences)

	apiBase.Get("/videos", handlerGet.GetVideo)
	apiBase.Post("/videos", handlerPost.PostVideo)
	apiBase.Patch("/videos", handlerPatch.PatchVideo)
	apiBase.Delete("/videos", handlerDelete.DeleteVideo)
}
