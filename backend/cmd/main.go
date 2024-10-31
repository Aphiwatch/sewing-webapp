package main

import (
    "github.com/gofiber/fiber/v2"
    "log"
)

func main() {
    app := fiber.New()

    app.Get("/api", func(c *fiber.Ctx) error {
        return c.SendString("Hello from Go Fiber!")
    })

    log.Fatal(app.Listen(":8080"))
}
