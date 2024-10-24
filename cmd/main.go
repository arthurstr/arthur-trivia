package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/arthurstr/arthur-trivia/database"
	"github.com/gofiber/template/html/v2"
	"github.com/arthurstr/arthur-trivia/handlers"
)

func main() {
    database.ConnectDb()

	engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine, 
		ViewsLayout: "layouts/main",
    })

    setupRoutes(app)

    app.Listen(":3000")
}
