package handlers

import (
	"src/database"
	"src/models"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/utils"
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

// // UserGet returns a user
// func UserList(c *fiber.Ctx) error {
// 	users := database.Get()

// 	return c.JSON(fiber.Map{
// 		"success": true,
// 		"users":   users,
// 	})
// }

// // UserCreate registers a user
// func UserCreate(c *fiber.Ctx) error {
// 	user := &models.User{
// 		// Note: when writing to external database,
// 		// we can simply use - Name: c.FormValue("user")
// 		Name: utils.CopyString(c.FormValue("user")),
// 	}
// 	database.Insert(user)

// 	return c.JSON(fiber.Map{
// 		"success": true,
// 		"user":    user,
// 	})
// }

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
