package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type MongoInstance struct{
	Client
	Db
}

var mg MongoInstance

const dbName = "fiber_mongodb"
const mongoURI = "127.0.0.1:27017"+ dbName



func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to fiber con mongodb")
}

func setupRoutes(app *fiber.App){
	app.Get("/employee", welcome)
	app.Post("/employee", welcome)
	app.Put("/employee:id", welcome)
	app.Delete("/employee:id", welcome)
	
}

func main(){
	app := fiber.New()

	//app.Get("/", welcome)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}