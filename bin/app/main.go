package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var mg MongoInstance
type MongoInstance struct{
	Client *mongo.Client
	Db     *mongo.Database
}



const dbName = "fiber_mongodb"
const MongoURI = "127.0.0.1:27017/"+ dbName

type Employee struct{
	ID string
	Name string
	Salary float64
	Age float64
}

func ConnectDb() error{
  client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()
  if err != nil{
	return err
  }
  err = client.Connect(ctx)
  if err != nil{
	return err
  }
  db := client.Database(dbName)

  mg := MongoInstance{
	Client: client,
	Db: db,
  }
return nil
}







func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to fiber con mongodb")
}

func setupRoutes(app *fiber.App){
	
	//app.Post("/employee", welcome)
	//app.Put("/employee:id", welcome)
	//app.Delete("/employee:id", welcome)
	
}

func main(){

	err := ConnectDb()
	if err != nil{
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func (c *fiber.Ctx) error  {
		query := bson.D{}

	cursor, err :=	mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil{
		return c.Status(500).SendString(err.Error())
	}

		var employees []Employee = make([]Employee, 0)

         err = cursor.All(c.Context(), &employees)
		 if err != nil{
          return c.Status(500).SendString(err.Error())
		 }
		 return c.JSON(employees)
	})

	log.Fatal(app.Listen(":3000"))
}