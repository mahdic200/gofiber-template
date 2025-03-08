package Routes

import (
	"github.com/mahdic200/gofiber-template/Controllers/AuthController"
	"github.com/mahdic200/gofiber-template/Controllers/UserController"
	"github.com/mahdic200/gofiber-template/Validations/Admin/UserValidator"

	// "github.com/mahdic200/gofiber-template/Middlewares/Auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

    app.Post("/login", AuthController.Login).Name("app.user.index")
    // adminGroup := app.Group("/admin", Auth.AuthMiddleware)
    adminGroup := app.Group("/admin")

    userGroup := adminGroup.Group("user")
    userGroup.Get("/", UserController.Index).Name("app.user.index")
    userGroup.Get("/show/:id", UserController.Show).Name("app.user.show")
    userGroup.Post("/store", UserValidator.Store(), UserController.Store).Name("app.user.store")

    app.Use("*", func(c *fiber.Ctx) error {
        return c.Status(404).JSON(fiber.Map{
            "message": "Not found !",
        })
    })
}


