package routes

import "github.com/gofiber/fiber/v2"
import "fiber/Controllers/UserController"

func SetupRoutes(app *fiber.App) {

    userGroup := app.Group("/user")
    userGroup.Get("/", UserController.Index).Name("app.user.index")
}


