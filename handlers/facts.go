package handlers

import (
    "github.com/gofiber/fiber/v2"
	"github.com/arthurstr/arthur-trivia/database"
	"github.com/arthurstr/arthur-trivia/models"
)

func ListFacts(c *fiber.Ctx) error {
    facts := []models.Fact{}
    database.DB.Db.Find(&facts)

    return c.Render("index", fiber.Map{
        "Title": "Arthur Trivia Time",
        "Subtitle": "Facts for funtimes with friends!",
		"Facts": facts,
    })
}

func NewFactView(c *fiber.Ctx) error{
	return c.Render("new", fiber.Map{
		"Title": "New fact",
		"Subtitle": "Add fact",
	})
}

func CreateFact(c *fiber.Ctx) error {
    fact := new(models.Fact)
    if err := c.BodyParser(fact); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&fact)

    return ConfirmationView(c)
}

func ConfirmationView(c *fiber.Ctx) error {
    return c.Render("confirmation", fiber.Map{
        "Title":    "Fact added successfully",
        "Subtitle": "Add more wonderful facts to the list!",
    })
}