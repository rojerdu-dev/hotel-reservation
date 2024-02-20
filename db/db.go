package db

const (
	DBNAME     = "hotel-reservations"
	DBURI      = "mongodb://localhost:27017"
	TestDBNAME = "hotel-reservations-test"
)

type Pagination struct {
	Limit int64
	Page  int64
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
