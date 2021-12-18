package main

import (
	"BazarMessenger/config"
	"BazarMessenger/user/controller"
	"BazarMessenger/user/mongodb"
	"BazarMessenger/user/usercase"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
)
import "github.com/gofiber/fiber/v2"

var MongoCollection *mongo.Collection

func init() {
	MongoCollection = config.InitMongodb()
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	MongoCollection = config.InitMongodb()

	userRepository := mongodb.NewMongodbUserRepository(MongoCollection)
	userUsecase := usercase.NewUserUsecase(userRepository)
	controller.NewUserController(app, userUsecase)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	app.Listen(":3000")
}
