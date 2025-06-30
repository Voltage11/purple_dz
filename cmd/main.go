package main

import (
	"fmt"
	"log"
	"purple1/config"
	"purple1/internal/pages"
	"purple1/pkg/logger"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	log.Println("Загрузка конфигурации")
	cfg := config.GetConfig() 
	log.Println("Конфигурация загружена")

	engine := html.New("./html", ".html")
	engine.AddFuncMap(map[string]interface{}{
		"ToUpper": func (c string) string {
			return strings.ToUpper(c)
		},
	})

	app := fiber.New(fiber.Config{
			Views: engine,
		})
		
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	app.Use(slogfiber.New(customLogger))
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