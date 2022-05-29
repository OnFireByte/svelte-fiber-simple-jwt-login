package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/onfirebyte/simple-jwt-login/db"
	"github.com/onfirebyte/simple-jwt-login/routes"
)




func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}


	db.Connect()
	app := fiber.New()

	allowOrigin := os.Getenv("ALLOW_ORIGINS")
	if allowOrigin == "" {
		allowOrigin = "*"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigin,
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	
	}))

	routes.RoutesSetup(app)
	app.Listen(":5000")
}