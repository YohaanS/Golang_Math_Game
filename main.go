package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Create a route to handle the POST request
	app.Post("/post", func(c *fiber.Ctx) error {
		// Parse request body into User struct
		question := new(Question)
		if err := c.BodyParser(question); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse request body",
			})
		}

		// Do something with the user object, e.g. save to a database
		// ...

		// Return a response
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Question created successfully",
			"question":    question,
		})
	})

	// Start the Fiber app
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
