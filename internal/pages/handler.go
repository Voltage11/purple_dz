package pages

import (
	"purple1/pkg/tadapter"
	"purple1/views"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
	Title string // Заголовок категории
	ImageUrl string // Ссылка на изображение, заготовка на будущее
	Url string // Ссылка для перехода на категорию, заготовка на будущее
}

func NewPageHandler(router *fiber.App) {
	
	router.Get("/", func (c *fiber.Ctx) error {
		

		component := views.Main()
	
	return tadapter.Render(c, component)
	})
}