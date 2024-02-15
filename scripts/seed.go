package main

import (
	"context"
	"fmt"
	"github.com/rojerdu-dev/hotel-reservation/db/fixtures"
	"log"
	"time"

	"github.com/rojerdu-dev/hotel-reservation/api"
	"github.com/rojerdu-dev/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Booking: db.NewMongoBookingStore(client),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "michael", "jordan", false)
	fmt.Println("michael jordan ->", api.CreateTokenFromUser(user))

	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin ->", api.CreateTokenFromUser(admin))

	hotel := fixtures.AddHotel(store, "Classy Hotel", "New York City", 5, nil)
	fmt.Println("hotel ->", hotel.ID)

	room := fixtures.AddRoom(store, "large", true, 299.99, hotel.ID)
	fmt.Println("room ->", room.ID)

	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))
	fmt.Println("booking ->", booking.ID)
}

//func init() {
//	var err error
//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatalf("Failed to ping MongoDB server: %v", err)
//	}
//
//	if err = client.Database(db.DBNAME).Drop(ctx); err != nil {
//		log.Fatal(err)
//	}
//	hotelStore := db.NewMongoHotelStore(client)
//	roomStore := db.NewMongoRoomStore(client, hotelStore)
//	userStore := db.NewMongoUserStore(client)
//	bookingStore := db.NewMongoBookingStore(client)
//	fmt.Println(hotelStore)
//	fmt.Println(roomStore)
//	fmt.Println(userStore)
//	fmt.Println(bookingStore)
//}

//var (
//	client       *mongo.Client
//	roomStore    db.RoomStore
//	hotelStore   db.HotelStore
//	userStore    db.UserStore
//	bookingStore db.BookingStore
//	ctx          = context.Background()
//)
//
//func seedUser(isAdmin bool, fname, lname, email, password string) *types.User {
//	user, err := types.NewUserFromParams(types.CreateUserParams{
//		FirstName: fname,
//		LastName:  lname,
//		Email:     email,
//		Password:  password,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	user.IsAdmin = isAdmin
//	insertedUser, err := userStore.InsertUser(context.TODO(), user)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%s -> %s\n", user.Email, api.CreateTokenFromUser(user))
//	return insertedUser
//}
//
//func seedRoom(size string, seaSide bool, price float64, hotelID primitive.ObjectID) *types.Room {
//	room := &types.Room{
//		Size:    size,
//		Seaside: seaSide,
//		Price:   price,
//		HotelID: hotelID,
//	}
//	insertedRoom, err := roomStore.InsertRoom(context.Background(), room)
//	if err != nil {
//		return nil
//	}
//	return insertedRoom
//}
//
//func seedBooking(userID, roomID primitive.ObjectID, from, until time.Time) {
//	booking := &types.Booking{
//		UserID:    userID,
//		RoomID:    roomID,
//		FromDate:  from,
//		UntilDate: until,
//	}
//	resp, err := bookingStore.InsertBooking(context.Background(), booking)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("booking:", resp.ID)
//}
//
//func seedHotel(name, location string, rating int) *types.Hotel {
//	hotel := types.Hotel{
//		Name:     name,
//		Location: location,
//		Rooms:    []primitive.ObjectID{},
//		Rating:   rating,
//	}
//	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return insertedHotel
//}
