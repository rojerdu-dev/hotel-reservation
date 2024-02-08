package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rojerdu-dev/hotel-reservation/db"
	"github.com/rojerdu-dev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Four Seasons",
		Location: "London",
	}
	rooms := []types.Room{
		{Type: types.SingleRoomType,
			BasePrice: 99.99},
		{Type: types.DeluxeRoomType,
			BasePrice: 199.99},
		{Type: types.SeaSideRoomType,
			BasePrice: 299.99},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatalf("failed to insert room: %w", err)
		}
		fmt.Println(insertedRoom)
	}
}
