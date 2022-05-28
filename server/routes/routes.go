package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onfirebyte/simple-jwt-login/controllers"
)

func RoutesSetup(app *fiber.App){
	
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)
	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	app.Post("api/note", controllers.AddNote)
	app.Get("api/note", controllers.SeeNote)

}