package infrastructura

import (
	"techinical/challenges/handlers"
	"techinical/sentences"

	"github.com/gofiber/fiber/v2"

	sharedRepo "techinical/shared/repository"
)

func SetupRoutes(apiBase fiber.Router, chatGptApi sharedRepo.ChatGptApi, sentencesRepo sentences.Sentences) {

	handlerGet := handlers.NewHandlerGetChallenge(sentencesRepo)
	handlerPost := handlers.NewHandlerPostChallenge(chatGptApi)
	handlerPatch := handlers.NewHandlerPatchChallenge(chatGptApi)
	handlerDelete := handlers.NewHandlerDeleteChallenge(sentencesRepo)

	apiBase.Get("/challenges", handlerGet.GetChallenge)
	apiBase.Post("/challenges", handlerPost.PostChallenge)
	apiBase.Patch("/challenges", handlerPatch.PatchChallenge)
	apiBase.Delete("/challenges", handlerDelete.DeleteChallenge)
}
