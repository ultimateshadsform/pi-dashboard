package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/ultimateshadsform/pi-dashboard/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}

	app := fiber.New()

	app.Static("/", "./static")

	routes.SetupRouter(app)

	log.Fatal(app.Listen("0.0.0.0:3000"))
}
