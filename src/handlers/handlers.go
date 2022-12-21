package handlers

import (
	"src/database"
	"src/models"

	"github.com/gofiber/fiber/v2"
)

func ItemList(c *fiber.Ctx) error {
	var items []models.Item
	db := database.DB
	db.Find(&items)
	return c.JSON(fiber.Map{
		"success":true,
		"itmes": items,
	})
}

func ItemDetail(c *fiber.Ctx) error {
	var item models.Item
	itemId := c.Params("itemId")
	db := database.DB
	db.First(&item, itemId)
	return c.JSON(fiber.Map{
		"success":true,
		"item": item,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
