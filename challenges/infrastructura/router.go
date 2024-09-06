package infrastructura

import (
	"techinical/challenges/handlers"
	"techinical/sentences"

	"github.com/gofiber/fiber/v2"

	sharedRepo "techinical/shared/repository"
)

func SetupRoutes(route fiber.Router, chatGptApi sharedRepo.ChatGptApi, sentencesRepo sentences.Sentences) {

	handlerGet := handlers.NewHandlerGetChallenge(sentencesRepo)
	handlerPost := handlers.NewHandlerPostChallenge(chatGptApi)
	handlerPatch := handlers.NewHandlerPatchChallenge(chatGptApi)
	handlerDelete := handlers.NewHandlerDeleteChallenge(sentencesRepo)

	route = route.Group("/challenges")
	route.Get("", handlerGet.GetChallenge)
	route.Post("", handlerPost.PostChallenge)
	route.Patch("", handlerPatch.PatchChallenge)
	route.Delete("", handlerDelete.DeleteChallenge)
}
