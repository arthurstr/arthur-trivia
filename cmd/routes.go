package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/arthurstr/arthur-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", handlers.ListFacts)


	app.Get("/fact",handlers.NewFactView)
	app.Post("/fact", handlers.CreateFact)
}