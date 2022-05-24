package main

import (
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
	app.Use(cors.New(cors.Config{AllowCredentials: true,}))

	routes.RoutesSetup(app)
	app.Listen(":5000")
}