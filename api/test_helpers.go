package api

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rojerdu-dev/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	client *mongo.Client
	*db.Store
}

func (tdb *testdb) teardown(t *testing.T) {
	dbName := os.Getenv(db.MongoDBNameEnvName)
	if err := tdb.client.Database(dbName).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(_ *testing.T) *testdb {
	dbURI := os.Getenv("MONGO_DB_URL_TEST")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	return &testdb{
		client,
		&db.Store{
			User:    db.NewMongoUserStore(client),
			Hotel:   hotelStore,
			Room:    db.NewMongoRoomStore(client, hotelStore),
			Booking: db.NewMongoBookingStore(client),
		},
	}
}
