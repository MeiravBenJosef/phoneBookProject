package main

import (
    "github.com/gofiber/fiber/v2"


)

func main() {

    app := fiber.New()

    setupRoutes(app)

    //the app will run on port 3000
    app.Listen(":3000")
}