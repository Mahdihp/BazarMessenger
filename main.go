package main

import (
	"BazarMessenger/config"
	"go.mongodb.org/mongo-driver/mongo"
)
import "github.com/gofiber/fiber/v2"

var MongoCollection *mongo.Collection

func init() {
	MongoCollection = config.InitMongodb()

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
