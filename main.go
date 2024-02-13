package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rojerdu-dev/hotel-reservation/api"
	"github.com/rojerdu-dev/hotel-reservation/db"
	"github.com/rojerdu-dev/hotel-reservation/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	userColl = "users"
)

var (
	config = fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.JSON(map[string]string{"error": err.Error()})
		},
	}
)

func main() {
	addr := flag.String("address", ":5000", "API server listen address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	// handlers init
	var (
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		userStore    = db.NewMongoUserStore(client)
		bookingStore = db.NewMongoBookingStore(client)
		store        = &db.Store{
			User:    userStore,
			Hotel:   hotelStore,
			Room:    roomStore,
			Booking: bookingStore,
		}
		userHandler  = api.NewUserHandler(db.NewMongoUserStore(client))
		hotelHandler = api.NewHotelHandler(store)
		authHandler  = api.NewAuthHandler(userStore)
		roomHandler  = api.NewRoomHandler(store)
		app          = fiber.New(config)
		auth         = app.Group("/api")
		apiv1        = app.Group("/api/v1", middleware.JWTAuthentication(userStore))
	)

	// auth handler
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// versioned api routes
	// user handlers
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	// hotel handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	// room handlers
	apiv1.Get("/room", roomHandler.HandleGetRooms)
	apiv1.Post("/room/:id/book", roomHandler.HandleBookRoom)

	app.Listen(*addr)
}
