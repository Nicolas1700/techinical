package infrastructura

import (
	"techinical/sentences"
	"techinical/video/handlers"

	"github.com/gofiber/fiber/v2"

	sharedRepo "techinical/shared/repository"
)

func SetupRoutes(route fiber.Router, chatGptApi sharedRepo.ChatGptApi, sentences sentences.Sentences) {

	handlerGet := handlers.NewHandlerGetVideo(sentences)
	handlerPost := handlers.NewHandlerPostVideo(chatGptApi)
	handlerPatch := handlers.NewHandlerPatchVideo(chatGptApi)
	handlerDelete := handlers.NewHandlerDeleteVideo(sentences)

	route = route.Group("/videos")
	route.Get("", handlerGet.GetVideo)
	route.Post("", handlerPost.PostVideo)
	route.Patch("", handlerPatch.PatchVideo)
	route.Delete("", handlerDelete.DeleteVideo)
}
