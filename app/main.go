package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)
func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to fiber con mongodb")
}

func setupRoutes(app *fiber.App){
	app.Get("/", welcome)
	
}

func main(){
	app := fiber.New()

	//app.Get("/", welcome)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}