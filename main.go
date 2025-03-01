package main

import "log"
import "fiber/Routes"
import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New();

    routes.SetupRoutes(app)

    log.Fatal(app.Listen(":8000"))
}

