package main

import (
	"log"

	"github.com/mahdic200/gofiber-template/Config"
	"github.com/mahdic200/gofiber-template/Routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New();

    Config.Connect()
    Routes.SetupRoutes(app)

    log.Fatal(app.Listen(":8000"))
}
