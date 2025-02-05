package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		p := new(Person)
		if err := c.BodyParser(p); err != nil {
			return err
		}

		return c.SendString("Hello, World!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})

	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	app.Listen(":3000")
}

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}
