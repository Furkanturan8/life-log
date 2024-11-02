package app

import (
	"fmt"
	"github.com/bmdavis419/the-better-backend/pkg/config"
	"github.com/bmdavis419/the-better-backend/pkg/database"
	"os"

	"github.com/bmdavis419/the-better-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() error {
	// Ortam değişkenlerini yükle
	err := config.LoadENV()
	if err != nil {
		return fmt.Errorf("environment variables loading error: %w", err)
	}

	// PostgreSQL veritabanını başlat
	err = database.StartPostgresDB()
	if err != nil {
		return fmt.Errorf("database connection error: %w", err)
	}
	defer database.ClosePostgresDB() // uygulama kapandıktan sonra veritabanını kapat

	// Fiber uygulamasını oluştur
	app := fiber.New()

	// Middleware'leri ekle
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// Rotaları yapılandır
	router.SetupRoutes(app)

	// Portu al ve uygulamayı başlat
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // varsayılan port
	}

	// Uygulamayı başlat ve hataları yakala
	if err = app.Listen(":" + port); err != nil {
		return fmt.Errorf("application startup error: %w", err)
	}

	return nil
}
