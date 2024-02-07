package main

import (
	"context"
	"flag"
	"github.com/rojerdu-dev/hotel-reservation/db"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/rojerdu-dev/hotel-reservation/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbURI    = "mongodb://localhost:27017"
	dbName   = "hotel-reservations"
	userColl = "users"
)

var (
	config = fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.JSON(map[string]string{"error": err.Error()})
		},
	}
)

func main() {
	addr := flag.String("address", ":5000", "API server listen address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	app.Get("/home", handleHome)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	app.Listen(*addr)
}

func handleHome(c fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "home is where the heart is"})
}
