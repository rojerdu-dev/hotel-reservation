package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/rojerdu-dev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/rojerdu-dev/hotel-reservation/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    = "mongodb://localhost:27017"
	dbName   = "hotel-reservations"
	userColl = "users"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbName).Collection(userColl)

	user := types.User{
		FirstName: "James",
		LastName:  "Bond",
	}

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var james types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
		log.Fatal(err)
	}

	fmt.Println(james)

	addr := flag.String("address", ":5000", "API server listen address")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/home", handleHome)
	apiv1.Get("/user/:id", api.HandleGetUser)
	apiv1.Get("/user", api.HandleGetUsers)
	app.Listen(*addr)
}

func handleHome(ctx fiber.Ctx) error {
	return ctx.JSON(map[string]string{"msg": "home is where the heart is"})
}
