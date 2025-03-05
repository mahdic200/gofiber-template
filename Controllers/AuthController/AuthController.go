package AuthController

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
    
    return c.Status(200).JSON(fiber.Map{
        "token" : "Login route",
    })
}
