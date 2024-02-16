package fixtures

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/rojerdu-dev/hotel-reservation/db"
	"github.com/rojerdu-dev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBooking(store *db.Store, userID, roomID primitive.ObjectID, from, until time.Time) *types.Booking {
	source := rand.NewSource(time.Now().UnixNano())
	randomNumGen := rand.New(source)
	booking := &types.Booking{
		UserID:     userID,
		RoomID:     roomID,
		NumPersons: randomNumGen.Intn(5) + 1,
		FromDate:   from,
		UntilDate:  until,
		Canceled:   false,
	}
	insertedBooking, err := store.Booking.InsertBooking(context.TODO(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}

func AddRoom(store *db.Store, size string, seaSide bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		ID:      primitive.ObjectID{},
		Size:    size,
		Seaside: seaSide,
		Price:   price,
		HotelID: hotelID,
	}
	insertedRoom, err := store.Room.InsertRoom(context.TODO(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func AddHotel(store *db.Store, name, location string, rating int, rooms []primitive.ObjectID) *types.Hotel {
	var roomIDs = rooms
	if rooms == nil {
		roomIDs = []primitive.ObjectID{}
	}
	hotel := types.Hotel{
		ID:       primitive.ObjectID{},
		Name:     name,
		Location: location,
		Rooms:    roomIDs,
		Rating:   rating,
	}
	insertedHotel, err := store.Hotel.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddUser(store *db.Store, fname, lname string, admin bool) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: fname,
		LastName:  lname,
		Email:     fmt.Sprintf("%s_%s@email.com", fname, lname),
		Password:  fmt.Sprintf("%s_%s", fname, lname),
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}
