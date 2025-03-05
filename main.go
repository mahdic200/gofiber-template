package main

import (
    "log"
    "fiber/Routes"
    "github.com/gofiber/fiber/v2"
    "fiber/Config"
    
)

func main() {
    app := fiber.New();

    Config.Connect()
    Routes.SetupRoutes(app)

    log.Fatal(app.Listen(":8000"))
}

