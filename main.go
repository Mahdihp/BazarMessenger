package main

import (
	"BazarMessenger/config"
	"BazarMessenger/user/controller"
	"BazarMessenger/user/mongodb"
	"BazarMessenger/user/usercase"
	"go.mongodb.org/mongo-driver/mongo"
)
import "github.com/gofiber/fiber/v2"

var MongoCollection *mongo.Collection

func init() {
	MongoCollection = config.InitMongodb()

}

func main() {
	app := fiber.New()
	userRepository := mongodb.NewMongodbUserRepository(MongoCollection)
	userUsecase := usercase.NewUserUsecase(userRepository)
	controller.NewUserController(app, userUsecase)

	app.Listen(":3000")
}
