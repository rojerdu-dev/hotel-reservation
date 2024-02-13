package db

const (
	DBNAME     = "hotel-reservations"
	DBURI      = "mongodb://localhost:27017"
	TestDBNAME = "hotel-reservations-test"
)

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
