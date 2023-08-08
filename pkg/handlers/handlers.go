package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ToDo struct {
	ID    string `json:"id" xml:"id" form:"id"`
	Title string `json:"title" xml:"title" form:"title"`
}

func GetHealth(c *fiber.Ctx) error {
	fmt.Println(c.GetReqHeaders())
	return c.JSON(fiber.Map{
		"ok": true,
	})
}

func GetToDos(c *fiber.Ctx) error {
	return c.JSON([]ToDo{
		{ID: "1", Title: "Buy milk"},
		{ID: "2", Title: "Buy eggs"},
	})
}

func CreateToDo(c *fiber.Ctx) error {
	id := uuid.New().String()
	toDo := new(ToDo)
	if err := c.BodyParser(toDo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if toDo.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing title",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id, "title": toDo.Title})
}

func DeleteToDo(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "It is not a valid UUID",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"id": id,
	})
}
