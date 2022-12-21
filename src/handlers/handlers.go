package handlers

import (
	"log"
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

func ItemCreate(c *fiber.Ctx) error {
	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		return err
	}
	log.Println(item.Title)
	log.Println(item.Content)
	return c.Send(c.Body())
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
