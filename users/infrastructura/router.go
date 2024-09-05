package infrastructura

import (
	"techinical/sentences"
	"techinical/users/handlers"

	"github.com/gofiber/fiber/v2"

	sharedRepo "techinical/shared/repository"
)

func SetupRoutes(apiBase fiber.Router, chatGptApi sharedRepo.ChatGptApi, sentences sentences.Sentences) {

	handlerGet := handlers.NewHandlerGetUser(sentences)
	handlerPost := handlers.NewHandlerPostUser(chatGptApi)
	handlerPatch := handlers.NewHandlerPatchUser(chatGptApi)
	handlerDelete := handlers.NewHandlerDeleteUser(sentences)

	apiBase.Get("/users", handlerGet.GetUser)
	apiBase.Post("/users", handlerPost.PostUser)
	apiBase.Patch("/users", handlerPatch.PatchUser)
	apiBase.Delete("/users", handlerDelete.DeleteUser)
}
