package main

import (
	"fmt"
	"log"
	"purple1/config"
	"purple1/internal/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	log.Println("Загрузка конфигурации")
	cfg := config.GetConfig() 
	log.Println("Конфигурация загружена")

	app := fiber.New()
	app.Use(recover.New())
	pages.NewPageHandler(app)

	if cfg.Server.Debug {
		// Просто проверка загрузки bool значения из .env, не несет логики в данном примере
		log.Println("Запуск сервера в Debug")
	}

	log.Printf("Запуск сервера %s:%s", cfg.Server.Host, cfg.Server.Port)

	runServer := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port) 
	log.Fatal(app.Listen(runServer))
}