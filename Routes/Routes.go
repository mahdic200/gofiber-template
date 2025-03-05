package Routes

import (
    "fiber/Controllers/AuthController"
    "fiber/Controllers/UserController"
    "fiber/Middlewares/Auth"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

    app.Post("/login", AuthController.Login).Name("app.user.index")
    adminGroup := app.Group("/admin", Auth.AuthMiddleware)

    userGroup := adminGroup.Group("user")
    userGroup.Get("/", UserController.Index).Name("app.user.index")
    userGroup.Get("/show/:id", UserController.Show).Name("app.user.show")
    userGroup.Post("/store", UserController.Store).Name("app.user.index")

}


