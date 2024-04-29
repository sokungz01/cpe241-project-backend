package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sokungz01/cpe241-project-backend/platform"
	"github.com/sokungz01/cpe241-project-backend/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	myDB, err := platform.NewSql(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	api := fiber.New()
	routes.RoutesRegister(api, myDB)
	api.Listen(":3000")
}
