package infrastructura

import (
	"techinical/users/handlers"

	"github.com/gofiber/fiber/v2"

	sharedRepo "techinical/shared/repository"
)

func SetupRoutes(apiBase fiber.Router, chatGptApi sharedRepo.ChatGptApi) {

	handlerGet := handlers.NewHandlerGetUser()
	handlerPost := handlers.NewHandlerPostUser(chatGptApi)
	handlerPatch := handlers.NewHandlerPatchUser(chatGptApi)
	handlerDelete := handlers.NewHandlerDeleteUser()

	apiBase.Get("/users", handlerGet.GetUser)
	apiBase.Post("/users", handlerPost.PostUser)
	apiBase.Patch("/users", handlerPatch.PatchUser)
	apiBase.Delete("/users", handlerDelete.DeleteUser)
}
