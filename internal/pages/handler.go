package pages

import (
	"github.com/gofiber/fiber/v2"
)

type Category struct {
	Title string // Заголовок категории
	ImageUrl string // Ссылка на изображение, заготовка на будущее
	Url string // Ссылка для перехода на категорию, заготовка на будущее
}

func NewPageHandler(router *fiber.App) {
	
	router.Get("/", func (c *fiber.Ctx) error {
		categories := []Category {
			{
				Title: "Животные",
				ImageUrl: "",
				Url: "#",
			},
			{
				Title: "Бабочки",
				ImageUrl: "",
				Url: "#",
			},
			{
				Title: "Крокодилы",
				ImageUrl: "",
				Url: "#",
			},
			{
				Title: "Сверчки",
				ImageUrl: "",
				Url: "#",
			},
			{
				Title: "Черепахи",
				ImageUrl: "",
				Url: "#",
			},
			{
				Title: "Природа",
				ImageUrl: "",
				Url: "#",
			},
		}

		data := struct {
			Categories []Category
		}{Categories: categories,}

		return c.Render("index", data)
	})
}