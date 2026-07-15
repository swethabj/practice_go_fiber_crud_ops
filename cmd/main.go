package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main(){

    fmt.Println("Start Server !!")
	app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))
}