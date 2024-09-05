package main

import (
	routersChallenges "techinical/challenges/infrastructura"
	"techinical/config"
	"techinical/db"
	"techinical/sentences"
	sharedRepo "techinical/shared/repository"
	routersUsers "techinical/users/infrastructura"
	routersVideos "techinical/video/infrastructura"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Iniciando configuración...")
	config := config.NewInitConfig()

	log.Info().Msg("Conectando a la base de datos...")
	db.InitDb(config)

	log.Info().Msg("Creando server...")
	app := fiber.New()
	apiBase := app.Group(config.NameService)

	log.Info().Msg("Inicializando apis...")
	chatGptApi := sharedRepo.NewChatGptApi(config.KeyOpenIa)
	sentencesRepo := sentences.NewSentences()

	// Definimos un grupo, con cada sub-contexto
	routersUsers.SetupRoutes(apiBase, chatGptApi, sentencesRepo)
	routersChallenges.SetupRoutes(apiBase, chatGptApi, sentencesRepo)
	routersVideos.SetupRoutes(apiBase, chatGptApi, sentencesRepo)

	app.Listen(":" + config.PortService)
}
