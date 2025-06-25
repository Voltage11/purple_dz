package pages

import "github.com/gofiber/fiber/v2"

func NewPageHandler(router *fiber.App) {
	
	router.Get("/", func (c *fiber.Ctx) error {
		return  c.SendString("It's ok")
	})
}