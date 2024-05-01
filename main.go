package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sokungz01/cpe241-project-backend/config"
	"github.com/sokungz01/cpe241-project-backend/platform"
	"github.com/sokungz01/cpe241-project-backend/routes"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Can't load config", err)
	}

	myDB, err := platform.NewSql(cfg.DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	api := fiber.New()
	api.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))
	routes.RoutesRegister(api, myDB, cfg)
	api.Listen(":3000")
}
