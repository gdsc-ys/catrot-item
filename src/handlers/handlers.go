package handlers

import (
	"fmt"
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
	log.Println(item.Category)

	if form, err := c.MultipartForm(); err == nil {
		files := form.File["img"]

		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
				return err
			}
		}
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"success" : true,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
