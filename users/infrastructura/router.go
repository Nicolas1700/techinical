package infrastructura

import (
	"techinical/sentences"
	"techinical/users/handlers"

	"github.com/gofiber/fiber/v2"

	sharedRepo "techinical/shared/repository"
)

func SetupRoutes(route fiber.Router, chatGptApi sharedRepo.ChatGptApi, sentences sentences.Sentences) {

	handlerGet := handlers.NewHandlerGetUser(sentences)
	handlerPost := handlers.NewHandlerPostUser(chatGptApi)
	handlerPatch := handlers.NewHandlerPatchUser(chatGptApi)
	handlerDelete := handlers.NewHandlerDeleteUser(sentences)

	route = route.Group("/users")
	route.Get("", handlerGet.GetUser)
	route.Post("", handlerPost.PostUser)
	route.Patch("", handlerPatch.PatchUser)
	route.Delete("", handlerDelete.DeleteUser)
}
